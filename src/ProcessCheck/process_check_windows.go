package ProcessCheck

import (
	"errors"
	"golang.org/x/sys/windows"
)

type ProcessCheckerWindows struct{}

func (p ProcessCheckerWindows) ProcessExists(pid int) (bool, error) {
	h, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, uint32(pid))
	if err != nil {
		if err == windows.ERROR_INVALID_PARAMETER {
			return false, errors.New("The process does not exist")
		}
		return false, err
	}
	// Close the process handle to avoid a leak
	windows.CloseHandle(h)

	return true, err
}

func NewProcessChecker() ProcessChecker {
	return ProcessCheckerWindows{}
}
