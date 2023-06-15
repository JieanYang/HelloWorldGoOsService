package ProcessCheck

import "syscall"

type ProcessCheckerDarwin struct{}

func (p ProcessCheckerDarwin) ProcessExists(pid int) (bool, error) {
	err := syscall.Kill(pid, 0)
	if err == nil {
		return true, nil
	}
	if err == syscall.ESRCH {
		return false, nil
	}
	return false, err
}

func NewProcessChecker() ProcessChecker {
	return ProcessCheckerDarwin{}
}
