package customjson

import (
	"encoding/json"
	"testing"

	"github.com/project-flogo/core/support/log"
	"github.com/stretchr/testify/assert"
)

var indata string
var setFieldValuesFnRef = &setFieldValuesFn{}
var setFieldValuesFnTestLogger log.Logger
var inputJsonStringToUpdate interface{}
var inputJsonObjectToUpdate map[string]interface{}
var inputFieldsToUpdate []interface{}
var expectedJsonDataToUpdate interface{}
var expectedJsonDataUpdated map[string]interface{}
var actualJsonDataUpdated map[string]interface{}

func init() {
	setFieldValuesFnTestLogger = log.RootLogger()
	log.SetLogLevel(setFieldValuesFnTestLogger, log.DebugLevel)

	indata = `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(indata), &inputJsonObjectToUpdate)
	inputJsonStringToUpdate = indata
}

//Test 1 - positive test case
//Final replacement values are string values.

func Test_setFieldValues_1(t *testing.T) {

	//declare input array of jpath expressions
	arr := `[
		{
		  "name": "$.lead.lastname",
		  "value": "Bhatia"
		},
		{
		  "name": "$.lead.name",
		  "value": "Abhishek"
		}
	  ]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToUpdate := `{
        "lead": {
            "name": "Abhishek",
            "lastname": "Bhatia",			
			"password": "123123",
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonStringToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 2 - positive test case
//Final Replacement values are a combination of string and object.
func Test_setFieldValues_2(t *testing.T) {

	//declare input array of jpath expressions
	arr := `[
		{
		  "name": "$.lead.lastname",
		  "value": {"ln1": "Bhatia", "ln2": "BHATIA"}
		},
		{
		  "name": "$.lead.name",
		  "value": "Abhishek"
		}
	  ]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToUpdate := `{
        "lead": {
            "name": "Abhishek",
            "lastname": {"ln1": "Bhatia", "ln2": "BHATIA"},			
			"password": "123123",
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonStringToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 3 - positive test case
//Final Replacement values are a combination of string and object.
//Input is JSON Object.
func Test_setFieldValues_3(t *testing.T) {

	indata = `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"age": 29,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(indata), &inputJsonObjectToUpdate)
	inputJsonStringToUpdate = indata

	//declare input array of jpath expressions
	arr := `[
		{
		  "name": "$.lead.lastname",
		  "value": {"ln1": "Bhatia", "ln2": "BHATIA"}
		},
		{
		  "name": "$.lead.name",
		  "value": "Abhishek"
		},
		{
		  "name": "$.lead.age",
		  "value": 33
		}
	  ]`
	//declare expected output i.e. redacted json string.
	expectedJsonDataToUpdate := `{
        "lead": {
            "name": "Abhishek",
            "lastname": {"ln1": "Bhatia", "ln2": "BHATIA"},			
			"password": "123123",
			"age": 33,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonObjectToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 4 - positive test case
//Final Replacement values are a combination of string and object.
//Input JSON in empty.
func Test_setFieldValues_4(t *testing.T) {

	indata = `{}`
	var inputJsonObjectToUpdate map[string]interface{}
	setFieldValuesFnTestLogger.Debugf("Input data is : %+v", indata)

	err := json.Unmarshal([]byte(indata), &inputJsonObjectToUpdate)
	inputJsonStringToUpdate = indata

	setFieldValuesFnTestLogger.Debugf("Input data after marshalling is : %+v", inputJsonObjectToUpdate)

	if err != nil {
		setFieldValuesFnTestLogger.Debugf("Cannot set up input data. Error: %+v", err)

	}
	//declare input array of jpath expressions
	arr := `[
		{
		  "name": "$.lead.lastname",
		  "value": {"ln1": "Bhatia", "ln2": "BHATIA"}
		},
		{
		  "name": "$.lead.name",
		  "value": "Abhishek"
		},
		{
		  "name": "$.lead.age",
		  "value": 33
		}
	  ]`
	//declare expected output i.e. updated json string.
	expectedJsonDataToUpdate = `{}`
	var expectedJsonDataUpdated map[string]interface{}

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate.(string)), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonObjectToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 5 - positive test case
//Final Replacement values are a combination of string and object.
//Input JSON in nil.
func Test_setFieldValues_5(t *testing.T) {

	inputJsonObjectToUpdate = nil
	setFieldValuesFnTestLogger.Debugf("Input data is : %+v", inputJsonObjectToUpdate)

	//declare input array of jpath expressions
	arr := `[
		{
		  "name": "$.lead.lastname",
		  "value": {"ln1": "Bhatia", "ln2": "BHATIA"}
		},
		{
		  "name": "$.lead.name",
		  "value": "Abhishek"
		},
		{
		  "name": "$.lead.age",
		  "value": 33
		}
	  ]`

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	//declare expected output i.e. updated json string.

	var expectedJsonDataToUpdate string
	var expectedJsonDataUpdated map[string]interface{}
	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonObjectToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 6 - positive test case
//Final Replacement values array is empty.
//data returned as-is.
func Test_setFieldValues_6(t *testing.T) {

	indata = `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"age": 29,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(indata), &inputJsonObjectToUpdate)
	inputJsonStringToUpdate = indata

	//declare input array of jpath expressions
	arr := `[]`

	//declare expected output i.e. redacted json string.
	expectedJsonDataToUpdate := `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"age": 29,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	//unmarshal array of jpath expressions, and expected output updated json string.
	json.Unmarshal([]byte(arr), &inputFieldsToUpdate)
	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonObjectToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}

//Test 7 - positive test case
//Final Replacement values array is null.
//data returned as-is.
func Test_setFieldValues_7(t *testing.T) {

	indata = `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"age": 29,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	json.Unmarshal([]byte(indata), &inputJsonObjectToUpdate)
	inputJsonStringToUpdate = indata

	//declare expected output i.e. redacted json string.
	expectedJsonDataToUpdate := `{
        "lead": {
            "name": "Test1",
            "lastname": "Test2",			
			"password": "123123",
			"age": 29,
			"test": {
					"key":"123456"
			}
        },
        "eventType": "test"
    }`

	//declare input array and leave uninitialized
	var inputFieldsToUpdate []interface{}

	json.Unmarshal([]byte(expectedJsonDataToUpdate), &expectedJsonDataUpdated)

	//invoke function under test.
	actualJsonDataUpdatedInterface, err := setFieldValuesFnRef.Eval(inputJsonObjectToUpdate, inputFieldsToUpdate, "update", false, false)

	//convert function output to map[string] interface, using unmarshal.
	actualJsonDataUpdated := actualJsonDataUpdatedInterface.(map[string]interface{})

	//print actual output.
	setFieldValuesFnTestLogger.Debug("Actual Output = ", actualJsonDataUpdated)

	//assert error is nil.
	assert.Nil(t, err)

	//assert input matches output.
	assert.EqualValues(t, expectedJsonDataUpdated, actualJsonDataUpdated)

	//Print the updated json output.
	r, _ := json.Marshal(actualJsonDataUpdated)
	setFieldValuesFnTestLogger.Debug("Updated JSON String: ", string(r))
}
