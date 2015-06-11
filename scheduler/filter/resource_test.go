package filter

import (
	"testing"

	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/scheduler/node"
	"github.com/docker/swarm/pkg/bitmap"
	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
)

func testResources() []*node.Node {
	return []*node.Node{
		{
			ID:          "node-0-id",
			Name:        "node-0-name",
			Addr:        "node-0",
			Cpuset:      bitmap.ZeroBitmap(4),
			UsedMemory:  0,
			UsedCpus:    0,
			TotalMemory: 8,
			TotalCpus:   8,
		},

		{
			ID:          "node-1-id",
			Name:        "node-1-name",
			Addr:        "node-1",
			Cpuset:      bitmap.ZeroBitmap(4),
			UsedMemory:  1,
			UsedCpus:    1,
			TotalMemory: 6,
			TotalCpus:   6,
		},

		{
			ID:          "node-2-id",
			Name:        "node-2-name",
			Addr:        "node-2",
			Cpuset:      bitmap.ZeroBitmap(4),
			UsedMemory:  2,
			UsedCpus:    2,
			TotalMemory: 4,
			TotalCpus:   4,
		},
	}
}

func TestResourceFilter(t *testing.T) {
	var (
		f      = ResourceFilter{}
		nodes  = testResources()
		result []*node.Node
		err    error
	)

	
	// Testing CPUSET
	// Set cpuset used in the nodes -> Node0: Cpuset(0,1), Node1: Cpuset(0,1) Node2: Cpuset(0,2)
	bitmap.SetBit(&nodes[0].Cpuset, 0)
	bitmap.SetBit(&nodes[0].Cpuset, 1)
	bitmap.SetBit(&nodes[1].Cpuset, 0)
	bitmap.SetBit(&nodes[1].Cpuset, 1)
	bitmap.SetBit(&nodes[2].Cpuset, 0)
	bitmap.SetBit(&nodes[2].Cpuset, 2)

	// Without resource constraints we should get the unfiltered list of nodes back.
	result, err = f.Filter(&cluster.ContainerConfig{}, nodes)
	assert.NoError(t, err)
	assert.Equal(t, result, nodes)
	
	// Set a cpuset resource constraint that cannot be fulfilled and expect an error back.
	result, err = f.Filter(cluster.BuildContainerConfig(dockerclient.ContainerConfig{Cpuset: "0"}), nodes)
	assert.Error(t, err)

	// // Set a cpuset constraint that can only be filled by a single node.
	result, err = f.Filter(cluster.BuildContainerConfig(dockerclient.ContainerConfig{Cpuset: "1"}), nodes)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, result[0], nodes[2])

	// This cpuset constraint can be fulfilled by a subset of nodes.
	result, err = f.Filter(cluster.BuildContainerConfig(dockerclient.ContainerConfig{Cpuset: "2"}), nodes)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.NotContains(t, result, nodes[2])

	//TODO: test CPU/MEM
	// Node0: Cpu: 0/8, Mem: 0/8 Node1: Cpu 1/6 Mem 1/6, Node2: Cpu 2/4 Mem 2/4
	
}
