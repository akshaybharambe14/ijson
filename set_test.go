package ijson

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	type args struct {
		data  interface{}
		value interface{}
		path  []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "nil data/ invalid path",
			args: args{
				data:  nil,
				value: 1,
				path:  []string{""},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil/ valid array index",
			args: args{
				data:  nil,
				value: 1,
				path:  []string{"#1"},
			},
			want:    []interface{}{nil, 1},
			wantErr: false,
		},
		{
			name: "nil/ invalid index",
			args: args{
				data:  nil,
				value: 1,
				path:  []string{"#a"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "array/ valid index",
			args: args{
				data:  []interface{}{1},
				value: 1,
				path:  []string{"#1"},
			},
			want:    []interface{}{1, 1},
			wantErr: false,
		},
		{
			name: "array/ valid path/ nil nested object",
			args: args{
				data:  []interface{}{},
				value: 1,
				path:  []string{"#1", "name"},
			},
			want:    []interface{}{nil, map[string]interface{}{"name": 1}},
			wantErr: false,
		},
		{
			name: "array/ valid path/ invalid nested object",
			args: args{
				data:  []interface{}{[]interface{}{}},
				value: 1,
				path:  []string{"#0", "name"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "object/ valid index",
			args: args{
				data:  make(map[string]interface{}),
				value: 1,
				path:  []string{"#1"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "array/ append",
			args: args{
				data:  []interface{}{},
				value: 1,
				path:  []string{"#"},
			},
			want:    []interface{}{1},
			wantErr: false,
		},
		{
			name: "nil/ append",
			args: args{
				data:  nil,
				value: 1,
				path:  []string{"#"},
			},
			want:    []interface{}{1},
			wantErr: false,
		},
		{
			name: "object/ append",
			args: args{
				data:  make(map[string]interface{}),
				value: 1,
				path:  []string{"#"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nil/ valid field",
			args: args{
				data:  nil,
				value: 1,
				path:  []string{"name"},
			},
			want:    map[string]interface{}{"name": 1},
			wantErr: false,
		},
		{
			name: "object/ valid field",
			args: args{
				data:  map[string]interface{}{"visits": 1},
				value: "tom",
				path:  []string{"name"},
			},
			want:    map[string]interface{}{"visits": 1, "name": "tom"},
			wantErr: false,
		},
		{
			name: "object/ valid path/ invalid nested array",
			args: args{
				data:  map[string]interface{}{"visits": []interface{}{}},
				value: 1,
				path:  []string{"visits", "count"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "object/ nil path",
			args: args{
				data:  map[string]interface{}{"visits": []interface{}{}},
				value: 1,
				path:  nil,
			},
			want:    map[string]interface{}{"visits": []interface{}{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Set(tt.args.data, tt.args.value, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetF(t *testing.T) {
	type args struct {
		data  interface{}
		value interface{}
		path  []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "object/ valid index",
			args: args{
				data:  make(map[string]interface{}),
				value: 1,
				path:  []string{"#1"},
			},
			want:    []interface{}{nil, 1},
			wantErr: false,
		},
		{
			name: "array/ valid field",
			args: args{
				data:  []interface{}{nil, 1},
				value: 1,
				path:  []string{"name"},
			},
			want:    map[string]interface{}{"name": 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetF(tt.args.data, tt.args.value, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extend(t *testing.T) {
	type args struct {
		arr []interface{}
		idx int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "existing array with index within range",
			args: args{arr: []interface{}{1, 2, 3}, idx: 2},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "existing array with index out of range",
			args: args{arr: []interface{}{1, 2, 3}, idx: 5},
			want: []interface{}{1, 2, 3, nil, nil, nil},
		},
		{
			name: "nil array with index out of range",
			args: args{arr: nil, idx: 5},
			want: []interface{}{nil, nil, nil, nil, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extend(tt.args.arr, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extend() = %v, want %v", got, tt.want)
			}
		})
	}
}
