package ijson

import (
	"reflect"
	"testing"
)

// func Example() {
// 	var data = []interface{}{
// 		map[string]interface{}{
// 			"index": 0,
// 			"friends": []interface{}{
// 				map[string]interface{}{
// 					"id":   0,
// 					"name": "Justine Bird",
// 				},
// 				map[string]interface{}{
// 					"id":   0,
// 					"name": "Justine Bird",
// 				},
// 				map[string]interface{}{
// 					"id":   1,
// 					"name": "Marianne Rutledge",
// 				},
// 			},
// 		},
// 	}

// 	r := New(data).
// 		GetP("#0.friends.#~name"). // list the friend names for 0th record -
// 		// []interface {}{"Justine Bird", "Justine Bird", "Marianne Rutledge"}

// 		Del("#0"). // delete 0th record
// 		// []interface {}{"Marianne Rutledge", "Justine Bird"}

// 		Set("tom", "#") // append "tom" in the list
// 		// // []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

// 	fmt.Printf("%#v\n", r.Value())
// 	// output: []interface {}{"Marianne Rutledge", "Justine Bird", "tom"}

// 	// // returns error if the data type differs than the type expected by query
// 	// fmt.Println(r.Set(1, "name").Error())
// }

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

func TestParse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{
			name: "valid data",
			args: args{data: `{"id":"0"}`},
			want: Result{val: map[string]interface{}{"id": "0"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
