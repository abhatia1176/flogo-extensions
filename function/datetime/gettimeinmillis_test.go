package customdatetime

import (
	"testing"
	"time"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var getTimeInMillisFnRef = &getTimeInMillisFn{}
var getTimeInMillisFnTestLogger log.Logger
var outputTimeInMillisInterface interface{}
var outputTimeInMillis int64
var expectedTimeInMillis int64

func init() {
	getTimeInMillisFnTestLogger = log.RootLogger()
	log.SetLogLevel(getTimeInMillisFnTestLogger, log.DebugLevel)
}

func Test_getTimeInMillis_1(t *testing.T) {

	outputTimeInMillisInterface, err := getTimeInMillisFnRef.Eval()
	assert.Nil(t, err)
	getTimeInMillisFnTestLogger.Debug("In tester: Output of function call = ", outputTimeInMillisInterface)

	outputTimeInMillis = outputTimeInMillisInterface.(int64)
	expectedTimeInMillis = time.Now().UnixNano() / 1e6
	timeElapsedInMillis := expectedTimeInMillis - outputTimeInMillis

	getTimeInMillisFnTestLogger.Debug("In tester: Time elapsed = ", timeElapsedInMillis)

	assert.InDelta(t, expectedTimeInMillis, outputTimeInMillis, 1000, "Took longer than 1 ms to execute.")
}
