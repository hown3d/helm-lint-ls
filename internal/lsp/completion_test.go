package lsp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	lsp "go.lsp.dev/protocol"
)

var testData string = `
	name: {{ .Values.tes }}
`

func Test_completeValue(t *testing.T) {
	type args struct {
		pos lsp.Position
	}
	tests := []struct {
		name    string
		args    args
		want    []lsp.CompletionItem
		wantErr bool
	}{
		{
			args: args{
				pos: lsp.Position{
					Line: uint32(5),
				},
			},
			name: "read value",
			want: []lsp.CompletionItem{{
				Label: ".Values.tes",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			openAndReadFile = func(filepath string, linenumber uint32) (string, error) {
				return testData, nil
			}
			got, err := completeValue("", tt.args.pos)
			if tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getBracetsText(t *testing.T) {
	type args struct {
		line string
		pos  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with bracets",
			args: args{
				line: "ABC 123 {{ .Values }} ABC",
				pos:  15,
			},
			want: "{{ .Values }}",
		},

		{
			name: "no bracets in text",
			args: args{
				line: "ABC 123 .Values ABC",
				pos:  15,
			},
			want: "ABC 123 .Values ABC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getBracetsText(tt.args.line, tt.args.pos)
			assert.Equal(t, tt.want, got)
		})
	}
}
