package customdatetime

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&convertTimestampToEpochTimeFn{})
}

type convertTimestampToEpochTimeFn struct {
}

// Name returns the name of the function
func (convertTimestampToEpochTimeFn) Name() string {
	return "convertTimestampToEpochTime"
}

// Sig returns the function signature
func (convertTimestampToEpochTimeFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny, data.TypeAny}, false
}

var convertTimestampToEpochTimeFnLogger = log.RootLogger()

// Eval executes the function
func (convertTimestampToEpochTimeFn) Eval(params ...interface{}) (interface{}, error) {
	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Entering function convertTimestampToEpochTime (eval) with params: [%+v], [%+v], [%+v]", params[0], params[1], params[2])
		convertTimestampToEpochTimeFnLogger.Debugf("Parameter 1 Timestamp = << %+v >>", params[0])
		convertTimestampToEpochTimeFnLogger.Debugf("Parameter 2 Timestamp Format= << %+v >>", params[1])
		convertTimestampToEpochTimeFnLogger.Debugf("Parameter 3 Desired Output time precision = << %+v >>", params[2])
	}

	const defaultTimePrecisionDivisor = int64(1)

	inputParamTimestampValue := params[0]
	inputParamTimestampFormatValue := params[1]
	inputParamDesiredTimePrecisionValue := params[2]

	var epochTimePrecisionDivisor int64
	var timestampFormat, timestamp, epochTimePrecision string

	var outputEpochTime int64
	var err error

	//Timestamp format handling.
	//Coerce to string and default to UTC format, if empty.
	timestampFormat, err = coerce.ToString(inputParamTimestampFormatValue)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce timestamp format to string. Error is %+v", err)
	}

	if len(timestampFormat) <= 0 {
		timestampFormat = time.RFC3339
	}

	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Timestamp format which will be used for parsing is: [%+v]", timestampFormat)
	}

	//Timestamp handling
	timestamp, err = coerce.ToString(inputParamTimestampValue)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce timestamp to string. Error is %+v", err)
	}
	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Input Timestamp is: [%+v]", timestamp)
	}

	//Assign time precision value. Default is ms
	epochTimePrecision, err = coerce.ToString(inputParamDesiredTimePrecisionValue)
	if err != nil {
		return nil, fmt.Errorf("Unable to coerce Desired Epoch Time Precision to string. Error is %+v", err)
	}
	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Desired Epoch Time Precision is: [%+v]", epochTimePrecision)
	}

	switch epochTimePrecision {
	case "us":
		epochTimePrecisionDivisor = int64(time.Microsecond)
	case "ns":
		epochTimePrecisionDivisor = int64(time.Nanosecond)
	case "ms":
		epochTimePrecisionDivisor = int64(time.Millisecond)
	default:
		epochTimePrecisionDivisor = defaultTimePrecisionDivisor
		epochTimePrecision = "s"
	}

	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Updated Epoch Time Desired Precision is: [%+v]", epochTimePrecision)
		convertTimestampToEpochTimeFnLogger.Debugf("Epoch Time Desired Precision divisor is: [%+v]", epochTimePrecisionDivisor)
	}

	//Parse Timestamp.
	datetime, err := time.Parse(timestampFormat, timestamp)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse input timestamp string using provided timestamp format. Error is %+v", err)
	}

	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Datetime after parsing is: [%+v]", datetime)
	}

	if epochTimePrecision == "s" {
		outputEpochTime = datetime.Unix() / epochTimePrecisionDivisor
	} else {
		outputEpochTime = datetime.UnixNano() / epochTimePrecisionDivisor
	}

	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Final Output is: [%+v]", outputEpochTime)
	}

	if convertTimestampToEpochTimeFnLogger.DebugEnabled() {
		convertTimestampToEpochTimeFnLogger.Debugf("Exiting function convertTimestampToEpochTime (eval)")
	}

	return outputEpochTime, nil
}
