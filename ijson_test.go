package ijson

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantR Result
	}{
		{
			name:  "init result",
			args:  args{data: []interface{}{}},
			wantR: Result{val: []interface{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := New(tt.args.data); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("New() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestResult_Get(t *testing.T) {
	r := Result{val: Object(), err: Err{o: errExpArr, a: "SET"}}
	type args struct {
		path []string
	}
	tests := []struct {
		name string
		r    Result
		args args
		want Result
	}{
		{
			name: "object/ valid path",
			r:    New(Object()),
			args: args{path: []string{"id"}},
			want: Result{val: Object()["id"]},
		},
		{
			name: "object/ valid path/ existing error",
			r:    r,
			args: args{path: []string{"id"}},
			want: r,
		},
		// FIXME: failed tests because of error interface
		// {
		// 	name: "object/ invalid path",
		// 	r:    New(Object()),
		// 	args: args{path: []string{"#0"}},
		// 	want: Result{val: nil, err: Err{o: errExpObj, a: "GET"}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Get(tt.args.path...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
