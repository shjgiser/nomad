// +build !linux

package driver

import (
	"os/exec"
	"syscall"

	cgroupConfig "github.com/opencontainers/runc/libcontainer/configs"
)

// isolateCommand sets the setsid flag in exec.Cmd to true so that the process
// becomes the process leader in a new session and doesn't receive signals that
// are sent to the parent process.
func isolateCommand(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Setsid = true
}

func destroyCgroup(group *cgroupConfig.Cgroup) error {
	return nil
}
