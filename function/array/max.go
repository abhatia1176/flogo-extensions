package customarray

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&maxFn{})
}

type maxFn struct {
}

// Name returns the name of the function
func (maxFn) Name() string {
	return "max"
}

// Sig returns the function signature
func (maxFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

var maxFnLogger = log.RootLogger()

/* Eval executes the function
* This function returns max of elements in a numeric array (if any value is not a number, an error is returned.)
*
 */
func (maxFn) Eval(params ...interface{}) (interface{}, error) {
	if maxFnLogger.DebugEnabled() {
		maxFnLogger.Debugf("Entering function max (eval) with param: %+v", params[0])
	}

	inputParamValue := params[0]
	var outputValue interface{}

	inputArray, ok := inputParamValue.([]interface{})
	if !ok {
		if maxFnLogger.DebugEnabled() {
			maxFnLogger.Debugf("First argument is not an array. Argument Type is: %T. Will return error.", inputParamValue)
		}
		return nil, fmt.Errorf("First argument is not an array. Argument Type is: %T", inputParamValue)
	}

	if inputArray == nil || len(inputArray) == 0 {
		//Do nothing
		if maxFnLogger.DebugEnabled() {
			maxFnLogger.Debugf("Input arguments are nil or empty. Will return nil as output.")
		}
		return nil, nil
	}

	maxValue := 0.0
	indexForMaxValue := 0

	for j := 0; j < len(inputArray); j++ {
		valueToCompare := inputArray[j]

		if maxFnLogger.DebugEnabled() {
			maxFnLogger.Debugf("[%+v]: Value at index [%+v] is [%+v], of type %T.", j, j, valueToCompare, valueToCompare)
			maxFnLogger.Debugf("[%+v]: Attempt to coerce the value to float64.", j)
		}

		tempValueToCompare, err := coerce.ToFloat64(valueToCompare)

		if err != nil {
			if maxFnLogger.DebugEnabled() {
				maxFnLogger.Debugf("[%+v]: Value at index [%+v] is [%+v], which is of type %T, and is not a number.", j, j, valueToCompare, valueToCompare)
				maxFnLogger.Debugf("[%+v]: Array is not an array of go number types. Cannot compute max.")
			}
			return nil, fmt.Errorf("Value at index [%+v] is [%+v], which is of type %T, and cannot be coerced to float64. "+
				"Array is not an array of go number types. Cannot compute max.", j, valueToCompare, valueToCompare)
		}

		if maxFnLogger.DebugEnabled() {
			maxFnLogger.Debugf("[%+v]: Successfully coerced the value to float64.", j)
			maxFnLogger.Debugf("[%+v]: Coerced value is = [%+v]", j, tempValueToCompare)
		}

		if j == 0 {
			maxValue = tempValueToCompare
		} else {
			if maxValue < tempValueToCompare {
				maxValue = tempValueToCompare
				indexForMaxValue = j
			}
		}

		if maxFnLogger.DebugEnabled() {
			maxFnLogger.Debugf("[%+v]: Current max value is [%+v].", j, maxValue)
			maxFnLogger.Debugf("[%+v]: Current max index is [%+v].", j, indexForMaxValue)
		}

	}

	outputValue = inputArray[indexForMaxValue]

	if maxFnLogger.DebugEnabled() {
		maxFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if maxFnLogger.DebugEnabled() {
		maxFnLogger.Debugf("Exiting function max (eval)")
	}

	return outputValue, nil
}
