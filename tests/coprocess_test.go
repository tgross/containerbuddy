package tests

import (
	//	"strings"
	"fmt"
	"testing"
)

func TestCoprocesses(t *testing.T) {
	const yml = "testdata/TestCoprocesses/compose.yml"
	composeUp(t, yml)
	defer func() {
		if !t.Failed() {
			err := composeTearDown(yml)
			if err != nil {
				t.Fatal(err)
			}
		}
	}()

	consulId := composePsIds(t, yml, "consul")
	out := dockerExec(t, consulId[:10], "consul", "info", "|", "grep", "leader_addr")
	fmt.Println(out)

	// consul, err := NewConsulProbe()
	// if err != nil {
	// 	t.Fatalf("Expected to be able to create consul client before the test starts: %s\n", err)
	// }

	// err = consul.WaitForLeader()
	// if err != nil {
	// 	t.Fatalf("Expected consul to elect leader before the test starts: %s\n", err)
	// }

}
