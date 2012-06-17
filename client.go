package process

import (
	"os"
	"strconv"
)

type CProcess struct {
	pid int
}

// creates a new client process
func CFromPID(pid int) *CProcess {
	return &CProcess{pid}
}

// PID getter
func (proc *CProcess) PID() int {
	return proc.pid
}

// returns true if process has property prop
func (proc *CProcess) HasProperty(prop string) bool {

	// if the file exists, it has the property.
	_, err := os.Lstat("/system/process/" + strconv.Itoa(proc.pid))
	if err != nil {
		return false
	}
	return true
}
