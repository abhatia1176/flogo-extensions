package customdatetime

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var getCurrentTimestampFnRef = &getCurrentTimestampFn{}
var getCurrentTimestampFnTestLogger log.Logger
var getCurrentTimestampActualOutput interface{}
var getCurrentTimestampExpectedOutput interface{}
var getCurrentTimestampInput1 interface{}
var getCurrentTimestampInput2 interface{}
var getCurrentTimestampInput3 interface{}
var getCurrentTimestampError error

func init() {
	getCurrentTimestampFnTestLogger = log.RootLogger()
	log.SetLogLevel(getCurrentTimestampFnTestLogger, log.DebugLevel)
}

func Test_getCurrentTimestamp_1(t *testing.T) {
	getCurrentTimestampInput1 = "yyyy-MM-ddThh:mm:ss"
	getCurrentTimestampInput2 = "ms"
	getCurrentTimestampInput3 = "Z"
	getCurrentTimestampExpectedOutput = nil

	getCurrentTimestampActualOutput, getCurrentTimestampError = getCurrentTimestampFnRef.Eval(getCurrentTimestampInput1, getCurrentTimestampInput2, getCurrentTimestampInput3)

	getCurrentTimestampFnTestLogger.Debug("In tester: Output of function call = ", getCurrentTimestampExpectedOutput)
	assert.Nil(t, getCurrentTimestampError)
	//assert.EqualValues(t, getCurrentTimestampExpectedOutput, getCurrentTimestampActualOutput)
}

func Test_getCurrentTimestamp_2(t *testing.T) {
	getCurrentTimestampInput1 = "yyyy-MM-ddThh:mm:ss"
	getCurrentTimestampInput2 = "ns"
	getCurrentTimestampInput3 = "Z-0000"
	getCurrentTimestampExpectedOutput = nil

	getCurrentTimestampActualOutput, getCurrentTimestampError = getCurrentTimestampFnRef.Eval(getCurrentTimestampInput1, getCurrentTimestampInput2, getCurrentTimestampInput3)

	getCurrentTimestampFnTestLogger.Debug("In tester: Output of function call = ", getCurrentTimestampExpectedOutput)
	assert.Nil(t, getCurrentTimestampError)
	//assert.EqualValues(t, getCurrentTimestampExpectedOutput, getCurrentTimestampActualOutput)
}
