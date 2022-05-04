package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

func ReadFileAtLine(filepath string, linenumber uint32) (string, error) {
	data, err := ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return readLine(bytes.NewBuffer(data), linenumber)
}

func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading file %v: %w", filepath, err)
	}
	return data, nil
}

func FindFile(directory string, name string) (files []string, err error) {
	pattern := filepath.Join(directory, "**", name)
	return filepath.Glob(pattern)
}
func readLine(r io.Reader, lineNumber uint32) (string, error) {
	sc := bufio.NewScanner(r)
	var currentLine uint32
	for sc.Scan() {
		currentLine++
		if currentLine == lineNumber {
			return sc.Text(), sc.Err()
		}
	}
	return "", io.EOF
}
