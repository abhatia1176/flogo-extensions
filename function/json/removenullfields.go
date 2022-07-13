package customjson

import (
	"encoding/json"
	"fmt"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
	"github.com/tidwall/sjson"
)

func init() {
	_ = function.Register(&removeNullFieldsFn{})
}

type removeNullFieldsFn struct {
}

// Name returns the name of the function
func (removeNullFieldsFn) Name() string {
	return "removeNullFields"
}

// Sig returns the function signature
func (removeNullFieldsFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny, data.TypeAny, data.TypeAny, data.TypeAny, data.TypeAny}, false
}

var removeNullFieldsFnLogger = log.RootLogger()
var updatedJsonStringGlobal string

// contains takes a slice and looks for an element in it. If found it will
// return true, otherwise it will return false.
func contains(slice []string, val string) (bool) {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func removeNullFieldsIsEmpty(inputParamValue interface{}, treatEmptyAsNull bool) (bool) {

	inputString := ""
	var outputValue bool
	var err error

	switch  inputParamValue.(type) {
	case nil:
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debugf("Processing value of type nil.")
			removeNullFieldsFnLogger.Debugf( "Input Parameter is nil. Returning true.")
		}
		outputValue = true
	case []interface{}:
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debugf("Processing an array of interfaces.")
		}
		if len(inputParamValue.([]interface{})) <= 0 {
			outputValue = true
			removeNullFieldsFnLogger.Debugf( "Input Parameter is empty. Returning true.")
		}
		if(!treatEmptyAsNull) {
			outputValue = false
		}
	case map[string]interface{}:
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debug( "Processing a map of interface (json object).")
		}
		if len(inputParamValue.(map[string]interface{}) ) <= 0 {
			outputValue = true
			removeNullFieldsFnLogger.Debugf( "Input Parameter is empty. Returning true.")
		}
		if(!treatEmptyAsNull) {
			outputValue = false
		}
	default:
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debugf( "Applying default processing logic for input param value of type %+T.", inputParamValue)
		}
		inputString, err =  coerce.ToString(inputParamValue)
		if err != nil {
			removeNullFieldsFnLogger.Debugf("Unable to coerece input value to a string. value = %+v.", inputParamValue)
			outputValue = false
		} else {
			if removeNullFieldsFnLogger.DebugEnabled() {
				removeNullFieldsFnLogger.Debugf("Input Parameter's string length is %+v.", len(inputString))
			}

			if len(inputString) <= 0 {
				outputValue = true
				removeNullFieldsFnLogger.Debugf("Input Parameter is empty or nil. Returning true.")
			}
		}
		if(!treatEmptyAsNull) {
			outputValue = false
		}
	}//ENDS - switch statement.

	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Final output value = %+v", outputValue)
	}

	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Exiting function isEmpty (eval)")
	}

	return outputValue
}

