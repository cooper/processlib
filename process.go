package process

/*

	This is the generic process interface.
	Both server process objects and client process objects are required to implement these methods.

	Server process objects (type *SProcess) are used by ProcessManager itself. They open and write
	to files in /system/process for storing information about each process. It will also create and
	delete directories as needed.

	Client process objects (type *CProcess) are used by other applications. They open and read files
	in /system/process to retrieve information about each process. Unlike SProcess, CProcess will
	immediately close the file after it has read from it. For this reason, objects of type CProcess
	do not have to be Free()'d like those of SProcess do.

	This interface is only for specification; in most production uses, a definite type of either
	CProcess or SProcess will be specified.

*/

type Process interface {

	// returns true if a process has a property
	HasProperty(prop string) bool

	// returns the PID of the process
	PID() int

	// returns a string property.
	// it should not bail but instead return a string error
	GetProperty(prop string) string
}
