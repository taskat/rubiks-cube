package errorhandler

type errorhandler struct {
	errors []Error
}

func newHandler() errorhandler {
	return errorhandler{errors: make([]Error, 0)}
}

func (e *errorhandler) addError(err Error) {
	e.errors = append(e.errors, err)
}

var handler = newHandler()

func AddError(err Error) {
	handler.addError(err)
}

func GetErrors() []Error {
	return handler.errors
}
