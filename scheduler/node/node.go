package node

import (
	"strconv"
	"strings"

	"github.com/docker/swarm/cluster"
	"github.com/docker/swarm/pkg/bitmap"
	"github.com/docker/swarm/pkg/utils"
)

// Node is an abstract type used by the scheduler
type Node struct {
	ID         string
	IP         string
	Addr       string
	Name       string
	Labels     map[string]string
	Containers []*cluster.Container
	Images     []*cluster.Image

	UsedMemory  int64
	UsedCpus    int64
	TotalMemory int64
	TotalCpus   int64
	Cpuset      bitmap.Bitmap

	IsHealthy bool
}

// NewNode creates a node from an engine
func NewNode(e *cluster.Engine) *Node {
	return &Node{
		ID:          e.ID,
		IP:          e.IP,
		Addr:        e.Addr,
		Name:        e.Name,
		Labels:      e.Labels,
		Containers:  e.Containers(),
		Images:      e.Images(),
		UsedMemory:  e.UsedMemory(),
		UsedCpus:    e.UsedCpus(),
		TotalMemory: e.TotalMemory(),
		TotalCpus:   e.TotalCpus(),
		Cpuset:      e.Cpuset,
		IsHealthy:   e.IsHealthy(),
	}
}

// Container returns the container with IDOrName in the engine.
func (n *Node) Container(IDOrName string) *cluster.Container {
	// Abort immediately if the name is empty.
	if len(IDOrName) == 0 {
		return nil
	}

	for _, container := range n.Containers {
		// Match ID prefix.
		if strings.HasPrefix(container.Id, IDOrName) {
			return container
		}

		// Match name, /name or engine/name.
		for _, name := range container.Names {
			if name == IDOrName || name == "/"+IDOrName || container.Engine.ID+name == IDOrName || container.Engine.Name+name == IDOrName {
				return container
			}
		}
	}

	return nil
}

// AddContainer inject a container into the internal state.
func (n *Node) AddContainer(container *cluster.Container) error {
	if container.Config != nil {
        // Update node resources: usedMemory, usedCpus and Cpuset
        n.UsedMemory = n.UsedMemory + container.Config.Memory
        n.UsedCpus = n.UsedCpus + container.Config.CpuShares
        for _, s := range utils.StringListSplit(container.Config.Cpuset) {
        		p,err := strconv.ParseUint(s, 10, 64)
        	    if err == nil {
        	    	bitmap.SetBit(&n.Cpuset, p)
        	    }
        }
	}
	n.Containers = append(n.Containers, container)
	return nil
}