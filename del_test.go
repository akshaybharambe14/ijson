package ijson

import (
	"reflect"
	"testing"
)

func TestDel(t *testing.T) {
	type args struct {
		data interface{}
		path []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "data/ invalid path",
			args:    args{data: []interface{}{1}, path: []string{""}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nil/ valid field",
			args:    args{data: nil, path: []string{"name"}},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "object/ valid field",
			args:    args{data: map[string]interface{}{"name": "tom", "age": 23}, path: []string{"name"}},
			want:    map[string]interface{}{"age": 23},
			wantErr: false,
		},
		{
			name:    "array/ valid field",
			args:    args{data: []interface{}{}, path: []string{"name"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nested/ valid field path",
			args: args{data: map[string]interface{}{
				"name": "tom",
				"friends": map[string]interface{}{
					"jerry": "mouse",
				},
			},
				path: []string{"friends", "jerry"},
			},
			want: map[string]interface{}{
				"name":    "tom",
				"friends": map[string]interface{}{},
			},
			wantErr: false,
		},
		{
			name: "nested/ valid index path",
			args: args{data: []interface{}{
				"tom",
				[]interface{}{"jerry"},
			},
				path: []string{"#1", "#0"},
			},
			want: []interface{}{
				"tom",
				[]interface{}{},
			},
			wantErr: false,
		},
		{
			name: "nested/ invalid field path",
			args: args{data: map[string]interface{}{
				"name": "tom",
				"friends": map[string]interface{}{
					"jerry": "mouse",
				},
			},
				path: []string{"friends", "#0"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nested/ invalid index path",
			args: args{data: []interface{}{
				"tom",
				[]interface{}{"jerry"},
			},
				path: []string{"#0", "name"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "array/ valid index",
			args:    args{data: []interface{}{1, 2, 3, 4}, path: []string{"#1"}},
			want:    []interface{}{1, 4, 3},
			wantErr: false,
		},
		{
			name:    "array/ invalid index",
			args:    args{data: []interface{}{nil, nil}, path: []string{"#a"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "array/ valid index out of range",
			args:    args{data: []interface{}{nil, nil}, path: []string{"#2"}},
			want:    []interface{}(nil),
			wantErr: true,
		},
		{
			name:    "empty array/ valid index",
			args:    args{data: []interface{}{}, path: []string{"#1"}},
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "array/ valid index PO",
			args:    args{data: []interface{}{1, 2, 3, 4}, path: []string{"#~1"}},
			want:    []interface{}{1, 3, 4},
			wantErr: false,
		},
		{
			name:    "array/ valid index PO",
			args:    args{data: []interface{}{1, 2, 3, 4}, path: []string{"#~1"}},
			want:    []interface{}{1, 3, 4},
			wantErr: false,
		},
		{
			name:    "array/ invalid index PO",
			args:    args{data: []interface{}{nil, nil}, path: []string{"#~a"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "array/ valid index PO out of range",
			args:    args{data: []interface{}{nil, nil}, path: []string{"#~2"}},
			want:    []interface{}(nil),
			wantErr: true,
		},
		{
			name:    "empty array/ valid index PO",
			args:    args{data: []interface{}{}, path: []string{"#~1"}},
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "array/ valid delete from end",
			args:    args{data: []interface{}{1, 2, 3, 4}, path: []string{"#"}},
			want:    []interface{}{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "non array/ valid delete from end",
			args:    args{data: 2, path: []string{"#"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Del(tt.args.data, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Del() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
