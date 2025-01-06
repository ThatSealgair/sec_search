package tui

type state int

const (
	stateSelectFromPlatform state = iota
	stateSelectToPlatform
	stateEnterQuery
	stateShowResult
)

func (s state) String() string {
	switch s {
	case stateSelectFromPlatform:
		return "Select Source Platform"
	case stateSelectToPlatform:
		return "Select Target Platform"
	case stateEnterQuery:
		return "Enter Query"
	case stateShowResult:
		return "Show Result"
	default:
		return "Unknown State"
	}
}
