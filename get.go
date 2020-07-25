package ijson

// Get returns the result corresponding to provided path.
//
// If it fails to resolve the path, an error will be returned.
//
// Returns same data if path is not provided, with nil error.
//
// Path syntax - Get(data, "#0", "friends", "#~name", "#")
//
// Explanation -
//
//  "#0" - access the 0th element in the array
//  "friends" - access the friends key from object
//  "#~name" - return an array of all the objects having "name" field.
//  "#" - return the length of the result array
//
func Get(data interface{}, path ...string) (interface{}, error) {
	var err error

	for i := range path {

		switch pathType := DetectGetPath(path[i]); pathType {
		case PGet_Obj:
			data, err = GetObject(data, path[i])

		case PGet_ArrIdx:
			idx, idxErr := index(path[0], PGet_ArrIdx)
			if idxErr != nil {
				return nil, idxErr
			}

			data, err = GetArrayIndex(data, idx)

		case PGet_ArrFld:
			data, err = GetArrayField(data, field(path[i]))

		case PGet_ArrLen:
			return GetArrayLen(data)

		case P_Unknown:
			return nil, errInvPth
		}

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetObject returns the data against the provided field.
//
// It expects the input data to be an object with string key.
//
// An error will be returned if the input is not a valid `map[string]interface{}` OR field does not exists.
func GetObject(data interface{}, field string) (interface{}, error) {
	object, exists := data.(map[string]interface{})
	if !exists {
		return nil, errExpObj
	}

	data, exists = object[field]
	if !exists {
		return nil, errNotFnd
	}

	return data, nil
}

// GetArrayIndex returns the data present at provided index.
//
// It expects the input data to be an array.
//
// An error will be returned if the input is not a valid `[]interface{}` OR index is out of range.
func GetArrayIndex(data interface{}, idx int) (interface{}, error) {

	array, exists := data.([]interface{})
	if !exists {
		return nil, errExpArr
	}

	if idx < 0 || idx >= len(array) {
		return nil, errOutBnd
	}

	return array[idx], nil
}

// GetArrayField returns all the objects matching the provided field.
//
// It expects the input data to be an array of objects.
//
// If provided data is array of values, it will return an empty array.
//
// An error will be returned if the input is not a valid `[]interface{}`.
func GetArrayField(data interface{}, field string) ([]interface{}, error) {
	array, exists := data.([]interface{})
	if !exists {
		return nil, errExpArr
	}

	result := make([]interface{}, len(array))

	var j, k int
	for j = range array {
		o, ok := array[j].(map[string]interface{})
		if !ok {
			continue
		}

		v, ok := o[field]
		if ok {
			result[k] = v
			k++
		}
	}

	return result[:k], nil
}

// GetArrayLen returns length of the array.
//
// It expects the input data to be an array OR nil. Length will be zero if data is nil.
//
// An error will be returned if the input is not nil and not a valid `[]interface{}`.
func GetArrayLen(data interface{}) (int, error) {
	if data == nil {
		return 0, nil
	}

	array, exists := data.([]interface{})
	if !exists {
		return 0, errExpArr
	}

	return len(array), nil
}
