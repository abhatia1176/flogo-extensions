package customflogo

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var getAppPropertyValueFnRef = &getAppPropertyValueFn{}
var getAppPropertyValueFnTestLogger log.Logger
var appPropValueActualOutput interface{}
var appPropValueExpectedOutput interface{}
var appPropValueInput interface{}
var appPropValueError error

func init() {
	getAppPropertyValueFnTestLogger = log.RootLogger()
	log.SetLogLevel(getAppPropertyValueFnTestLogger, log.DebugLevel)
}

func Test_getAppPropertyValue_1(t *testing.T) {
	appPropValueInput = "TEST"
	appPropValueExpectedOutput = nil

	appPropValueActualOutput, appPropValueError = getAppPropertyValueFnRef.Eval(appPropValueInput)

	getAppPropertyValueFnTestLogger.Debug("In tester: Output of function call = ", appPropValueExpectedOutput)
	assert.Nil(t, appPropValueError)
	assert.EqualValues(t, appPropValueExpectedOutput, appPropValueActualOutput)
}
