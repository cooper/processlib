package process

import (
	"os"
	"strconv"
)

var processes map[int]*SProcess

type SProcess struct {
	pid   int
	files map[string]*os.File
}

func SFromPID(pid int) *SProcess {
	if processes == nil {
		processes = make(map[int]*SProcess)
	}

	// already exists
	if processes[pid] != nil {
		return processes[pid]
	}

	// create new
	proc := &SProcess{
		pid: pid,
	}
	processes[pid] = proc
	os.Mkdir("/system/process/"+strconv.Itoa(pid), os.ModeDir)

	return proc
}

func Free(proc *SProcess) {

	// close any open files
	for _, file := range proc.files {
		file.Close()
	}

	// delete the directory
	os.RemoveAll("/system/process/" + strconv.Itoa(proc.pid))

	delete(processes, proc.pid)
}

func (proc *SProcess) PID() int {
	return proc.pid
}

func (proc *CProcess) HasProperty(prop string) bool {
	return false
}
