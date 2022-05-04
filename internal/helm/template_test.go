package helm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTemplate = `
{{- if .Values.serviceAccount.create -}}
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ include "test-chart.serviceAccountName" . }}
  labels:
    {{- include "test-chart.labels" . | nindent 4 }}
  {{- with .Values.serviceAccount.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
{{- end }}
`

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseTemplate(testTemplate)
			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
