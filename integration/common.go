package integration

import (
	"net"
	"testing"

	"github.com/taskgraph/taskgraph"
	"github.com/taskgraph/taskgraph/framework"
)

func createListener(t *testing.T) net.Listener {
	l, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("net.Listen(\"tcp4\", \"\") failed: %v", err)
	}
	return l
}

func drive(t *testing.T, jobName string, etcds []string, taskBuilder taskgraph.TaskBuilder, topo map[string]taskgraph.Topology) {

	bootstrap := framework.NewBootStrap(jobName, etcds, createListener(t), nil)
	bootstrap.SetTaskBuilder(taskBuilder)
	for i, _ := range topo {
		bootstrap.AddLinkage(i, topo[i])
	}
	bootstrap.Start()
}
