package customdatetime

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&getTimeInMillisFn{})
}

type getTimeInMillisFn struct {
}

// Name returns the name of the function
func (getTimeInMillisFn) Name() string {
	return "getTimeInMillis"
}

// Sig returns the function signature
func (getTimeInMillisFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

var getTimeInMillisFnLogger = log.RootLogger()

// Eval executes the function
func (getTimeInMillisFn) Eval(params ...interface{}) (interface{}, error) {
	if getTimeInMillisFnLogger.DebugEnabled() {
		getTimeInMillisFnLogger.Debugf("Entering function getTimeInMillis (eval)")
	}

	outputTimeInMillis := time.Now().UnixNano() / int64(time.Millisecond)

	if getTimeInMillisFnLogger.DebugEnabled() {
		getTimeInMillisFnLogger.Debugf("Output time in milliseconds is = %+v", outputTimeInMillis)
	}

	if getTimeInMillisFnLogger.DebugEnabled() {
		getTimeInMillisFnLogger.Debugf("Exiting function getTimeInMillis (eval)")
	}

	return outputTimeInMillis, nil
}
