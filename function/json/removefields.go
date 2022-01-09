package customjson

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data/coerce"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func init() {
	_ = function.Register(&removeFieldsFn{})
}

type removeFieldsFn struct {
}

// Name returns the name of the function
func (removeFieldsFn) Name() string {
	return "removeFields"
}

// Sig returns the function signature
func (removeFieldsFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny}, false
}

var removeFieldsFnLogger = log.RootLogger()

// Eval executes the function
func (removeFieldsFn) Eval(params ...interface{}) (interface{}, error) {

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debugf("Entering function removeFields (eval), with 2 parameters")
		removeFieldsFnLogger.Debugf("Parameter 1 json object = <<%+v>>", params[0])
		removeFieldsFnLogger.Debugf("Parameter 2 array of expressions = <<%+v>>", params[1])
	}

	//First parameter is json object where fields need to be removed.
	//Second parameter is array of expressions (json path expressions),
	//to identify the fields which need to be removed.
	inputParamJsonToUpdate := params[0]
	inputParamArrayOfExpressions := params[1]
	var inputJsonToUpdate map[string]interface{}
	var inputJsonToUpdateByteArray []byte
	var inputJsonToUpdateString string
	var inputArrayOfExpressions []interface{}
	var outputJsonObject map[string]interface{}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debugf("Json to update is = <<%+v>> with type = [%T]", inputParamJsonToUpdate, inputParamJsonToUpdate)
		removeFieldsFnLogger.Debugf("Array of expressions is = <<%+v>> with type [%T]", inputParamArrayOfExpressions, inputParamArrayOfExpressions)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debug("Starting data validation.")
	}

	//If the input object or array of expressions to redact is nil, return the input json object as-is.
	if inputParamJsonToUpdate == nil || inputParamArrayOfExpressions == nil {
		removeFieldsFnLogger.Debug("Input Object or array of field expressions to remove are nil. Returning input as-is.")
		removeFieldsFnLogger.Debugf("Exiting function removeFields (eval)")
		return inputParamJsonToUpdate, nil
	}

	//Check if the input is a json object
	inputJsonObjectToUpdate, ok := inputParamJsonToUpdate.(map[string]interface{})
	if !ok {
		removeFieldsFnLogger.Debugf("Error: First parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
		return nil, fmt.Errorf("Second parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debugf("Input is a JSON Object = <<%+v>> with type [%T]", inputJsonObjectToUpdate, inputJsonObjectToUpdate)
	}

	//If the input object is empty, return the input json object as-is.
	if len(inputJsonObjectToUpdate) == 0 {
		if removeFieldsFnLogger.DebugEnabled() {
			removeFieldsFnLogger.Debug("Input JSON Object is empty. Returning as-is.")
			removeFieldsFnLogger.Debugf("Exiting function removeFields (eval)")
		}

		return inputParamJsonToUpdate, nil
	}

	//Check if the second parameter is an array
	inputArrayOfExpressions, ok = inputParamArrayOfExpressions.([]interface{})
	if !ok {
		removeFieldsFnLogger.Debugf("Error: Second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfExpressions)
		return nil, fmt.Errorf("second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfExpressions)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debug("Printing input values after validation.")
		removeFieldsFnLogger.Debugf("Json Object to update is = << %+v >> of type << %T >>", inputJsonToUpdate, inputJsonToUpdate)
		removeFieldsFnLogger.Debugf("Array of expressions is = << %+v >> of type << %T >>", inputArrayOfExpressions, inputArrayOfExpressions)
	}
	//END - Data Validation

	//Convert input json object to string.
	inputJsonToUpdateString, err := coerce.ToString(inputParamJsonToUpdate)
	//	inputJsonToUpdateString, ok = inputParamJsonToUpdate.(string)
	if err != nil {
		removeFieldsFnLogger.Debugf("Error: Unable to coerce input json object <<%+v>> to string.", inputParamJsonToUpdate)
		return nil, fmt.Errorf("unable to coerce first parameter to string. Error is %+v", err)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debug("Printing Json String after type conversion.")
		removeFieldsFnLogger.Debugf("Json String to redact is = <<%+v>> with type = [%T]", inputJsonToUpdateString, inputJsonToUpdateString)
	}

	//Parse string to JSON Object using gjson.
	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debug("Parsing string to json object using gjson.")
	}
	inputJsonToUpdateGJsonObject := gjson.Parse(inputJsonToUpdateString)

	//NOTE: If individual expressions are not of string type i.e. it is not
	//an array of strings, then input will be returned as is.
	inputJsonToUpdateByteArray = []byte(inputJsonToUpdateString)
	for i, expression := range inputArrayOfExpressions {

		expressionString, ok := expression.(string)

		if ok {
			removeFieldsFnLogger.Debugf("[%d] Input Expression = [%s]", i, expression)

			//remove "$." or "$loop." as sjson does not require it.
			tmpExpression := strings.Replace(expressionString, "$loop.", "", -1)
			tmpExpression = strings.Replace(tmpExpression, "$.", "", -1)

			removeFieldsFnLogger.Debugf("[%d] Modified Expression = [%s]", i, tmpExpression)

			if inputJsonToUpdateGJsonObject.Get(tmpExpression).Exists() {

				result := gjson.GetBytes(inputJsonToUpdateByteArray, tmpExpression)
				tmpArray := result.Array()
				tmpArrayLength := len(tmpArray)
				//fmt.Println("Length raw is: ", len(result.Raw))
				removeFieldsFnLogger.Debugf("[%d] Array Length = [%d]", i, tmpArrayLength)

				if tmpArrayLength > 0 {

					for j := range tmpArray {

						tmpExpressionUpdated := strings.Replace(tmpExpression, "#", strconv.Itoa(j), -1)
						removeFieldsFnLogger.Debugf("[%d][%d] Modified Expression = [%s]", i, j, tmpExpressionUpdated)

						removeFieldsFnLogger.Debugf("[%d][%d] Removing field from JSON.", i, j)
						inputJsonToUpdateByteArray, err = sjson.DeleteBytes(inputJsonToUpdateByteArray, tmpExpressionUpdated)

						if err != nil {
							removeFieldsFnLogger.Debugf("[%d][%d] Unable to delete field [%s] from json object <<%+v>>.", i, j, tmpExpressionUpdated, string(inputJsonToUpdateByteArray))
							return nil, fmt.Errorf("unable to delete field [%s] from json object <<%+v>>", tmpExpressionUpdated, string(inputJsonToUpdateByteArray))
						}

						if removeFieldsFnLogger.DebugEnabled() {
							removeFieldsFnLogger.Debugf("[%d][%d] Updated JSON is = <<%v>>", i, j, string(inputJsonToUpdateByteArray))
						}

					}
					if removeFieldsFnLogger.DebugEnabled() {

						removeFieldsFnLogger.Debugf("[%d] Completed removal of field(s) for expression = [%v]", i, expressionString)
						removeFieldsFnLogger.Debugf("[%d] Updated JSON at the end of iteration is = <<%v>>", i, string(inputJsonToUpdateByteArray))
					}
				} else {
					removeFieldsFnLogger.Debugf("[%d] Ignore Input Expression For Array Search = [%s] as it does not exist in the input json.", i, expression)
				}
			} else {
				removeFieldsFnLogger.Debugf("[%d] Ignore Input Expression = [%s] as it does not exist in the input json.", i, expression)

			}
		} else {
			removeFieldsFnLogger.Debugf("[%d] Ignore Input Expression = [%s] as it is of type [%T]", i, expression, expression)
		}
	}

	//unmarshal to a map.
	err = json.Unmarshal(inputJsonToUpdateByteArray, &outputJsonObject)

	if err != nil {
		removeFieldsFnLogger.Debugf("Unable to unmarshal updated json to an object. Updated JSON is <<%+v>>", string(inputJsonToUpdateByteArray))
		return nil, fmt.Errorf("Unable to unmarshal updated json to an object. Error is %+v.", err)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debugf("Final output value = <<%+v>>", outputJsonObject)
	}

	if removeFieldsFnLogger.DebugEnabled() {
		removeFieldsFnLogger.Debugf("Exiting function removeFields (eval)")
	}

	return outputJsonObject, nil
}
