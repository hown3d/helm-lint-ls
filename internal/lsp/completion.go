package lsp

import (
	"fmt"
	"strings"

	"github.com/mrjosh/helm-lint-ls/internal/util"
	lsp "go.lsp.dev/protocol"
)

var openAndReadFile = util.ReadFileAtLine

const (
	CompletionKindValues  = "Values"
	CompletionKindRelease = "Release"
	CompletionKindFile    = "File"
	CompletionKindChart   = "Chart"
)

func completionKind(line string) {

}

func completeValue(filepath string, pos lsp.Position) ([]lsp.CompletionItem, error) {
	line, err := openAndReadFile(filepath, pos.Line)
	if err != nil {
		return nil, fmt.Errorf("reading file %v at line %v: %w", filepath, pos.Line, err)
	}
	// get the strings inside the current {{ }}
	insideBracets := strings.TrimSpace(getBracetsText(line, int(pos.Character)))

	// {{ .Values.data.test.<CURSOR> }}
	logger.Debug(insideBracets)
	return []lsp.CompletionItem{{
		Label: insideBracets,
	}}, nil
}

func getBracetsText(line string, pos int) string {
	leftDelimiter := 0
	rightDelimiter := len(line)
	for i := pos; i >= 0; i-- {
		if line[i] == '{' && line[i-1] == '{' {
			leftDelimiter = i - 1
			break
		}
	}
	for i := pos; i < len(line); i++ {
		if line[i] == '}' && line[i+1] == '}' {
			rightDelimiter = i + 1
			break
		}
	}
	// to prevent out of bounds
	if rightDelimiter == len(line) {
		return line[leftDelimiter:rightDelimiter]
	}
	// +1 to include the bracets
	return line[leftDelimiter : rightDelimiter+1]
}
