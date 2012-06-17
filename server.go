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

// breaks the reference, allowing the object to be disposed of
func Free(proc *SProcess) {

	// close any open files
	for _, file := range proc.files {
		file.Close()
	}

	// delete the directory
	os.RemoveAll("/system/process/" + strconv.Itoa(proc.pid))

	delete(processes, proc.pid)
}

// PID getter
func (proc *SProcess) PID() int {
	return proc.pid
}

// returns true if proc has property prop
func (proc *SProcess) HasProperty(prop string) bool {

	// first, check if a File is open.
	if proc.files[prop] != nil {
		return true
	}

	// otherwise, do a dirty check and see if the file exists.
	_, err := os.Lstat("/system/process/" + strconv.Itoa(proc.pid))
	if err != nil {
		return false
	}

	return true
}

// returns string property prop
func (proc *SProcess) GetProperty(prop string) string {
	return "(undefined)"
}

// assign a property
func (proc *SProcess) SetProperty(prop string, value string) {

}
