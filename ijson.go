package ijson

import "fmt"

type (
	Result struct {
		val interface{}
		err Err
	}

	Err struct {
		o error  // original error
		a string // action that caused error
	}
)

func New(data interface{}) (r Result) { r.val = data; return }

func (r Result) Value() interface{} { return r.val }

func (r Result) Error() error {
	if r.err.o == nil {
		return nil
	}

	return r.err
}

func (r Result) Get(path ...string) Result {
	if r.Error() != nil {
		return r
	}

	data, err := Get(r.val, path...)
	if err != nil {
		return Result{err: Err{o: err, a: "GET"}}
	}

	return Result{val: data}
}

func (r Result) Set(value interface{}, path ...string) Result {
	if r.Error() != nil {
		return r
	}

	data, err := Set(r.val, value, path...)
	if err != nil {
		return Result{err: Err{o: err, a: "SET"}}
	}

	return Result{val: data}
}

func (r Result) Del(path ...string) Result {
	if r.Error() != nil {
		return r
	}

	data, err := Del(r.val, path...)
	if err != nil {
		return Result{err: Err{o: err, a: "DELETE"}}
	}

	return Result{val: data}
}

func (e Err) Error() string {
	return fmt.Sprintf("failed to %s : %v", e.a, e.o)
}
