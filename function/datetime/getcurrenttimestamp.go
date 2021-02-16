/*
* References: Timezone loading: https://www.golangprograms.com/golang-get-current-date-and-time-in-est-utc-and-mst.html
*           : Layout Options: https://yourbasic.org/golang/format-parse-string-time-date-example/#layout-options
*           :
 */
package customdatetime

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&getCurrentTimestampFn{})
}

type getCurrentTimestampFn struct {
}

// Name returns the name of the function
func (getCurrentTimestampFn) Name() string {
	return "getCurrentTimestamp"
}

// Sig returns the function signature
func (getCurrentTimestampFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString, data.TypeString}, false
}

var getCurrentTimestampFnLogger = log.RootLogger()

// Eval executes the function
func (getCurrentTimestampFn) Eval(params ...interface{}) (interface{}, error) {
	if getCurrentTimestampFnLogger.DebugEnabled() {
		getCurrentTimestampFnLogger.Debugf("Entering function getCurrentTimestamp (eval) with params: [%+v], [%+v], [%+v]", params[0], params[1], params[2])
	}

	defaultTimestampFormat := "2006-01-02T15:04:05"
	defaultTimePrecisionFormat := ".000" //milliseconds
	defaultTimezoneFormat := "Z"

	inputParamTimestampFormatValue := params[0]
	inputParamTimePrecisionValue := params[1]
	inputParamTimezoneValue := params[2]
	timePrecisionString := ""
	timestampFormatString := ""
	timezoneFormatString := ""

	var outputFormat string
	var outputTimestamp string
	var location *time.Location
	var err error

	//Validate Timestamp format.
	//This is for future enhancement.
	//Only support one format today
	switch inputParamTimestampFormatValue {
	case "yyyy-MM-ddThh:mm:ss":
		timestampFormatString = defaultTimestampFormat
	default:
		timestampFormatString = defaultTimestampFormat
	}

	//Assign time precision value. Default is ms
	switch inputParamTimePrecisionValue {
	case "us":
		timePrecisionString = ".000000"
	case "ns":
		timePrecisionString = ".000000000"
	default:
		timePrecisionString = defaultTimePrecisionFormat
	}

	//Validate Timezone format.
	//This is for future enhancement.
	//Only support default timezone today.
	switch inputParamTimezoneValue {
	default:
		timezoneFormatString = defaultTimezoneFormat
		location, err = time.LoadLocation("UTC")

	}

	if err != nil {
		return nil, fmt.Errorf("Invalid Timezone value requested. Error is %+v", err)
	}
	outputFormat = timestampFormatString + timePrecisionString + timezoneFormatString

	if getCurrentTimestampFnLogger.DebugEnabled() {
		getCurrentTimestampFnLogger.Debugf("Final Output Timestamp Format is: [%+v]", outputFormat)
	}

	date := time.Now().In(location)
	outputTimestamp = date.Format(outputFormat)

	if getCurrentTimestampFnLogger.DebugEnabled() {
		getCurrentTimestampFnLogger.Debugf("Final output value = [%+v]", outputTimestamp)
	}

	if getCurrentTimestampFnLogger.DebugEnabled() {
		getCurrentTimestampFnLogger.Debugf("Exiting function getCurrentTimestamp (eval)")
	}

	return outputTimestamp, nil
}
