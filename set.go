package ijson

// Set sets the provide value to the path. It creates the structure if not present.
// An error is returned if it fails to resolve the path OR encounters different type than expected by path.
func Set(data, value interface{}, path ...string) (interface{}, error) {
	return set(data, value, false, path...)
}

// SetP is same as Set(). It just takes `"."` separated path.
func SetP(data, value interface{}, path string) (interface{}, error) {
	return setP(data, value, false, path)
}

// SetF is same as Set(). It just forcefully replaces the structure if it is not same as expected by the path.
func SetF(data, value interface{}, path ...string) (interface{}, error) {
	return set(data, value, true, path...)
}

// SetFP is same as SetF(). It just takes `"."` separated path.
func SetFP(data, value interface{}, path string) (interface{}, error) {
	return setP(data, value, true, path)
}

func set(data interface{}, value interface{}, force bool, path ...string) (interface{}, error) {
	if len(path) == 0 {
		if data == nil {
			return value, nil
		}

		return data, nil
	}

	pathType := DetectSetPath(path[0])
	switch pathType {
	case PSet_Obj:

		object, valid := data.(map[string]interface{})
		if !valid {
			if data == nil || force {
				object = make(map[string]interface{}, 1)
			} else {
				return nil, errExpObj
			}

			// object = make(map[string]interface{})
		}

		newData, err := set(object[path[0]], value, force, path[1:]...)
		if err != nil {
			return nil, err
		}

		object[path[0]] = newData

		return object, nil

	case PSet_ArrIdx:

		idx, err := index(path[0], PSet_ArrIdx)
		if err != nil {
			return nil, err
		}

		array, valid := data.([]interface{})
		if !valid {
			if data == nil || force {
				array = make([]interface{}, idx+1)
			} else {
				return nil, errExpArr
			}
		} else {
			array = extend(array, idx)
		}

		newData, err := set(array[idx], value, force, path[1:]...)
		if err != nil {
			return nil, err
		}

		array[idx] = newData

		return array, nil

	case PSet_ArrAppend:
		array, valid := data.([]interface{})
		if !valid {
			if data == nil || force {
				array = make([]interface{}, 0, 1)
			} else {
				return nil, errExpArr
			}
		}

		array = append(array, value)

		return array, nil

	default:
		return nil, errInvPth
	}
}

func setP(data interface{}, value interface{}, force bool, path string) (interface{}, error) {
	return set(data, value, force, split(path)...)
}

func extend(arr []interface{}, idx int) []interface{} {
	max := len(arr) - 1
	if idx > max {
		arr = append(arr, make([]interface{}, idx-max)...)
	}

	return arr
}
