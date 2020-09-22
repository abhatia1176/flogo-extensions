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

/* Eval executes the function
* This function returns sum of elements in a numeric array (if any value is not a number, an error is returned.
* The output of this function could be affected by rounding errors typical with floating point arithmetic.
* See some of the articles about this:
* What every Computer scientist should know about Floating point arithmetic - https://docs.oracle.com/cd/E19957-01/806-3568/ncg_goldberg.html
* Float rounding errors - https://stackoverflow.com/a/36055172/6385674
* Why not use double or float to represent currency - https://stackoverflow.com/questions/3730019/why-not-use-double-or-float-to-represent-currency
* Rounding off errors in Java - https://www.geeksforgeeks.org/rounding-off-errors-java/
*
* Following article provide information on various options which could help with the aforementioned errors:
* Use big.Float for higher precision - https://golang.org/src/math/big/example_test.go
* Another float conversion article   - https://sourcegraph.com/github.com/golang/go/-/blob/src/math/big/floatconv_test.go
* Improving accuracy of sum - http://alex.uwplse.org/2015/10/16/improving-accuracy-summation.html
* Tips for handling tricky floating point arithmetic - https://www.soa.org/news-and-publications/newsletters/compact/2014/may/com-2014-iss51/losing-my-precision-tips-for-handling-tricky-floating-point-arithmetic/
*
 */
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

	if inputArray == nil {
		//Do nothing
		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("Input arguments are nil or empty. Will return 0 as output.")
		}
		return 0.0, nil
	}

	outputValue = 0.0
	for k, v := range inputArray {

		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("[%+v]: Value at index [%+v] is [%+v], of type %T.", k, k, v, v)
			sumFnLogger.Debugf("[%+v]: Attempt to coerce the value to float64.", k)
		}

		tempOutputValue, err := coerce.ToFloat64(v)

		if err != nil {
			if sumFnLogger.DebugEnabled() {
				sumFnLogger.Debugf("[%+v]: Value at index [%+v] is [%+v], which is of type %T, and is not a number.", k, k, v, v)
				sumFnLogger.Debugf("[%+v]: Array is not an array of go number types. Cannot compute sum.")
			}
			return nil, fmt.Errorf("Value at index [%+v] is [%+v], which is of type %T, and cannot be coerced to float64. "+
				"Array is not an array of go number types. Cannot compute sum.", k, v, v)
		}
		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("[%+v]: Successfully coerced the value to float64.", k)
			sumFnLogger.Debugf("[%+v]: Coerced value is = [%+v]", k, tempOutputValue)
		}

		outputValue = outputValue + tempOutputValue

		if sumFnLogger.DebugEnabled() {
			sumFnLogger.Debugf("[%+v]: Current running sum value is [%+v].", k, outputValue)
		}
	}

	if sumFnLogger.DebugEnabled() {
		sumFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if sumFnLogger.DebugEnabled() {
		sumFnLogger.Debugf("Exiting function sum (eval)")
	}

	return outputValue, nil
}
