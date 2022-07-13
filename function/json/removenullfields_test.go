package customjson

import (
	"encoding/json"
	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

var removeNullFieldsFnRef = &removeNullFieldsFn{}
var removeNullFieldsFnTestLogger log.Logger
var removeNullFieldsFnActualOutput map[string] interface{}
var removeNullFieldsFnActualOutputJsonObject interface{}
var removeNullFieldsFnExpectedOutputJsonString string
var removeNullFieldsFnExpectedOutputJsonObject interface{}
var removeNullFieldsFnInputJsonString string
var removeNullFieldsFnInputJsonObject interface{}
var removeNullFieldsFnInputArrayOfArrayNamesString string
var removeNullFieldsFnInputArrayOfArrayNamesObject []interface{}
var removeNullFieldsFnErr error

func init() {
	removeNullFieldsFnTestLogger = log.RootLogger()
	log.SetLogLevel(removeNullFieldsFnTestLogger, log.DebugLevel)
}
//sunny path case
//remove simple null keys.
func Test_removeNullFields_1(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray":[{"test":"1123"},{"test":"123123"}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray":[{"test":"1123"},{"test":"123123"}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case
//remove entire object from array, because all elements of that object are null.
func Test_removeNullFields_2(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray":[{"test":null},{"test":"123123"}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray":[{"test":"123123"}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case
//1. objectArray1 -- remove entire object of the array, because one element is null and it is a matching array.
//2. objectArray2 -- remove only the key which is null from the object, because non-matching array.
func Test_removeNullFields_3(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1"},{"prop_name":"prop2","prop_value":"value2"}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case
//1. objectArray1 -- remove entire object of the array, because one element is null and it is a matching array.
//2. objectArray2 -- remove only the key which is null from the object, because non-matching array.
//3. objectArray3 -- remove the entire array because all elements in all objects are null.
func Test_removeNullFields_4(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray3":[{"prop_name":null,"prop_value":null},{"prop_name":null,"prop_value":null}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1"},{"prop_name":"prop2","prop_value":"value2"}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case
//1. objectArray1 -- remove entire object of the array, because one element is null and it is a matching array.
//2. objectArray2 -- remove only the key which is null from the object, because non-matching array.
//3. objectArray3 -- remove the entire array because all elements in all objects are null.
//4. objectArray4 -- remove only address of type home, because addr1 field is null.
func Test_removeNullFields_5(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray3":[{"prop_name":null,"prop_value":null},{"prop_name":null,"prop_value":null}],
			"objectArray4":[{"address":{"type":"home","addr1":null}}, {"address":{"type":"home2","addr1":"300 lane"}}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1"},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray4":[{"address":{"type":"home"}}, {"address":{"type":"home2","addr1":"300 lane"}}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}


//sunny path case
//lead object is returned as empty object, because both fields are null and removed.
//check test case 7, to see how removing the empty object works.
func Test_removeNullFields_6(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
		"lead": {
			"key":null,
			"key2": null
		},
		"id":123
	}`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
		"lead":{},
		"id":123
	}`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case
//Empty object lead will be removed. It will be empty because all null fields will be deleted.
func Test_removeNullFields_7(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
		"lead": {
			"key":null,
			"key2": null
		},
		"id":123
	}`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
		"id":123
	}`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, true, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case - variation of test case 5
//1. objectArray1 -- remove entire object of the array, because one element is null and it is a matching array.
//                    Also, since removeArrayIfAnyObject removed is true, it removed the entire array
//2. objectArray2 -- remove only the key which is null from the object, because non-matching array.
//3. objectArray3 -- remove the entire array because all elements in all objects are null.
//4. objectArray4 -- remove only address of type home, because addr1 field is null.
//4. objectArray4 -- remove entire array because it is a matching array, having at least one value removed.
func Test_removeNullFields_8(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"removeNullKey": null
			},
			"removeNulllObject": null,
			"simpleArray":["123","123"],
			"objectArray1":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray2":[{"prop_name":"prop1","prop_value":null},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray3":[{"prop_name":null,"prop_value":null},{"prop_name":null,"prop_value":null}],
			"objectArray4":[{"address":{"type":"home","addr1":null}}, {"address":{"type":"home2","addr1":"300 lane"}}]
        },
        "eventType": null
    }`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["simpleArray", "objectArray1", "objectArray4"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456"
			},
			"simpleArray":["123","123"],
			"objectArray2":[{"prop_name":"prop1"},{"prop_name":"prop2","prop_value":"value2"}],
			"objectArray4":[{},{"address":{"type":"home2","addr1":"300 lane"}}]
        }
    }`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, true, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case - variation of test case 5
//remove arrays with
func Test_removeNullFields_9(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
		"messageMetadata": {
			"retryAttempt": null,
			"retryAttemptStartTimestamp": 642,
			"retryDestination": "ABCD",
			"tibcoAuditSafeSequence": 790.75,
			"transactionId": "ABCDEFGHIJKLMNOPQR",
			"action": "ABCDEFGHI",
			"additionalProps": [{
				"prop_name": "name",
				"prop_value": null
			}],
			"auditSafeProps": [{
				"prop_name": "abc",
				"prop_value": null
			}]
		},
		"messagePayload": ""
	}`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["additionalProps", "auditSafeProps"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
		"messageMetadata": {
			"retryAttemptStartTimestamp": 642,
			"retryDestination": "ABCD",
			"tibcoAuditSafeSequence": 790.75,
			"transactionId": "ABCDEFGHIJKLMNOPQR",
			"action": "ABCDEFGHI"
		},
		"messagePayload": ""
	}`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, false)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}

//sunny path case - variation of test case 5
//treat empty as null
func Test_removeNullFields_10(t *testing.T) {
	//Set input parameter 1 = json object
	removeNullFieldsFnInputJsonString = `{
		"messageMetadata": {
			"retryAttempt": null,
			"retryAttemptStartTimestamp": 642,
			"retryDestination": "ABCD",
			"tibcoAuditSafeSequence": 790.75,
			"transactionId": "ABCDEFGHIJKLMNOPQR",
			"action": "",
			"additionalProps": [{
				"prop_name": "name",
				"prop_value": ""
			},
			{	
				"prop_name": "name2",
				"prop_value": "test1"
			}],
			"auditSafeProps": [{
				"prop_name": "abc",
				"prop_value": null
			}],
			"redactionInfo":{
				"pathOfFieldsToRedactInPayload": null
			}
		},
		"messagePayload": ""
	}`
	json.Unmarshal([]byte(removeNullFieldsFnInputJsonString), &removeNullFieldsFnInputJsonObject)

	//Set input parameter 2 = json array
	removeNullFieldsFnInputArrayOfArrayNamesString := `["additionalProps", "auditSafeProps", "pathOfFieldsToRedactInPayload"]`
	json.Unmarshal([]byte(removeNullFieldsFnInputArrayOfArrayNamesString), &removeNullFieldsFnInputArrayOfArrayNamesObject)

	//Setup expected output
	removeNullFieldsFnExpectedOutputJsonString = `{
		"messageMetadata": {
			"retryAttemptStartTimestamp": 642,
			"retryDestination": "ABCD",
			"tibcoAuditSafeSequence": 790.75,
			"transactionId": "ABCDEFGHIJKLMNOPQR",
			"additionalProps": [{	
				"prop_name": "name2",
				"prop_value": "test1"
			}],
			"redactionInfo":{}
		}
	}`
	json.Unmarshal([]byte(removeNullFieldsFnExpectedOutputJsonString), &removeNullFieldsFnExpectedOutputJsonObject)

	//invoke function under test.
	removeNullFieldsFnActualOutput, err := removeNullFieldsFnRef.Eval(removeNullFieldsFnInputJsonObject, removeNullFieldsFnInputArrayOfArrayNamesObject, false, false, true)

	//convert function output to map[string] interface{}.
	removeNullFieldsFnActualOutputJsonObject = removeNullFieldsFnActualOutput.(map[string]interface {})

	//print actual output.
	removeNullFieldsFnTestLogger.Debug("Actual Output = ",removeNullFieldsFnActualOutputJsonObject)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, removeNullFieldsFnExpectedOutputJsonObject, removeNullFieldsFnActualOutputJsonObject)
}