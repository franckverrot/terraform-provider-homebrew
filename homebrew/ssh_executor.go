package homebrew

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// SSHExecutor stores the configuration
type SSHExecutor struct {
	Config *Config
}

// CommandResult stores the result of a command execution
type CommandResult struct {
	Args   string
	Output string
}

// Command executes a command through SSH
func (s *SSHExecutor) Command(args ...string) (*CommandResult, error) {
	arg := []string{}

	// Login is optional so we want to make sure we forge a valid command
	login := ""
	if s.Config.Login != "" {
		login = fmt.Sprintf("%s@", s.Config.Login)
	}
	arg = append(arg, fmt.Sprintf("%s%s", login, s.Config.Host))

	arg = append(arg, s.Config.HomebrewBinaryPath)
	arg = append(arg, args...)

	combinedArgs := strings.Join(arg, " ")

	cmd := exec.Command("ssh", arg...)
	stdout, err := cmd.CombinedOutput()

	if stdout != nil {
		output := string(stdout)
		if strings.Contains(output, "is already installed") {
			return &CommandResult{Args: combinedArgs, Output: output}, nil
		}
	}

	fullCommand := fmt.Sprintf("ssh %s", combinedArgs)

	if err != nil {
		return nil, fmt.Errorf("error waiting for %+v: exit code %+v, command\n%s", cmd, err, fullCommand)
	}

	if stdout == nil {
		return nil, errors.New("weird")
	}
	output := string(stdout)
	return &CommandResult{Args: combinedArgs, Output: output}, nil
}
