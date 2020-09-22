package customarray

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var maxFnRef = &maxFn{}
var maxFnTestLogger log.Logger
var maxFnActualOutput interface{}
var maxFnExpectedOutput interface{}
var maxFnInput interface{}
var maxFnErr error

func init() {
	maxFnTestLogger = log.RootLogger()
	log.SetLogLevel(maxFnTestLogger, log.DebugLevel)
}

//sunny path case
func Test_max_1(t *testing.T) {
	data := `[0,1,2,3,4,5,6,7]`
	err := json.Unmarshal([]byte(data), &maxFnInput)
	maxFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	maxFnExpectedOutput = 7

	maxFnActualOutput, maxFnErr = maxFnRef.Eval(maxFnInput)

	maxFnTestLogger.Debug("In tester: Output of function call = ", maxFnActualOutput)
	assert.Nil(t, maxFnErr)
	assert.EqualValues(t, maxFnExpectedOutput, maxFnActualOutput)
}

//sunny path - numeric string array with non-integer number
func Test_max_2(t *testing.T) {
	data := `["0","7.8","2","3","4","5","6","7.7"]`
	err := json.Unmarshal([]byte(data), &maxFnInput)
	maxFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	maxFnExpectedOutput = "7.8"

	maxFnActualOutput, maxFnErr = maxFnRef.Eval(maxFnInput)

	maxFnTestLogger.Debug("In tester: Output of function call = ", maxFnActualOutput)
	assert.Nil(t, maxFnErr)
	assert.EqualValues(t, maxFnExpectedOutput, maxFnActualOutput)
}

//error test -- non-numeric element in array
func Test_max_3(t *testing.T) {
	data := `["0","asdasd","2","3","asd","5","6","7"]`
	err := json.Unmarshal([]byte(data), &maxFnInput)
	maxFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	maxFnExpectedOutput = nil

	maxFnActualOutput, maxFnErr = maxFnRef.Eval(maxFnInput)

	maxFnTestLogger.Debug("In tester: Output of function call = ", maxFnActualOutput)
	maxFnTestLogger.Debug("In tester: Error from function call = ", maxFnErr)
	assert.NotNil(t, maxFnErr)
}

//empty array
func Test_max_4(t *testing.T) {
	data := `[]`
	err := json.Unmarshal([]byte(data), &maxFnInput)
	maxFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	maxFnExpectedOutput = nil

	maxFnActualOutput, maxFnErr = maxFnRef.Eval(maxFnInput)

	maxFnTestLogger.Debug("In tester: Output of function call = ", maxFnActualOutput)
	assert.Nil(t, maxFnErr)
	assert.EqualValues(t, maxFnExpectedOutput, maxFnActualOutput)
}

func Test_max_5(t *testing.T) {
	data := `[-21,1433,2,3,4223,5,6,7]`
	err := json.Unmarshal([]byte(data), &maxFnInput)
	maxFnTestLogger.Debug("In tester: Unmarshaled status = ", err != nil)

	maxFnExpectedOutput = 4223

	maxFnActualOutput, maxFnErr = maxFnRef.Eval(maxFnInput)

	maxFnTestLogger.Debug("In tester: Output of function call = ", maxFnActualOutput)
	assert.Nil(t, maxFnErr)
	assert.EqualValues(t, maxFnExpectedOutput, maxFnActualOutput)
}
