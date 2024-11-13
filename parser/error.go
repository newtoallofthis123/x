package parser

type ParserError struct {
	Line string
	Msg  string
}

func (e *ParserError) Error() string {
	return e.Msg
}
