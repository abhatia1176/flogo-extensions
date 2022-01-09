package customjson

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var removeFieldsFnRef = &removeFieldsFn{}
var removeFieldsFnTestLogger log.Logger
var inputJsonDataToRemove interface{}
var inputFieldsToRemove []interface{}
var expectedJsonDataToRemove interface{}
var expectedJsonDataRemoved interface{}
var actualJsonDataRemoved map[string]interface{}
var arr string

func init() {
	removeFieldsFnTestLogger = log.RootLogger()
	log.SetLogLevel(removeFieldsFnTestLogger, log.DebugLevel)

	data := `{
        "lead": {
            "email": "ab@test.com",
			"password": "123456",
			"test": {
					"key":"123456",
					"key2":"1789"
			}
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(data), &inputJsonDataToRemove)
}

//remove fields from a JSON object.
func Test_removeFields_1(t *testing.T) {

	//declare input array of jpath expressions
	arr = `["$.lead.password", "$.lead.test.key"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
        "lead": {
            "email": "ab@test.com",
			"test": {
					"key2":"1789"
			}
        },
        "eventType": "test"
    }`

	//unmarshal array of jpath expressions, and expected output redacted json string.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	//invoke function under test.
	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the Removeed json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//remove fields from an array.
func Test_removeFields_2(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.lead.password", "$.rootArray.#.lead.test.key", "$.rootArray.#.obj.test.key3", "$.rootArray.#.lead.test.key3", "$.eventType", "$.eventType2"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
			}
		}]
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//TC3 - Array of expression is empty.
func Test_removeFields_3(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = ``
	//declare expected output i.e. updated json string.
	expectedJsonDataToRemove := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}],
		"eventType": "test"
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	inputFieldsToRemove = nil
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove all objects from array.
// test with same expression for field removal twice in the array.
func Test_removeFields_4(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.lead", "$.rootArray.#.lead"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
		"rootArray": [{		}, {		}],
		"eventType": "test"
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove top level array.
func Test_removeFields_5(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = `["$.rootArray", "$.rootArray.#.lead"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
		"eventType": "test"
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove fields from a null object.
func Test_removeFields_6(t *testing.T) {

	//declare input data - json to be redacted.
	var data map[string]interface{}
	//data = nil
	inputJsonDataToRemove = data

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.lead.password", "$.rootArray.#.lead.test.key", "$.rootArray.#.lead.test.key3", "$.eventType", "$.eventType2"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := data

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	//json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataToRemove, actualJsonDataRemoved)

	//Print the updated json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove fields from an empty object.
func Test_removeFields_7(t *testing.T) {

	//declare input data - json to be redacted.
	//var data map[string]interface{}
	data := `{}`

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.lead.password", "$.rootArray.#.lead.test.key", "$.rootArray.#.lead.test.key3", "$.eventType", "$.eventType2"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := data

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the updated json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove fields from an array.
func Test_removeFields_8(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890",
					"key3": "7890"					
				}
			}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.lead.password", "$.rootArray.#.lead.test.key", "$.rootArray.#.lead.test.key3", "$.eventType", "$.eventType2"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
		"rootArray": [{
			"lead": {
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
			}
		}, {
			"lead": {
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
			}
		}]
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}

//remove fields from an array.
func Test_removeFields_9(t *testing.T) {

	//declare input data - json to be redacted.
	data := `{
		"rootArray": [{
				"email": "ab@test.com",
				"password": "123456",
				"test": {
					"key": "123456",
					"key2": "7890"
				}
		}, {
				"email": "ab@test.com",
				"password": "123456",
				"apiKey":"asdasd",				
				"test": {
					"key": "123456",
					"key2": "7890",
					"key3": "7890"					
				}
		}],
		"eventType": "test"
	}`
	json.Unmarshal([]byte(data), &inputJsonDataToRemove)

	//declare input array of jpath expressions
	arr = `["$.rootArray.#.apiKey", "$.rootArray.#.password", "$.rootArray.#.test.key", "$.rootArray.#.test.key3", "$.eventType", "$.eventType2"]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToRemove := `{
		"rootArray": [{
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
		}, {
				"email": "ab@test.com",
				"test": {
					"key2": "7890"
				}
		}]
	}`

	//unmarshal input fields to redact array, and expected output.
	//expected output is unmarshaled, so it is easy to compare.
	json.Unmarshal([]byte(arr), &inputFieldsToRemove)
	json.Unmarshal([]byte(expectedJsonDataToRemove), &expectedJsonDataRemoved)

	actualJsonDataRemoved, err := removeFieldsFnRef.Eval(inputJsonDataToRemove, inputFieldsToRemove, 1)

	//print actual output.
	removeFieldsFnTestLogger.Debug("Actual Output = ", actualJsonDataRemoved)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataRemoved, actualJsonDataRemoved)

	//Print the redacted json output. Could print the string output of function directly as well.
	r, _ := json.Marshal(actualJsonDataRemoved)
	removeFieldsFnTestLogger.Debug("Updated JSON String: ", string(r))

}