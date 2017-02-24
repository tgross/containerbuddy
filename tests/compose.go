package tests

import (
	"fmt"
	"github.com/joyent/containerpilot/commands"
	"testing"
)

const projectName = "cptests"

// helper function to generate Compose argument string
func compose(file string, composeArgs ...string) []string {
	args := []string{
		"docker-compose",
		"-p", projectName,
		"-f", file,
	}
	for _, arg := range composeArgs {
		args = append(args, arg)
	}
	return args

}

// helper function running 'docker-compose build'
func composeBuild(composeFileName string) error {
	cmd, err := commands.NewCommand(compose(composeFileName, "build"), "5m")
	if err != nil {
		return err
	}
	return commands.RunWithTimeout(cmd, nil)
}

// helper function running 'docker-compose rm -f'
func composeTearDown(composeFileName string) error {
	cmd, err := commands.NewCommand(compose(composeFileName, "stop"), "5m")
	if err != nil {
		return err
	}
	err = commands.RunWithTimeout(cmd, nil)
	if err != nil {
		return err
	}
	cmd, err = commands.NewCommand(compose(composeFileName, "rm", "-f"), "5m")
	if err != nil {
		return err
	}
	return commands.RunWithTimeout(cmd, nil)
}

// helper function running 'docker-compose up'
func composeUp(t *testing.T, composeFileName string) {
	cmd, err := commands.NewCommand(
		compose(composeFileName, "up", "-d"), "1m")
	if err != nil {
		t.Fatal(err)
	}
	if err := commands.RunWithTimeout(cmd, nil); err != nil {
		t.Fatal(err)
	}
}

// helper function running 'docker-compose run'
func composeRun(t *testing.T, composeFileName, serviceName string) string {
	cmd, err := commands.NewCommand(
		compose(composeFileName, "run", serviceName), "1m")
	if err != nil {
		t.Fatal(err)
	}
	out, runErr := commands.RunAndWaitForOutput(cmd)
	if runErr != nil {
		t.Fatal(runErr)
	}
	return out
}

// helper function running 'docker-compose logs'
func composeLogs(t *testing.T, composeFileName, serviceName string) string {
	cmd, err := commands.NewCommand(
		compose(composeFileName, "logs", serviceName), "1m")
	if err != nil {
		t.Fatal(err)
	}
	out, runErr := commands.RunAndWaitForOutput(cmd)
	if runErr != nil {
		t.Fatal(runErr)
	}
	return out
}

// helper function running 'docker-compose ps -q'
func composePsIds(t *testing.T, composeFileName, serviceName string) string {
	cmd, err := commands.NewCommand(
		compose(composeFileName, "ps", "-q", serviceName), "10s")
	if err != nil {
		t.Fatal(err)
	}
	out, runErr := commands.RunAndWaitForOutput(cmd)
	if runErr != nil {
		t.Fatal(runErr)
	}
	return out
}

// helper function to execute command via 'docker exec'
func dockerExec(t *testing.T, id string, execArgs ...string) string {
	args := []string{"docker", "exec", "-it", id}
	for _, arg := range execArgs {
		args = append(args, arg)
	}
	fmt.Println(args)
	cmd, err := commands.NewCommand(args, "10s")
	if err != nil {
		t.Fatal(err)
	}
	out, runErr := commands.RunAndWaitForOutput(cmd)
	if runErr != nil {
		t.Fatal(runErr)
	}
	return out
}
