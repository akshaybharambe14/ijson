package ijson

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
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
			name:    "nil/ nil",
			args:    args{data: nil, path: nil},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "array/ valid index",
			args:    args{data: Array(), path: []string{"#0"}},
			want:    Array()[0],
			wantErr: false,
		},
		{
			name:    "object/ valid field",
			args:    args{data: Object(), path: []string{"name"}},
			want:    Object()["name"],
			wantErr: false,
		},
		{
			name:    "nested/ valid path",
			args:    args{data: Nested(), path: []string{"#0", "friends", "#~name", "#"}},
			want:    3,
			wantErr: false,
		},
		{
			name:    "nested/ invalid path",
			args:    args{data: Nested(), path: []string{"", "", "#~name", "#"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "array/ invalid index",
			args:    args{data: Array(), path: []string{"#a"}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nil/ valid index",
			args:    args{data: nil, path: []string{"#2"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.data, tt.args.path...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetObject(t *testing.T) {
	type args struct {
		data  interface{}
		field string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid data with field",
			args:    args{data: Object(), field: "name"},
			want:    Object()["name"],
			wantErr: false,
		},
		{
			name:    "valid data with invalid field",
			args:    args{data: Object(), field: "qualification"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid data",
			args:    args{data: Array(), field: "qualification"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetObject(tt.args.data, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArrayIndex(t *testing.T) {
	type args struct {
		data interface{}
		idx  int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid array with valid index",
			args:    args{data: Array(), idx: 0},
			want:    Array()[0],
			wantErr: false,
		},
		{
			name:    "valid array with invalid index",
			args:    args{data: Array(), idx: 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid array",
			args:    args{data: Object(), idx: 10},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArrayIndex(tt.args.data, tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrayIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArrayIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArrayField(t *testing.T) {
	type args struct {
		data  interface{}
		field string
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name:    "valid array with valid filed",
			args:    args{data: Array(), field: "id"},
			want:    []interface{}{0, 0, 1},
			wantErr: false,
		},
		{
			name:    "valid array with mixed valid filed",
			args:    args{data: append(Array(), nil), field: "id"},
			want:    []interface{}{0, 0, 1},
			wantErr: false,
		},
		{
			name:    "valid array with invalid filed",
			args:    args{data: Array(), field: "tags" /* any random name */},
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "invalid array",
			args:    args{data: Object(), field: "tags" /* any random name */},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArrayField(tt.args.data, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrayField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArrayField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArrayLen(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "valid array",
			args:    args{data: Array()},
			want:    3,
			wantErr: false,
		},
		{
			name:    "nil array",
			args:    args{data: nil},
			want:    0,
			wantErr: false,
		},
		{
			name:    "invalid array",
			args:    args{data: Object()},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArrayLen(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArrayLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
