package customdatetime

import (
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var convertTimestampToEpochTimeFnRef = &convertTimestampToEpochTimeFn{}
var convertTimestampToEpochTimeFnTestLogger log.Logger
var actualOutputTimeInterface interface{}
var actualOutputEpochTime int64
var expectedOutputEpochTime int64

func init() {
	convertTimestampToEpochTimeFnTestLogger = log.RootLogger()
	log.SetLogLevel(convertTimestampToEpochTimeFnTestLogger, log.DebugLevel)
}

//Test output epoch time in seconds.
func Test_convertTimestampToEpochTime_1(t *testing.T) {

	expectedOutputEpochTime = 1586977085
	actualOutputTimeInterface, err := convertTimestampToEpochTimeFnRef.Eval("2020-04-15T18:58:05Z", "", "")

	assert.Nil(t, err)

	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Actual Output of function call = ", actualOutputTimeInterface)
	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Expected Output epoch time = ", expectedOutputEpochTime)

	actualOutputEpochTime = actualOutputTimeInterface.(int64)

	assert.Equal(t, expectedOutputEpochTime, actualOutputEpochTime)
}

//Test output epoch time in milliseconds.
func Test_convertTimestampToEpochTime_2(t *testing.T) {

	expectedOutputEpochTime = 1586977085012
	actualOutputTimeInterface, err := convertTimestampToEpochTimeFnRef.Eval("2020-04-15T18:58:05.012Z", "", "ms")

	assert.Nil(t, err)

	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Actual Output of function call = ", actualOutputTimeInterface)
	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Expected Output epoch time = ", expectedOutputEpochTime)

	actualOutputEpochTime = actualOutputTimeInterface.(int64)

	assert.Equal(t, expectedOutputEpochTime, actualOutputEpochTime)
}

//Test microsecond output.
func Test_convertTimestampToEpochTime_3(t *testing.T) {

	expectedOutputEpochTime = 1586977085012267
	actualOutputTimeInterface, err := convertTimestampToEpochTimeFnRef.Eval("2020-04-15T18:58:05.012267Z", "", "us")

	assert.Nil(t, err)

	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Actual Output of function call = ", actualOutputTimeInterface)
	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Expected Output epoch time = ", expectedOutputEpochTime)

	actualOutputEpochTime = actualOutputTimeInterface.(int64)

	assert.Equal(t, expectedOutputEpochTime, actualOutputEpochTime)
}

//Test nanosecond output.
func Test_convertTimestampToEpochTime_4(t *testing.T) {

	expectedOutputEpochTime = 1585767485012567777
	actualOutputTimeInterface, err := convertTimestampToEpochTimeFnRef.Eval("2020-04-01T18:58:05.012567777Z", "", "ns")

	assert.Nil(t, err)

	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Actual Output of function call = ", actualOutputTimeInterface)
	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Expected Output epoch time = ", expectedOutputEpochTime)

	actualOutputEpochTime = actualOutputTimeInterface.(int64)

	assert.Equal(t, expectedOutputEpochTime, actualOutputEpochTime)
}

//Test custom timestamp format.
func Test_convertTimestampToEpochTime_5(t *testing.T) {

	expectedOutputEpochTime = 1585813565
	actualOutputTimeInterface, err := convertTimestampToEpochTimeFnRef.Eval("Thu Apr 02 02:46:05 -0500 2020", "Mon Jan 02 15:04:05 -0700 2006", "s")

	assert.Nil(t, err)

	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Actual Output of function call = ", actualOutputTimeInterface)
	convertTimestampToEpochTimeFnTestLogger.Debug("In tester: Expected Output epoch time = ", expectedOutputEpochTime)

	actualOutputEpochTime = actualOutputTimeInterface.(int64)

	assert.Equal(t, expectedOutputEpochTime, actualOutputEpochTime)
}
