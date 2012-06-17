package process

import (
	"io"
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

// returns string property prop
func (proc *CProcess) GetProperty(prop string) string {
	file, err := os.Open("/system/process/" + strconv.Itoa(proc.pid) + "/" + prop)
	if err != nil {
		return "(undefined)"
	}

	// read up to 1024 bytes
	b := make([]byte, 1024)
	_, err = file.Read(b)

	// an error occured, and it was not an EOF
	if err != nil && err != io.EOF {
		return "(undefined)"
	}

	// file was more than 1M
	if err != io.EOF {
		return "(maxed out)"
	}

	file.Close()
	return string(b)
}
