package ijson

func Set(data, value interface{}, path ...string) (interface{}, error) {
	return set(data, value, false, path...)
}

func SetF(data, value interface{}, path ...string) (interface{}, error) {
	return set(data, value, true, path...)
}

func set(data interface{}, value interface{}, force bool, path ...string) (interface{}, error) {
	if len(path) == 0 {
		return value, nil
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

func extend(arr []interface{}, idx int) []interface{} {
	max := len(arr) - 1
	if idx > max {
		arr = append(arr, make([]interface{}, idx-max)...)
	}

	return arr
}
