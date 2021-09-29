package customurl

import (
	"fmt"
	"net/url"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&queryUnescapeFn{})
}

type queryUnescapeFn struct {
}

// Name returns the name of the function
func (queryUnescapeFn) Name() string {
	return "queryUnescape"
}

// Sig returns the function signature
func (queryUnescapeFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

var queryUnescapeFnLogger = log.RootLogger()

/* Eval executes the function
* This function returns.
 */
func (queryUnescapeFn) Eval(params ...interface{}) (interface{}, error) {
	if queryUnescapeFnLogger.DebugEnabled() {
		queryUnescapeFnLogger.Debugf("Entering function queryUnescape (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	var outputValue string

	if inputParamValue == nil {
		//Do nothing
		if queryUnescapeFnLogger.DebugEnabled() {
			queryUnescapeFnLogger.Debugf("Input argument is nil. Will return nil as output.")
		}
		return nil, nil
	}

	inputString, ok := inputParamValue.(string)
	if !ok {
		if queryUnescapeFnLogger.DebugEnabled() {
			queryUnescapeFnLogger.Debugf("First argument is not a string. Argument Type is: %T. Will return error.", inputParamValue)
		}
		return nil, fmt.Errorf("first argument is not a string. Argument Type is: %T", inputParamValue)
	}

	outputValue, err := url.QueryUnescape(inputString)
	if err != nil {
		if queryUnescapeFnLogger.DebugEnabled() {
			queryUnescapeFnLogger.Debugf("queryUnescape function returned an error = %s", err)
		}
		return nil, fmt.Errorf("queryUnescape function returned an error = %s", err)
	}
	if queryUnescapeFnLogger.DebugEnabled() {
		queryUnescapeFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if queryUnescapeFnLogger.DebugEnabled() {
		queryUnescapeFnLogger.Debugf("Exiting function queryUnescape (eval)")
	}

	return outputValue, nil
}
