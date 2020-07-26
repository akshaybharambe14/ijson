package ijson

// Del deletes element form the the data pointed by the path.
// An error is returned if it fails to resolve the path.
func Del(data interface{}, path ...string) (interface{}, error) {
	if len(path) == 0 || data == nil {
		return data, nil
	}

	pathType := DetectDelPath(path[0])
	switch pathType {
	case PDel_Obj:
		object, valid := data.(map[string]interface{})
		if !valid {
			return nil, errExpObj
		}

		if len(path) == 1 {
			delete(object, path[0])
			return object, nil
		}

		newData, err := Del(object[path[0]], path[1:]...)
		if err != nil {
			return nil, err
		}

		object[path[0]] = newData
		return object, nil

	case PDel_ArrIdx, PDel_ArrIdxPO:
		idx, err := index(path[0], pathType)
		if err != nil {
			return nil, err
		}

		array, valid := data.([]interface{})
		if !valid {
			return nil, errExpArr
		}

		if len(path) == 1 {
			return DeleteAtArrayIndex(array, idx, pathType == PDel_ArrIdxPO)
		}

		newData, err := Del(array[idx], path[1:]...)
		if err != nil {
			return nil, err
		}

		array[idx] = newData
		return array, nil

	case PDel_ArrEnd:
		array, valid := data.([]interface{})
		if !valid {
			return nil, errExpArr
		}

		l := len(array)
		array[l-1] = nil
		return array[:l-1], nil

	default:
		return nil, errInvPth
	}
}

// DelP is same as Del() function. It just takes `"."` separated path.
func DelP(data interface{}, path string) (interface{}, error) {
	return Del(data, split(path)...)
}

// DeleteAtArrayIndex deletes the provides index from array.
// Set po to true if you want to preserve the order while deleting the index.
// An error is returned if the index is out of range.
func DeleteAtArrayIndex(
	arr []interface{},
	idx int,
	po bool, /* preser order(inefficient) */
) ([]interface{}, error) {

	if po {
		return DeleteAtIndexPO(arr, idx)
	}

	return DeleteAtIndex(arr, idx)
}

// DeleteAtIndexPO deletes the provides index from array with preserving the order.
func DeleteAtIndexPO(arr []interface{}, idx int) ([]interface{}, error) {
	l := len(arr)

	if l == 0 {
		return arr, nil
	}

	if idx < 0 || idx >= l {
		return nil, errOutBnd
	}

	// Remove the element at index i from a.
	copy(arr[idx:], arr[idx+1:]) // Shift a[i+1:] left one index.
	arr[l-1] = nil               // Erase last element (write zero value).

	// Truncate slice.
	return arr[:l-1], nil
}

// DeleteAtIndex deletes the provides index from array withput preserving the order.
func DeleteAtIndex(arr []interface{}, idx int) ([]interface{}, error) {
	l := len(arr)
	if l == 0 {
		return arr, nil
	}

	if idx < 0 || idx >= l {
		return nil, errOutBnd
	}

	// Remove the element at index i from a.
	arr[idx] = arr[l-1] // Copy last element to index i.
	arr[l-1] = nil      // Erase last element (write zero value).

	// Truncate slice.
	return arr[:l-1], nil
}
