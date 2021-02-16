package customdatetime

import (
	"testing"
	"time"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var getTimeInNanoFnRef = &getTimeInNanoFn{}
var getTimeInNanoFnTestLogger log.Logger
var outputTimeInNanoInterface interface{}
var outputTimeInNano int64
var expectedTimeInNano int64

func init() {
	getTimeInNanoFnTestLogger = log.RootLogger()
	log.SetLogLevel(getTimeInNanoFnTestLogger, log.DebugLevel)
}

func Test_getTimeInNano_1(t *testing.T) {

	outputTimeInNanoInterface, err := getTimeInNanoFnRef.Eval()
	assert.Nil(t, err)
	getTimeInNanoFnTestLogger.Debug("In tester: Output of function call = ", outputTimeInNanoInterface)

	outputTimeInNano = outputTimeInNanoInterface.(int64)
	expectedTimeInNano = time.Now().UnixNano()
	timeElapsedInNano := expectedTimeInNano - outputTimeInNano

	getTimeInNanoFnTestLogger.Debug("In tester: Time elapsed = ", timeElapsedInNano)

	assert.InDelta(t, expectedTimeInNano, outputTimeInNano, 1000, "Took longer than 1 ms to execute.")
}
