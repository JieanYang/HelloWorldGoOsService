package ProcessCheck

type ProcessChecker interface {
	ProcessExists(pid int) (bool, error)
}
