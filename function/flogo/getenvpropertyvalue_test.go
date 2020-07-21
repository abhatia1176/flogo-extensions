package customflogo

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var getEnvPropertyValueFnRef = &getEnvPropertyValueFn{}
var getEnvPropertyValueFnTestLogger log.Logger
var envPropValueActualOutput interface{}
var envPropValueExpectedOutput interface{}
var envPropValueInput interface{}
var envPropValueError error

func init() {
	getEnvPropertyValueFnTestLogger = log.RootLogger()
	log.SetLogLevel(getEnvPropertyValueFnTestLogger, log.DebugLevel)
}

func Test_returnDefaultIfEmpty_1(t *testing.T) {
	envPropValueInput = "TEST"
	envPropValueExpectedOutput = nil

	envPropValueActualOutput, envPropValueError = getEnvPropertyValueFnRef.Eval(envPropValueInput)

	getEnvPropertyValueFnTestLogger.Debug("In tester: Output of function call = ", envPropValueExpectedOutput)
	assert.Nil(t, envPropValueError)
	assert.EqualValues(t, envPropValueExpectedOutput, envPropValueActualOutput)
}
