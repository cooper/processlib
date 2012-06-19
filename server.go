package process

import (
	"io"
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
		pid:   pid,
		files: make(map[string]*os.File),
	}
	processes[pid] = proc
	os.Mkdir("/system/process/"+strconv.Itoa(pid), 0755)

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
	var (
		file *os.File
		err  error
	)

	// file exists
	if proc.files[prop] != nil {
		file = proc.files[prop]
		file.Seek(0, 0)

		// doesn't exist; create
	} else {
		file, err = os.Create("/system/process/" + strconv.Itoa(proc.pid) + "/" + prop)
		file.Chmod(0755)
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

	return string(b)
}

// assign a property
func (proc *SProcess) SetProperty(prop string, value string) {
	var file *os.File

	// file exists; empty
	if proc.files[prop] != nil {
		file = proc.files[prop]
		file.Truncate(0)

		// doesn't exist; create
	} else {
		file, _ = os.Create("/system/process/" + strconv.Itoa(proc.pid) + "/" + prop)
		file.Chmod(0755)
	}

	proc.files[prop] = file

	// write
	file.Seek(0, 0)
	file.WriteString(value)
}
