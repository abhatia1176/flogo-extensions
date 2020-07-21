package customflogo

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/data/property"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&getAppPropertyValueFn{})
}

type getAppPropertyValueFn struct {
}

// Name returns the name of the function
func (getAppPropertyValueFn) Name() string {
	return "getAppPropertyValue"
}

// Sig returns the function signature
func (getAppPropertyValueFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

var getAppPropertyValueFnLogger = log.RootLogger()

// Eval executes the function
func (getAppPropertyValueFn) Eval(params ...interface{}) (interface{}, error) {
	if getAppPropertyValueFnLogger.DebugEnabled() {
		getAppPropertyValueFnLogger.Debugf("Entering function getAppPropertyValue (eval) with param: [%+v]", params[0])
	}

	inputParamValue := params[0]
	inputString := ""
	var outputValue interface{}
	var err error

	inputString, err = coerce.ToString(inputParamValue)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerece input value to a string. value = %+v.", inputParamValue)
	}

	if getAppPropertyValueFnLogger.DebugEnabled() {
		getAppPropertyValueFnLogger.Debugf("Input Parameter's string length is %+v.", len(inputString))
	}

	if len(inputString) <= 0 {
		getAppPropertyValueFnLogger.Debugf("Input Parameter is empty or nil. Returning nil.")
		return nil, nil
	}

	manager := property.DefaultManager()
	outputValue, exists := manager.GetProperty(inputString) //should we add the path and reset it to ""

	if !exists {
		if getAppPropertyValueFnLogger.DebugEnabled() {
			getAppPropertyValueFnLogger.Debugf("failed to resolve Property: '%s', ensure that property is configured in the application", inputString)
		}
		return nil, nil
	}
	if getAppPropertyValueFnLogger.DebugEnabled() {
		getAppPropertyValueFnLogger.Debugf("Final output value = [%+v]", outputValue)
	}

	if getAppPropertyValueFnLogger.DebugEnabled() {
		getAppPropertyValueFnLogger.Debugf("Exiting function getAppPropertyValue (eval)")
	}

	return outputValue, nil
}
