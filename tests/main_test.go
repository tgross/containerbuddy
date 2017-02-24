package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestMain will build all container images we need for the test
// suite. It will automatically be run ahead of any tests here.
func TestMain(m *testing.M) {
	cwd, _ := os.Getwd()
	binPath := filepath.Join(cwd, filepath.Dir("../build/containerpilot"), "containerpilot")
	if _, err := os.Stat(binPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("ContainerPilot not built. Did you make?")
			os.Exit(1)
		}
	}
	fmt.Printf("ContainerPilot binary found at: %s\n", binPath)
	os.Setenv("CONTAINERPILOT_BIN", binPath)

	fmt.Println("=== creating test fixtures")
	err := composeBuild("testdata/build.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exitCode := m.Run()
	if exitCode == 0 {
		fmt.Println("=== tearing down containers")
		err := composeTearDown("testdata/build.yml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	os.Exit(exitCode)
}
