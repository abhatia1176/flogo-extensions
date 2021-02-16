package customdatetime

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&getTimeInNanoFn{})
}

type getTimeInNanoFn struct {
}

// Name returns the name of the function
func (getTimeInNanoFn) Name() string {
	return "getTimeInNano"
}

// Sig returns the function signature
func (getTimeInNanoFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

var getTimeInNanoFnLogger = log.RootLogger()

// Eval executes the function
func (getTimeInNanoFn) Eval(params ...interface{}) (interface{}, error) {
	if getTimeInNanoFnLogger.DebugEnabled() {
		getTimeInNanoFnLogger.Debugf("Entering function getTimeInNano (eval)")
	}

	outputTimeInNano := time.Now().UnixNano()

	if getTimeInNanoFnLogger.DebugEnabled() {
		getTimeInNanoFnLogger.Debugf("Output time in Nanoeconds is = %+v", outputTimeInNano)
	}

	if getTimeInNanoFnLogger.DebugEnabled() {
		getTimeInNanoFnLogger.Debugf("Exiting function getTimeInNano (eval)")
	}

	return outputTimeInNano, nil
}
