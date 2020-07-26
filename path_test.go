package ijson

import (
	"reflect"
	"testing"
)

func TestDetectPath(t *testing.T) {
	type args struct {
		a Actn
		p string
	}
	tests := []struct {
		name string
		args args
		want Path
	}{
		{
			name: "Get object",
			args: args{a: Act_Get, p: "name"},
			want: PGet_Obj,
		},
		{
			name: "Get array length",
			args: args{a: Act_Get, p: "#"},
			want: PGet_ArrLen,
		},
		{
			name: "Get array index",
			args: args{a: Act_Get, p: "#1"},
			want: PGet_ArrIdx,
		},
		{
			name: "Get array field",
			args: args{a: Act_Get, p: "#~name"},
			want: PGet_ArrFld,
		},
		{
			name: "Get unknown path",
			args: args{a: Act_Get, p: ""},
			want: P_Unknown,
		},
		{
			name: "Get path for unknown action",
			args: args{a: Act_Ukn, p: ""},
			want: P_Unknown,
		},
		{
			name: "Set object",
			args: args{a: Act_Set, p: "name"},
			want: PSet_Obj,
		},
		{
			name: "Set array index",
			args: args{a: Act_Set, p: "#1"},
			want: PSet_ArrIdx,
		},
		{
			name: "Set array append",
			args: args{a: Act_Set, p: "#"},
			want: PSet_ArrAppend,
		},
		{
			name: "Set unknown path",
			args: args{a: Act_Set, p: ""},
			want: P_Unknown,
		},

		{
			name: "Del object",
			args: args{a: Act_Del, p: "name"},
			want: PDel_Obj,
		},
		{
			name: "Del array index",
			args: args{a: Act_Del, p: "#1"},
			want: PDel_ArrIdx,
		},
		{
			name: "Del array index with order preserved",
			args: args{a: Act_Del, p: "#~1"},
			want: PDel_ArrIdxPO,
		},
		{
			name: "Del array end",
			args: args{a: Act_Del, p: "#"},
			want: PDel_ArrEnd,
		},
		{
			name: "Del unknown path",
			args: args{a: Act_Del, p: ""},
			want: P_Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectPath(tt.args.a, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_field(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid array field",
			args: args{p: "#~friends"},
			want: "friends",
		},
		{
			name: "invalid array field",
			args: args{p: "friends"},
			want: "iends",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := field(tt.args.p); got != tt.want {
				t.Errorf("field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_index(t *testing.T) {
	type args struct {
		p string
		t Path
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "set array index",
			args:    args{p: "#10", t: PSet_ArrIdx},
			want:    10,
			wantErr: false,
		},
		{
			name:    "del array index with order preserved",
			args:    args{p: "#~10", t: PDel_ArrIdxPO},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := index(tt.args.p, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("index() = %v, want %v", got, tt.want)
			}
		})
	}
}
