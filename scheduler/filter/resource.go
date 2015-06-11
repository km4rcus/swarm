package filter

import (
	"errors"
	"strconv"

	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/pkg/bitmap"
	"github.com/docker/swarm/pkg/utils"
	"github.com/docker/swarm/scheduler/node"
)

// ResourceFilter filter the nodes that satisfy resource constraint for the container
// resources: CPU, Memory, Cpuset
type ResourceFilter struct {
}

// Name returns the name of the filter
func (r *ResourceFilter) Name() string {
	return "resource"
}

// TODO:
// Those cases happen if there is the possibility to schedule containers with no particular
// contraints on the same nodes ...
//    a) if the container has a cpuset constraint cannot be scheduled on a node with containers
//		 without cpuset contraints
//    b) if a container has no resource constraint cannot be scheduled on a node with containers
//		 with cpuset constraint
//    containers with cpuset constraint can be scheduled on node with containers with cpu constraint ??
//    containers with cpu constraint can be scheduled on nodes with containers with cpuset contraint ?
//    is it possible to introduce a soft constraint for resources (cpu/cpuset)??
//
// Filter is exported 
func (r *ResourceFilter) Filter(config *cluster.ContainerConfig, nodes []*node.Node) ([]*node.Node, error) {
	memory := config.Memory
	cpus := config.CpuShares

	candidates := []*node.Node{}
	for _, node := range nodes {
		// we need to setup the container cpuset
		cpuset := bitmap.ZeroBitmap(uint64(node.TotalCpus))
		for _, s := range utils.StringListSplit(config.Cpuset) {
			bitPos, err := strconv.ParseInt(s, 10, 64)
			// we need to check if cpuset specified are within NCPUS boundary;
			// this is due to the fact that cpuset can be specified without specifying # of cpus
			if err == nil {
				if bitPos >= node.TotalCpus {
					// we set cpus to bitPos+1 and break so next "if node.TotalCpus-cpus>=0" will fail and the node will 
					// not be a candidate
					cpus = bitPos + 1
					break
				}
				bitmap.SetBit(&cpuset, uint64(bitPos))
			}
		}
		if node.TotalMemory-memory >= 0 && node.TotalCpus-cpus >= 0 && bitmap.IsZero(bitmap.AND(node.Cpuset, cpuset)) {
			candidates = append(candidates, node)
		}
	}

	if len(candidates) == 0 {
		return nil, errors.New("No Resources Available to Schedule Container")
	}
	nodes = candidates
	return nodes, nil
}






