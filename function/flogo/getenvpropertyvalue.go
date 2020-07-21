package customflogo

import (
	"fmt"
	"os"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&getEnvPropertyValueFn{})
}

type getEnvPropertyValueFn struct {
}

// Name returns the name of the function
func (getEnvPropertyValueFn) Name() string {
	return "getEnvPropertyValue"
}

// Sig returns the function signature
func (getEnvPropertyValueFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

var getEnvPropertyValueFnLogger = log.RootLogger()

// Eval executes the function
func (getEnvPropertyValueFn) Eval(params ...interface{}) (interface{}, error) {
	if getEnvPropertyValueFnLogger.DebugEnabled() {
		getEnvPropertyValueFnLogger.Debugf("Entering function getEnvPropertyValue (eval) with param: [%+v]", params[0])
	}

	inputParamValue := params[0]
	inputString := ""
	var outputValue interface{}
	var err error

	inputString, err = coerce.ToString(inputParamValue)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerece input value to a string. Value = [%+v].", inputParamValue)
	}

	if getEnvPropertyValueFnLogger.DebugEnabled() {
		getEnvPropertyValueFnLogger.Debugf("Input Parameter's string length is [%+v].", len(inputString))
	}

	if len(inputString) <= 0 {
		getEnvPropertyValueFnLogger.Debugf("Input Parameter is empty or nil. Returning nil.")
		return nil, nil
	}

	outputValue, exists := os.LookupEnv(inputString)
	if !exists {
		if getEnvPropertyValueFnLogger.DebugEnabled() {
			getEnvPropertyValueFnLogger.Debugf("failed to resolve Env Property: '%s', ensure that property is configured in the application", inputString)
		}
		return nil, nil
	}

	if getEnvPropertyValueFnLogger.DebugEnabled() {
		getEnvPropertyValueFnLogger.Debugf("Final output value = [%+v]", outputValue)
	}

	if getEnvPropertyValueFnLogger.DebugEnabled() {
		getEnvPropertyValueFnLogger.Debugf("Exiting function getEnvPropertyValue (eval)")
	}

	return outputValue, nil
}
