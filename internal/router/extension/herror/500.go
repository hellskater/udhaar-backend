package herror

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

// InternalError Internal error
type InternalError struct {
	// Err error
	Err error
	// Stack Stack trace
	Stack []byte
	// Fields ZAP log field
	Fields []zap.Field
	// Panic Whether Panic occurred
	Panic bool
}

func (i *InternalError) Error() string {
	if i.Panic {
		return fmt.Sprintf("[Panic] %s\n%s", i.Err.Error(), i.Stack)
	}
	return fmt.Sprintf("%s\n%s", i.Err.Error(), i.Stack)
}

func InternalServerError(err error) error {
	return &InternalError{
		Err:    err,
		Stack:  debug.Stack(),
		Fields: []zap.Field{zapdriver.ErrorReport(runtime.Caller(1)), zap.Error(err)},
		Panic:  false,
	}
}

func Panic(err error) error {
	return &InternalError{
		Err:    err,
		Stack:  debug.Stack(),
		Fields: []zap.Field{zapdriver.ErrorReport(runtime.Caller(1)), zap.Error(err)},
		Panic:  true,
	}
}
