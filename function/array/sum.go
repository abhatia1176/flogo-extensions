package customarray

import (
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&sumFn{})
}

type sumFn struct {
}

// Name returns the name of the function
func (sumFn) Name() string {
	return "sum"
}

// Sig returns the function signature
func (sumFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

var sumFnLogger = log.RootLogger()

// Eval executes the function
// This function returns true, if the input value is either empty or nil  (json null)
func (sumFn) Eval(params ...interface{}) (interface{}, error) {
	if sumFnLogger.DebugEnabled() {
		sumFnLogger.Debugf("Entering function sum (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	var outputValue float64

	inputArray, ok := inputParamValue.([]interface{})
	if !ok {
		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("First argument is not an array. Argument Type is: %T. Will return error.", inputParamValue)
		}
		return nil, fmt.Errorf("First argument is not an array. Argument Type is: %T", inputParamValue)
	}

	if inputArray == nil  {
		//Do nothing
		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("Input arguments are nil or empty. Will return 0 output.")
		}
		return 0, nil
	}

	counter:=0
	outputValue = 0
	for k, v := range inputArray {
		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("[%+v]: array value-%+v : %+v", counter, k, v)
		}
		switch v.(type) {
			case int, float64, float32, int8, int32, int64, int16, uint, uint8, uint16, uint32, uint64, string:
				tempOutputValue, err := coerce.ToFloat64(v)
				if (err != nil) {
					return nil, fmt.Errorf("Value at index [%+v] is [%+v], which is of type %T, and cannot be coerced to float64.", counter, v, v)
				}
				outputValue = outputValue + tempOutputValue
			default:
				if sumFnLogger.DebugEnabled() {
					sumFnLogger.Debugf("Array is not an array of go number types. Cannot compute sum.")
					sumFnLogger.Debugf("Value at index [%+v] is [%+v], which is of type %T, and is not a number.", counter, v, v)

				}
				return nil, fmt.Errorf("Array is not an array of go number types. Cannot compute sum.")
		}
		counter++
	}

	if sumFnLogger.DebugEnabled() {
		sumFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if sumFnLogger.DebugEnabled() {
		sumFnLogger.Debugf("Exiting function sum (eval)")
	}

	return outputValue, nil
}