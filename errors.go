package ijson

import "errors"

var (
	errExpObj = errors.New("expected an object")
	errExpArr = errors.New("expected an array")
	errNotFnd = errors.New("field or index does not exists")
	errOutBnd = errors.New("index out of range")
	errInvPth = errors.New("invalid path")
)
