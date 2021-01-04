package mapreader

func (err *Error) Error() string {
	return err.Error()
}

type ErrorWrapUnion interface {
	PkgmapreaderError()
	error
}

func ErrorWrap(err ErrorWrapUnion) *Error {
	return &Error{Err: err}
}
func (err *Error) Unwrap() ErrorWrapUnion {
	return err.Err.(ErrorWrapUnion)
}
func (err *Error) PkgptraceTraceError() {}
