package ProcessCheck

import "syscall"

type ProcessCheckerLinux struct{}

func (p ProcessCheckerLinux) ProcessExists(pid int) (bool, error) {
	// source: https://docs.oracle.com/cd/E86824_01/html/E54765/kill-2.html
	// The kill() function sends a signal to a process or a group of processes. The process or group of processes to which the signal is to be sent is specified by pid. The signal that is to be sent is specified by sig and is either one from the list given in signal (see signal.h(3HEAD)), or 0. If sig is 0 (the null signal), error checking is performed but no signal is actually sent. This can be used to check the validity of pid.
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
	return ProcessCheckerLinux{}
}
