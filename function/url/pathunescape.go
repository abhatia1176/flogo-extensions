package customurl

import (
	"fmt"
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&pathUnescapeFn{})
}

type pathUnescapeFn struct {
}

// Name returns the name of the function
func (pathUnescapeFn) Name() string {
	return "pathUnescape"
}

// Sig returns the function signature
func (pathUnescapeFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

var pathUnescapeFnLogger = log.RootLogger()

/* Eval executes the function
* This function returns.
 */
func (pathUnescapeFn) Eval(params ...interface{}) (interface{}, error) {
	if pathUnescapeFnLogger.DebugEnabled() {
		pathUnescapeFnLogger.Debugf("Entering function pathUnescape (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	var outputValue string

	if inputParamValue == nil {
		//Do nothing
		if pathUnescapeFnLogger.DebugEnabled() {
			pathUnescapeFnLogger.Debugf("Input argument is nil. Will return nil as output.")
		}
		return nil, nil
	}

	inputString, ok := inputParamValue.(string)
	if !ok {
		if pathUnescapeFnLogger.DebugEnabled() {
			pathUnescapeFnLogger.Debugf("First argument is not a string. Argument Type is: %T. Will return error.", inputParamValue)
		}
		return nil, fmt.Errorf("first argument is not a string. Argument Type is: %T", inputParamValue)
	}

	outputValue, err := url.PathUnescape(inputString)
	if err != nil {
		if pathUnescapeFnLogger.DebugEnabled() {
			pathUnescapeFnLogger.Debugf("pathUnescape function returned an error = %s", err)
		}
		return nil, fmt.Errorf("pathUnescape function returned an error = %s", err)
	}
	if pathUnescapeFnLogger.DebugEnabled() {
		pathUnescapeFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if pathUnescapeFnLogger.DebugEnabled() {
		pathUnescapeFnLogger.Debugf("Exiting function pathUnescape (eval)")
	}

	return outputValue, nil
}
