package process

type Process struct {
	pid int
}

func FromPID(pid int) *Process {
	return &Process{pid}
}
