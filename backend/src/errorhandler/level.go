package errorhandler

type Level string

const (
	INFO    Level = "INFO"
	WARNING Level = "WARNING"
	ERROR   Level = "ERROR"
)

func (l Level) String() string {
	return string(l)
}
