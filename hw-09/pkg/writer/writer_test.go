package writer

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	type args struct {
		a []any
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name:  "Strings",
			args:  args{[]any{"Hello", "World"}},
			wantW: "HelloWorld",
		},
		{
			name:  "Not Strings",
			args:  args{[]any{1, []string{"Hello", "World"}}},
			wantW: "",
		},
		{
			name:  "Mixed",
			args:  args{[]any{"Hello", 1, []string{"Hello", "World"}}},
			wantW: "Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Write(w, tt.args.a...)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Write() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
