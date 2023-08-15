package failure

import "fmt"

type FailureType string

type Failure struct {
	Type FailureType
	Args map[string]interface{}
}

func New(t FailureType) *Failure {
	return &Failure{
		Type: t,
	}
}

func NewWithArgs(t FailureType, a map[string]interface{}) *Failure {
	return &Failure{
		Type: t,
		Args: a,
	}
}

func (e *Failure) Error() string {
	return fmt.Sprintf("%v - %#v", e.Type, e.Args)
}
