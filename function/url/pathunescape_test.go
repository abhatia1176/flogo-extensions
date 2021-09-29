package customurl

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var pathUnescapeFnRef = &pathUnescapeFn{}
var pathUnescapeFnTestLogger log.Logger
var pathUnescapeFnActualOutput interface{}
var pathUnescapeFnExpectedOutput interface{}
var pathUnescapeFnInput interface{}
var pathUnescapeFnErr error

func init() {
	pathUnescapeFnTestLogger = log.RootLogger()
	log.SetLogLevel(pathUnescapeFnTestLogger, log.DebugLevel)
}

//sunny path case
func Test_pathUnescape_1(t *testing.T) {
	data := map[string]string{
		"":                           "",
		"Hi%2C%20How%20are%20you%3F": "Hi, How are you?",
		"Hi, is this valid?":         "Hi, is this valid?",
		"Hey%2C%20Did%20you%20get%20the%20list%20of%20houses%20that%20where%20sent%3F": "Hey, Did you get the list of houses that where sent?",
		"123%2C%20456+": "123, 456+",
		"Hi%2C%20this%20-%20is%20%2B%20a%20test%20%2B%20string.+And+No+Space.": "Hi, this - is + a test + string.+And+No+Space.",
	}

	for pathUnescapeFnInput, pathUnescapeFnExpectedOutput = range data {
		pathUnescapeFnActualOutput, pathUnescapeFnErr = pathUnescapeFnRef.Eval(pathUnescapeFnInput)

		pathUnescapeFnTestLogger.Debug("In tester: Output of function call = ", pathUnescapeFnActualOutput)

		assert.Nil(t, pathUnescapeFnErr)
		assert.EqualValues(t, pathUnescapeFnExpectedOutput, pathUnescapeFnActualOutput)
	}

}
