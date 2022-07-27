package customjson

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func init() {
	_ = function.Register(&setFieldValuesFn{})
}

type setFieldValuesFn struct {
}

// Name returns the name of the function
func (setFieldValuesFn) Name() string {
	return "setFieldValues"
}

// Sig returns the function signature
func (setFieldValuesFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny, data.TypeString, data.TypeBool, data.TypeBool}, false
}

var setFieldValuesFnLogger = log.RootLogger()

// Eval executes the function
func (setFieldValuesFn) Eval(params ...interface{}) (interface{}, error) {

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debugf("Entering function setFieldValues (eval), with 2 parameters")
		setFieldValuesFnLogger.Debugf("Parameter 1 json object = <<%+v>>", params[0])
		setFieldValuesFnLogger.Debugf("Parameter 2 array of NVPairs with expressions & values = <<%+v>>", params[1])
		setFieldValuesFnLogger.Debugf("Parameter 3 (is the operation) = <<%+v>>", params[0])
		setFieldValuesFnLogger.Debugf("Parameter 4 (is boolean indicating whether to throw error if duplicate key in JSON with insert/upsert operation) = <<%+v>>", params[0])
		setFieldValuesFnLogger.Debugf("Parameter 5 (is boolean indicating whether to throw error if key not in JSON with update/upsert operation)  = <<%+v>>", params[0])
	}

	//First parameter is supposed to be a json object (or json string) where field values need to be updated.
	//Second parameter is array of objects, where each object is a name/value pair - name contains the
	//json path expression for the field, and value contains the replacement value.
	//Other parameters are not implemented in this version.
	//Third param - insert, update, upsert
	//Fourth Param - true  - throw error if insert/upsert leads to duplicate key in json.
	//               false - silently ignore error but do not add duplicate key.
	//Fifth Param - true  - throw error if update/upsert does not find key in json.
	//              false - silently ignore key which is not found in json.

	inputParamJsonToUpdate := params[0]
	inputParamArrayOfExpressions := params[1]
	inputParamOperation := params[2]
	inputParamThrowErrorOnDupKeyInJson := params[3]
	inputParamThrowErrorOnKeyNotInJson := params[4]

	//input
	var inputJsonToUpdate map[string]interface{}
	var inputArrayOfExpressions []interface{}
	var inputStringToUpdate string
	var inputOperation string
	var inputThrowErrorOnDupKeyInJson bool
	var inputThrowErrorOnKeyNotInJson bool

	//intermediate
	var ok bool
	var err error
	var inputGJsonObject gjson.Result

	//output
	var updatedJsonString string
	var updatedJson map[string]interface{}

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debugf("Json to update is = <<%+v>> with type = [%T]", inputParamJsonToUpdate, inputParamJsonToUpdate)
		setFieldValuesFnLogger.Debugf("Array of expressions is = <<%+v>> with type [%T]", inputParamArrayOfExpressions, inputParamArrayOfExpressions)
	}

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Starting data validation.")
	}
	//If the input object or array of expressions to redact is nil, return the input json object as-is.
	if inputParamJsonToUpdate == nil || inputParamArrayOfExpressions == nil {
		return inputParamJsonToUpdate, nil
	}

	//check if first parameter is a json object
	switch inputParamJsonToUpdate.(type) {

	case string:
		//Using unmarshal instead of coerce.
		err = json.Unmarshal([]byte(inputParamJsonToUpdate.(string)), &inputJsonToUpdate)
		if err != nil {
			setFieldValuesFnLogger.Debugf("First parameter must be a Json Object but it is not. Error is %+v", err)
			return nil, fmt.Errorf("First parameter must be a Json Object but it is not. Error is %+v", err)
		}

		inputStringToUpdate = inputParamJsonToUpdate.(string)

	case map[string]interface{}:
		inputJsonToUpdate, ok = inputParamJsonToUpdate.(map[string]interface{})
		if !ok {
			setFieldValuesFnLogger.Debugf("First parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
			return nil, fmt.Errorf("Second parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
		}

		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debugf("Input is a JSON Object = <<%+v>> with type [%T]", inputJsonToUpdate, inputJsonToUpdate)
		}

		//Convert input json object to string.
		inputStringToUpdate, err = coerce.ToString(inputParamJsonToUpdate)
		if err != nil {
			setFieldValuesFnLogger.Debugf("Unable to coerce input json object <<%+v>> to string.", inputParamJsonToUpdate)
			return nil, fmt.Errorf("unable to coerce first parameter to string. Error is %+v", err)
		}

		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debug("Printing Json String after type conversion.")
			setFieldValuesFnLogger.Debugf("Json String to update is = <<%+v>> with type = [%T]", inputStringToUpdate, inputStringToUpdate)
		}

	default:
		setFieldValuesFnLogger.Debugf("First parameter must be a Json Object but it is not. Type is %T", inputParamJsonToUpdate)
		return nil, fmt.Errorf("First parameter must be a Json Object but it is not. Type is %T", inputParamJsonToUpdate)
	}

	//If the input object is empty, return the input json object as-is.
	if len(inputJsonToUpdate) == 0 {
		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debug("Input JSON Object is empty. Returning as-is.")
		}

		return inputJsonToUpdate, nil
	}

	//Check if the second parameter is an array
	inputArrayOfExpressions, ok = inputParamArrayOfExpressions.([]interface{})
	if !ok {
		setFieldValuesFnLogger.Debugf("Second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfExpressions)
		return nil, fmt.Errorf("second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfExpressions)
	}

	//Identify Operation to be performed
	//1. insert.
	//2. update.
	//3. upsert.
	if inputParamOperation == nil {
		inputOperation = "update"
	}
	inputOperation, err = coerce.ToString(inputParamOperation)
	if err != nil {
		setFieldValuesFnLogger.Debug("Operation could not be coerced to String. Using default operation (update).")
		inputOperation = "update"
	} else {
		setFieldValuesFnLogger.Debugf("Operation successfully coerced to String. Value is [%+v]", inputOperation)
	}
	switch inputOperation {
	case "insert", "update", "upsert":
		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debug("Valid Operation.")
		}
	default:
		inputOperation = "update"
		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debug("Defaulting Operation to (update).")
		}
	}

	//Check if the input is a boolean
	inputThrowErrorOnDupKeyInJson, ok = inputParamThrowErrorOnDupKeyInJson.(bool)
	if !ok {
		setFieldValuesFnLogger.Debugf("Fourth parameter must be a boolean, but it is of type %T. Defaulting to false.", inputParamThrowErrorOnDupKeyInJson)
		inputThrowErrorOnDupKeyInJson = false
	}

	//Check if the input is a boolean
	inputThrowErrorOnKeyNotInJson, ok = inputParamThrowErrorOnKeyNotInJson.(bool)
	if !ok {
		setFieldValuesFnLogger.Debugf("Fifth parameter must be a boolean, but it is of type %T. Defaulting to false.", inputParamThrowErrorOnKeyNotInJson)
		inputThrowErrorOnKeyNotInJson = false
	}

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Printing input values after validation.")
		setFieldValuesFnLogger.Debugf("Json Object to update is = << %+v >> of type << %T >>", inputJsonToUpdate, inputJsonToUpdate)
		setFieldValuesFnLogger.Debugf("Json Object (string representation) to update is = << %+v >> of type << %T >>", inputStringToUpdate, inputStringToUpdate)
		setFieldValuesFnLogger.Debugf("Array of expressions is = << %+v >> of type << %T >>", inputArrayOfExpressions, inputArrayOfExpressions)
		setFieldValuesFnLogger.Debugf("Operation is = << %+v >> of type << %T >>", inputOperation, inputOperation)
		setFieldValuesFnLogger.Debugf("Operation is = << %+v >> of type << %T >>", inputThrowErrorOnDupKeyInJson, inputThrowErrorOnDupKeyInJson)
		setFieldValuesFnLogger.Debugf("Operation is = << %+v >> of type << %T >>", inputThrowErrorOnKeyNotInJson, inputThrowErrorOnKeyNotInJson)
	}

	//If the input object is empty, return the input json object as-is.
	if len(inputArrayOfExpressions) == 0 {
		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debug("Input Array is empty. Returning as-is.")
		}

		return inputJsonToUpdate, nil
	}
	//Data Validation - END

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Convert input JSON String to GJSON object using gjson parse.")
	}
	inputGJsonObject = gjson.Parse(inputStringToUpdate)

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Successfully parsed to GJSON Result object.")
		setFieldValuesFnLogger.Debugf("Json Object to update is = << %+v >> of type << %T >>", inputGJsonObject, inputGJsonObject)
	}

	//Iterate over each NVPair object in the array, and update input JSON.
	//If array expressions find no match, then input will be returned as is.
	//Assign input to output in case of no changes.
	updatedJsonString = inputStringToUpdate
	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Logic to update JSON Object - START.")
	}

	for i, value := range inputArrayOfExpressions {

		if setFieldValuesFnLogger.DebugEnabled() {
			setFieldValuesFnLogger.Debugf("[%d] Object in array is = << %+v >> of type [%T]", i, value, value)
		}

		switch m := value.(type) {
		case map[string]interface{}:
			var expression, tmpExpression string
			var replacementValue interface{}
			var doNotProcess bool

			doNotProcess = false

			for k, v := range m {
				setFieldValuesFnLogger.Debugf("[%d] Object element key = [%s], value = [%s]", i, k, v)

				if k == "name" {
					setFieldValuesFnLogger.Debugf("[%d] Object key is 'name' i.e. field to be udpated = << %+v >> of type [%T]", i, v, v)

					expression = v.(string)

					setFieldValuesFnLogger.Debugf("[%d] Input Expression = [%s]", i, expression)

					//remove "$." or "$loop." as sjson does not require it.
					tmpExpression = strings.Replace(expression, "$loop.", "", -1)
					tmpExpression = strings.Replace(tmpExpression, "$.", "", -1)

					setFieldValuesFnLogger.Debugf("[%d] Modified Expression = [%s]", i, tmpExpression)
				} else if k == "value" {
					setFieldValuesFnLogger.Debugf("[%d] Object key is 'value' i.e. target field value = << %+v >> of type [%T]", i, v, v)
					replacementValue = v
				} else {
					setFieldValuesFnLogger.Debugf("[%d] Object in array contains unrecognized key = [%s]", i, k)
					setFieldValuesFnLogger.Debugf("[%d] Accepted keys are 'name' and 'value'", i)
					setFieldValuesFnLogger.Debugf("[%d] Skipping JSON Update step.", i)
					doNotProcess = true
				}
			}
			if !doNotProcess {
				if inputGJsonObject.Get(tmpExpression).Exists() {
					//use sjson set function to update the value and assign it back to JsonString.
					updatedJsonString, _ = sjson.Set(updatedJsonString, tmpExpression, replacementValue)
				} else {
					setFieldValuesFnLogger.Debugf("[%d] Ignore Input Expression = [%s] as it does not exist in the input json.", i, expression)

				}
			} else {
				setFieldValuesFnLogger.Debugf("[%d] JSON Update Skipped.", i)
			}
		default:
			print("Invalid Array Structure")
		}
	} //end for loop - iterate on expressions/value pair.

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Logic to update JSON Object - END.")
		setFieldValuesFnLogger.Debugf("Updated JSON String = [%s]", updatedJsonString)
		setFieldValuesFnLogger.Debug("Unmarshal JSON String to Object.")
	}
	err = json.Unmarshal([]byte(updatedJsonString), &updatedJson)
	if err != nil {
		setFieldValuesFnLogger.Debugf("Unable to unmarshal output string to output JSON object. Error is %+v", err)
		return nil, fmt.Errorf("Unable to unmarshal output string to output JSON object. Error is %+v", err)
	}

	if setFieldValuesFnLogger.DebugEnabled() {
		setFieldValuesFnLogger.Debug("Successfully unmarshalled.")
		setFieldValuesFnLogger.Debugf("Updated JSON Object = << %+v >>", updatedJson)
	}
	return updatedJson, nil
}
