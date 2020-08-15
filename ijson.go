package ijson

import "fmt"

type (
	Result struct {
		val interface{}
		err error
	}

	Err struct {
		o error  // original error
		a string // action that caused error
	}
)

func New(data interface{}) (r Result) { r.val = data; return }

func (r Result) Value() interface{} { return r.val }

func (r Result) Error() error {
	if r.err == nil {
		return nil
	}

	return r.err
}

func (r Result) GetP(path string) Result {
	return r.get(split(path)...)
}

func (r Result) Get(path ...string) Result {
	return r.get(path...)
}

func (r Result) get(path ...string) Result {
	if r.Error() != nil {
		return r
	}

	data, err := Get(r.val, path...)
	if err != nil {
		return Result{err: Err{o: err, a: "GET"}}
	}

	return Result{val: data}
}

func (r Result) SetP(value interface{}, path string) Result {
	return r.set(value, split(path)...)
}

func (r Result) Set(value interface{}, path ...string) Result {
	return r.set(value, path...)
}

func (r Result) set(value interface{}, path ...string) Result {
	if r.Error() != nil {
		return r
	}

	data, err := Set(r.val, value, path...)
	if err != nil {
		return Result{err: Err{o: err, a: "SET"}}
	}

	return Result{val: data}
}

func (r Result) DelP(path string) Result {
	return r.del(split(path)...)
}

func (r Result) Del(path ...string) Result {
	return r.del(path...)
}
func (r Result) del(path ...string) Result {
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
