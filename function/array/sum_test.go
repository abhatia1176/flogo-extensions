package customarray

import (
	"encoding/json"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sumFnRef = &sumFn{}
var sumFnTestLogger log.Logger
var sumFnActualOutput interface{}
var sumFnExpectedOutput interface{}
var sumFnInput interface{}
var sumFnErr error

func init() {
	sumFnTestLogger = log.RootLogger()
	log.SetLogLevel(sumFnTestLogger, log.DebugLevel)
}

//sunny path case
func Test_sum_1(t *testing.T) {
	data := `[0,1,2,3,4,5,6,7]`
	err := json.Unmarshal([]byte(data), &sumFnInput)
	sumFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	sumFnExpectedOutput = 28

	sumFnActualOutput, sumFnErr = sumFnRef.Eval(sumFnInput)

	sumFnTestLogger.Debug("In tester: Output of function call = ", sumFnActualOutput)
	assert.Nil(t, sumFnErr)
	assert.EqualValues(t, sumFnExpectedOutput, sumFnActualOutput)
}

//sunny path - numeric string array with non-integer number
func Test_sum_2(t *testing.T) {
	data := `["0","1","2","3","4","5","6","7.7"]`
	err := json.Unmarshal([]byte(data), &sumFnInput)
	sumFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	sumFnExpectedOutput = 28.7

	sumFnActualOutput, sumFnErr = sumFnRef.Eval(sumFnInput)

	sumFnTestLogger.Debug("In tester: Output of function call = ", sumFnActualOutput)
	assert.Nil(t, sumFnErr)
	assert.EqualValues(t, sumFnExpectedOutput, sumFnActualOutput)
}

//error test -- non-numeric element in array
func Test_sum_3(t *testing.T) {
	data := `["0","asdasd","2","3","asd","5","6","7"]`
	err := json.Unmarshal([]byte(data), &sumFnInput)
	sumFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	sumFnExpectedOutput = nil

	sumFnActualOutput, sumFnErr = sumFnRef.Eval(sumFnInput)

	sumFnTestLogger.Debug("In tester: Output of function call = ", sumFnActualOutput)
	sumFnTestLogger.Debug("In tester: Error from function call = ", sumFnErr)
	assert.NotNil(t, sumFnErr)
}

//empty array
func Test_sum_4(t *testing.T) {
	data := `[]`
	err := json.Unmarshal([]byte(data), &sumFnInput)
	sumFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	sumFnExpectedOutput = 0

	sumFnActualOutput, sumFnErr = sumFnRef.Eval(sumFnInput)

	sumFnTestLogger.Debug("In tester: Output of function call = ", sumFnActualOutput)
	assert.Nil(t, sumFnErr)
	assert.EqualValues(t, sumFnExpectedOutput, sumFnActualOutput)
}

//Test json.Number type
// Reference: https://github.com/golang/go/wiki/InterfaceSlice
// https://eager.io/blog/go-and-json/
// https://gobyexample.com/json
// https://yourbasic.org/golang/json-example/
type testStruct struct {
	Price []json.Number
}

func Test_sum_5(t *testing.T) {

	//Create a struct with array of json.Number values
	data := `{"Price": [1, 2.6, -0.3]}`
	res := testStruct{}
	err := json.Unmarshal([]byte(data), &res)

	sumFnTestLogger.Debug("In tester: Unmarshal status = ", err != nil)
	sumFnTestLogger.Debug("In tester: Unmarshal data to struct = ", res)
	sumFnTestLogger.Debugf("In tester: %+v %T", res.Price[0], res.Price[0])

	sumFnTestLogger.Debugf("In tester: %+v %T", res.Price, res.Price)

	//Extract json.Number array and assign to function input
	var interfaceSlice []interface{} = make([]interface{}, len(res.Price))
	for i, d := range res.Price {
		interfaceSlice[i] = d
	}
	sumFnTestLogger.Debugf("In tester: %+v %T", interfaceSlice, interfaceSlice)

	sumFnInput = interfaceSlice

	sumFnExpectedOutput = 3.4

	sumFnActualOutput, sumFnErr = sumFnRef.Eval(sumFnInput)

	sumFnTestLogger.Debug("In tester: Output of function call = ", sumFnActualOutput)
	assert.Nil(t, sumFnErr)
	assert.EqualValues(t, sumFnExpectedOutput, sumFnActualOutput)
}