func removeNullFieldsFromJson(updatedJsonString *string, key string, data interface{}, keyPath string,  arrayNames []string, isArray bool, matchingArrayFound bool, removeEmptyObjects bool, removeArrayIfAnyObjectRemoved bool, treatEmptyAsNull bool, recursionCounter int) (b bool, c bool) {
		counter := recursionCounter
		currentKeyPath := ""
		deleteArray := false;
		deleteElement := false;

		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debug("[", counter, "] ", "Enter Function.")
		}
		if (keyPath != "") {
			currentKeyPath = keyPath + "." + key
		} else {
			currentKeyPath = key
		}
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debugf("[%+v] with key = <<%+v>>", counter, key)
			removeNullFieldsFnLogger.Debugf("[%+v] having value = <<%+v>>", counter, data)
			removeNullFieldsFnLogger.Debugf("[%+v] of Type = <<%T>>", counter, data)
			removeNullFieldsFnLogger.Debugf("[%+v] current key path = <<%+v>>", counter, currentKeyPath)
		}
		switch m := data.(type) {
		case map[string]interface{}:
			if removeNullFieldsFnLogger.DebugEnabled() {
				removeNullFieldsFnLogger.Debug("[", counter, "] ", "Processing a map of interface (json object) and iterating on elements.")
			}
			removedElementsCounter := 0
			mapElementCounter:=0
			for k, v := range m {
				if removeNullFieldsFnLogger.DebugEnabled() {
					removeNullFieldsFnLogger.Debugf("[%+v][%+v] <<%+v>> = <<%+v>>", counter,mapElementCounter, k, v)
				}
				if removeNullFieldsIsEmpty(v, treatEmptyAsNull) {
					removeNullFieldsFnLogger.Debugf("[%+v][%+v] Value is nil. Remove key from json.",counter, mapElementCounter)
					removeNullFieldsFnLogger.Debugf("[%+v][%+v] Key = [%+v]", counter, mapElementCounter, k)
					deleteNullKeyFullPath := ""
					if (currentKeyPath != "") {
						deleteNullKeyFullPath = currentKeyPath + "." + k
					} else {
						deleteNullKeyFullPath = k
					}
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v][%+v] Base key path is : [%+v]", counter, mapElementCounter, currentKeyPath)
						removeNullFieldsFnLogger.Debugf("[%+v][%+v] Full key path is : [%+v]", counter, mapElementCounter,  deleteNullKeyFullPath)
					}

					if(matchingArrayFound) {
						if removeNullFieldsFnLogger.DebugEnabled() {
							removeNullFieldsFnLogger.Debugf("[%+v][%+v] Soft delete. Return true.", counter, mapElementCounter)
						}
					}
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v][%+v] Delete key from JSON.", counter, mapElementCounter)
					}
					deleteArray=true
					removedElementsCounter++
					//use sjson delete function to delete the value and assign it to redactedJsonString.
					*updatedJsonString, _ = sjson.Delete(*updatedJsonString, deleteNullKeyFullPath)
				} else {
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v][%+v] Element is of type<<%T>>, Calling recursive function.",counter, mapElementCounter, v)
					}
					counter ++;
					_, deleteElement1 := removeNullFieldsFromJson(updatedJsonString, k, v, currentKeyPath, arrayNames, isArray, matchingArrayFound, removeEmptyObjects, removeArrayIfAnyObjectRemoved , treatEmptyAsNull, counter)
					/*if(deleteArray1) {
						//removedElementsCounter++
					}*/
					if(deleteElement1) {
						removedElementsCounter++
					}
				}
				mapElementCounter++;
			}

			//if all elements in the map are removed. remove map itself.
			if(mapElementCounter == removedElementsCounter) {
				if(isArray || removeEmptyObjects) {
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v] All elements of object are removed. Removing object.", counter)
					}
					//use sjson delete function to delete the value.
					*updatedJsonString, _ = sjson.Delete(*updatedJsonString, currentKeyPath)

					//indicate that the entire object was removed from the array.
					deleteArray = true;
				}
			} else if (matchingArrayFound && removedElementsCounter >= 1) {
				//remove the object from array, because it is a matching array - requested by user for this behavior.
				//use sjson delete function to delete the value.
				*updatedJsonString, _ = sjson.Delete(*updatedJsonString, currentKeyPath)

				//indicate that the entire object was removed from the array.
				deleteArray = true
			}	else {
				if removeNullFieldsFnLogger.DebugEnabled() {
					removeNullFieldsFnLogger.Debugf("[%+v] No need to remove object from this array.", counter)
				}
			}
		case []interface{}:
			if removeNullFieldsFnLogger.DebugEnabled() {
				removeNullFieldsFnLogger.Debug("[", counter, "] ", "Processing an array of interface and iterating on elements.")
			}

			tempMatchingArrayFound := false
			removeArray := false
			removedItemCount := 0
			arrayItemCount := 0
			if(contains(arrayNames,key)) {
				tempMatchingArrayFound = true
			}
			for k, v := range m {
				if removeNullFieldsFnLogger.DebugEnabled() {
					removeNullFieldsFnLogger.Debugf("[%+v] <<%+v>> = <<%+v>>", counter, k, v)
				}
				//if v == nil {
				if removeNullFieldsIsEmpty(v, treatEmptyAsNull) {
					removeNullFieldsFnLogger.Debugf("[%+v] Value is nil. Remove key from json.",counter)
					removeNullFieldsFnLogger.Debugf("[%+v] Key = [%+v]", counter, k)

				} else {
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v] Element is of type<<%T>>. Calling recursive function.", counter, v)
					}
					counter++;
					tempKey, _ := coerce.ToString(k)
					tempRemoveArray, tempRemovedElements := removeNullFieldsFromJson(updatedJsonString, tempKey, v, currentKeyPath, arrayNames, true, tempMatchingArrayFound, removeEmptyObjects, removeArrayIfAnyObjectRemoved , treatEmptyAsNull, counter)

					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v] tempRemovedElements [%+v].", counter, tempRemovedElements)
						removeNullFieldsFnLogger.Debugf("[%+v] tempRemoveArray [%+v].", counter, tempRemoveArray)
					}
					//capture count of array items which are removed.
					if(tempRemoveArray) {
						removedItemCount++
					}

					//Once array has been set for deletion. Do not update the status based on other array element status.
					//This is used, if the user wants to delete the entire array, if even one object from the array is deleted.
					if (!removeArray) {
						removeArray = tempRemoveArray
					}

				}
				arrayItemCount++
			}
			//all objects in the object array were removed. remove array
			if(arrayItemCount == removedItemCount) {
				if removeNullFieldsFnLogger.DebugEnabled() {
					removeNullFieldsFnLogger.Debugf("[%+v] All objects of the array were deleted. Delete Array.", counter)
					removeNullFieldsFnLogger.Debugf("[%+v] Delete following key [%v].", counter, currentKeyPath)
					//use sjson delete function to delete the value and assign it to redactedJsonString.
					*updatedJsonString, _ = sjson.Delete(*updatedJsonString, currentKeyPath)
				}
			}else if(tempMatchingArrayFound) {
				//delete array, if one element was null and it is a matching array, and being requested to be removed.
				if(removeArrayIfAnyObjectRemoved) {
					if(removeArray) {
						if removeNullFieldsFnLogger.DebugEnabled() {
							removeNullFieldsFnLogger.Debugf("[%+v] Matching array which had null keys removed.", counter)
							removeNullFieldsFnLogger.Debugf("[%+v] Delete following key [%v].", counter, currentKeyPath)
						}
						//use sjson delete function to delete the value and assign it to redactedJsonString.
						*updatedJsonString, _ = sjson.Delete(*updatedJsonString, currentKeyPath)

					} else {
						removeNullFieldsFnLogger.Debugf("[%+v] Matching array but no need to remove. Array not updated.", counter)
					}
				} else {
					removeNullFieldsFnLogger.Debugf("[%+v] Matching array but not requested to be removed if empty.", counter)
				}
			}
		default:
			removeNullFieldsFnLogger.Debugf("[%+v] ** Default type value", counter)
			removeNullFieldsFnLogger.Debugf("[%+v] with value = <<%+v>>", counter, data)
			if removeNullFieldsIsEmpty(data, treatEmptyAsNull) {
				if(matchingArrayFound) {
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("[%+v] Soft delete. Return true.", counter)
						//deleteArray=true
						deleteElement=true
					}
				} else {
					if removeNullFieldsFnLogger.DebugEnabled() {
						removeNullFieldsFnLogger.Debugf("Iteration = [%+v] Remove Key - %+v : %+v", counter, key, data)
					}
					//use sjson delete function to delete the value and assign it to redactedJsonString.
					*updatedJsonString, _ = sjson.Delete(*updatedJsonString, key)
				}
			}
		}
		return deleteArray, deleteElement
}

