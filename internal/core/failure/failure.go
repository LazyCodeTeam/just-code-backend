package failure

import "fmt"

type FailureCode string

type FailureGroup int

const (
	FailureGroupUnknown FailureGroup = iota
	FailureGroupNotFound
	FailureGroupInput
	FailureGroupState
	FailureGroupAuth
)

type Failure struct {
	Group       FailureGroup
	Code        FailureCode
	ParentError error
	Args        map[string]interface{}
}

func NewAuthFailure(code FailureCode, rest ...interface{}) *Failure {
	return New(FailureGroupAuth, code, rest...)
}

func NewInputFailure(code FailureCode, rest ...interface{}) *Failure {
	return New(FailureGroupInput, code, rest...)
}

func NewStateFailure(code FailureCode, rest ...interface{}) *Failure {
	return New(FailureGroupState, code, rest...)
}

func NewNotFoundFailure(code FailureCode, rest ...interface{}) *Failure {
	return New(FailureGroupNotFound, code, rest...)
}

func NewUnknownFailure(code FailureCode, rest ...interface{}) *Failure {
	return New(FailureGroupUnknown, code, rest...)
}

func New(group FailureGroup, code FailureCode, rest ...interface{}) *Failure {
	if len(rest) == 0 {
		return &Failure{
			Group: group,
			Code:  code,
		}
	}

	var parentError error

	if err, ok := rest[0].(error); ok {
		parentError = err
		rest = rest[1:]
	}

	if len(rest) == 0 {
		return &Failure{
			Group:       group,
			Code:        code,
			ParentError: parentError,
		}
	}

	if len(rest)%2 != 0 {
		return &Failure{
			Group:       group,
			Code:        code,
			ParentError: parentError,
			Args: map[string]interface{}{
				"unknown": rest,
			},
		}
	}
	args := make(map[string]interface{})
	for i := 0; i < len(rest); i += 2 {
		key, ok := rest[i].(string)
		if !ok {
			unknownKey := fmt.Sprintf("unknown_%d", i/2)
			args[unknownKey] = rest[i]

			continue
		}
		args[key] = rest[i+1]
	}

	return &Failure{
		Group:       group,
		Code:        code,
		ParentError: parentError,
		Args:        args,
	}
}

func (f *Failure) Unwrap() error {
	return f.ParentError
}

func (e *Failure) Error() string {
	return fmt.Sprintf(
		"%s: %s\nParentError: %v\nArgs: %v",
		failureGroupToString(e.Group),
		e.Code,
		e.ParentError,
		e.Args,
	)
}

func failureGroupToString(group FailureGroup) string {
	switch group {
	case FailureGroupUnknown:
		return "Unknown"
	case FailureGroupNotFound:
		return "NotFound"
	case FailureGroupInput:
		return "Input"
	case FailureGroupState:
		return "State"
	default:
		return "Unknown"
	}
}