// Eval executes the function
// This function removes fields which have null values, from the input json object.
// It expects a json object and returns a json object with null fields removed.
// There is an optional second parameter (can be nil), where one can specify an array of strings,
// where each string value represents name of an object array, containing name-value pair
// elements in each object {"sampleArray": [{"name": "somefield", "value": "somevalue"}]}
// and the entire object needs to be removed, if either of the name or value is empty.
func (removeNullFieldsFn) Eval(params ...interface{}) (interface{}, error) {
	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Entering function removeNullFields (eval), with five parameters")
		removeNullFieldsFnLogger.Debugf("First Param = <<%+v>>", params[0])
		removeNullFieldsFnLogger.Debugf("Second Param = <<%+v>>",params[1])
		removeNullFieldsFnLogger.Debugf("Third Param = <<%+v>>",params[2])
		removeNullFieldsFnLogger.Debugf("Fourth Param = <<%+v>>",params[3])
		removeNullFieldsFnLogger.Debugf("Fourth Param = <<%+v>>",params[4])
	}

	inputParamJsonToUpdate := params[0]
	inputParamArrayOfNames := params[1]
	inputParamRemoveEmptyObjects := params[2]
	inputParamRemoveArrayIfAnyObjectRemoved := params[3]
	inputParamTreatEmptyAsNull := params[4]


	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Type of first parameter (json object) = [%T]", inputParamJsonToUpdate)
		removeNullFieldsFnLogger.Debugf("Type of second parameter (array of array names) = [%T]", inputParamArrayOfNames)
		removeNullFieldsFnLogger.Debugf("Type of third parameter (inputParamRemoveEmptyObjects) = [%T]", inputParamRemoveEmptyObjects)
		removeNullFieldsFnLogger.Debugf("Type of fourth parameter (inputParamRemoveArrayIfAnyObjectRemoved) = [%T]", inputParamRemoveArrayIfAnyObjectRemoved)
		removeNullFieldsFnLogger.Debugf("Type of fifth parameter (inputParamTreatEmptyAsNull) = [%T]", inputParamTreatEmptyAsNull)
	}

	//validate input
	//If the input json is nil, return the input json as-is.
	if inputParamJsonToUpdate == nil {
		return inputParamJsonToUpdate, nil
	}

	//VALIDATION --- START

	//Check if the input is a json object
	inputJsonObjectToUpdate, ok := inputParamJsonToUpdate.(map[string]interface{})
	if (!ok) {
		removeNullFieldsFnLogger.Debugf("First parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
		return nil, fmt.Errorf("Second parameter must be a json object, but it is of type %T", inputParamJsonToUpdate)
	}

	//Check if the input is an array
	inputArrayOfArrayNames, ok := inputParamArrayOfNames.([]interface{})
	if (!ok && inputParamArrayOfNames != nil) {
		removeNullFieldsFnLogger.Debugf("Second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfNames)
		return nil, fmt.Errorf("Second parameter must be an array of interface{}, but it is of type %T", inputParamArrayOfNames)
	}

	//Check if the input is a boolean
	inputRemoveEmptyObjects, ok := inputParamRemoveEmptyObjects.(bool)
	if (!ok && inputParamArrayOfNames != nil) {
		removeNullFieldsFnLogger.Debugf("Third parameter must be a boolean, but it is of type %T", inputParamRemoveEmptyObjects)
		return nil, fmt.Errorf("Third parameter must be a boolean, but it is of type %T", inputParamRemoveEmptyObjects)
	}

	//Check if the input is a boolean
	inputRemoveArrayIfAnyObjectRemoved, ok := inputParamRemoveArrayIfAnyObjectRemoved.(bool)
	if (!ok && inputParamArrayOfNames != nil) {
		removeNullFieldsFnLogger.Debugf("Fourth parameter must be a boolean, but it is of type %T", inputParamRemoveArrayIfAnyObjectRemoved)
		return nil, fmt.Errorf("Fourth parameter must be a boolean, but it is of type %T", inputParamRemoveArrayIfAnyObjectRemoved)
	}

	//Check if the input is a boolean
	inputTreatEmptyAsNull, ok := inputParamTreatEmptyAsNull.(bool)
	if (!ok && inputParamArrayOfNames != nil) {
		removeNullFieldsFnLogger.Debugf("Fifth parameter must be a boolean, but it is of type %T", inputParamTreatEmptyAsNull)
		return nil, fmt.Errorf("Fifth parameter must be a boolean, but it is of type %T", inputParamTreatEmptyAsNull)
	}

	//coerce json to string and validate
	inputJsonStringToUpdate, err := coerce.ToString(inputJsonObjectToUpdate)
	if (err != nil) {
		removeNullFieldsFnLogger.Debugf("Unable to coerce input json object <<%+v>> to string.", inputJsonObjectToUpdate)
		return nil, fmt.Errorf("Unable to coerce input json object to string. Error is %+v", err)
	}

	//convert input array names to an array of string
	var inputStringArrayOfArrayNames = make([]string,0)
	var tempArrayOfArrayNamesCounter = 0
	for k, v := range inputArrayOfArrayNames {
		if removeNullFieldsFnLogger.DebugEnabled() {
			removeNullFieldsFnLogger.Debugf("[%+v] <<%+v>> = <<%+v>>", tempArrayOfArrayNamesCounter, k, v)
		}

		if v == nil {
			removeNullFieldsFnLogger.Debugf("[%+v] Value is nil. Ignore value.",tempArrayOfArrayNamesCounter)
		} else {
			stringArrayName, ok := v.(string)
			if !ok {
			removeNullFieldsFnLogger.Debugf("[%+v] Value is not a string. Return error.",tempArrayOfArrayNamesCounter)
			return nil, fmt.Errorf("Second parameter must be an array of string, but it is not, as it contains a value of type %T", v)

			} else {
				inputStringArrayOfArrayNames = append(inputStringArrayOfArrayNames, stringArrayName)
			}
		}
		tempArrayOfArrayNamesCounter++
	}

	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debug("Printing input values after type conversion.")
		removeNullFieldsFnLogger.Debugf("Json Object to update is = <<%+v>> with type = [%T]", inputJsonObjectToUpdate, inputJsonObjectToUpdate)
		removeNullFieldsFnLogger.Debugf("Json String to update is  = <<%+v>> with type [%T]", inputJsonStringToUpdate, inputJsonStringToUpdate)
		removeNullFieldsFnLogger.Debugf("Array Names to check, for removal of null values = <<%+v>> with type [%T]", inputArrayOfArrayNames, inputArrayOfArrayNames)
		removeNullFieldsFnLogger.Debugf("Array Names to check, for removal of null values = <<%+v>> with type [%T]", inputStringArrayOfArrayNames, inputStringArrayOfArrayNames)

	}

	//VALIDATION --- COMPLETE

	var outputJsonObject map[string]interface{}
	var outputJsonString string

	updatedJsonStringGlobal = inputJsonStringToUpdate
	removeNullFieldsFromJson(&inputJsonStringToUpdate, "", inputJsonObjectToUpdate, "", inputStringArrayOfArrayNames, false, false,	inputRemoveEmptyObjects, inputRemoveArrayIfAnyObjectRemoved,   inputTreatEmptyAsNull, 0)
	outputJsonString = inputJsonStringToUpdate

	//unmarshal to a map.
	err = json.Unmarshal([]byte(outputJsonString), &outputJsonObject)

	if (err != nil) {
		removeNullFieldsFnLogger.Debugf("Unable to unmarshal updated json to an object. Updated JSON is <<%+v>>", outputJsonString)
		return nil, fmt.Errorf("Unable to unmarshal updated json to an object. Error is %+v.", err)
	}

	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Final output value = <<%+v>>", outputJsonString)
	}

	if removeNullFieldsFnLogger.DebugEnabled() {
		removeNullFieldsFnLogger.Debugf("Exiting function removeNullFields (eval)")
	}

	return outputJsonObject, nil
}