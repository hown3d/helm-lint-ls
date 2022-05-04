package helm

import (
	"fmt"
	"text/template/parse"
)

func ParseTemplate(templateText string) error {
	trees, err := parse.Parse("test", templateText, "{{", "}}")
	if err != nil {
		return err
	}
	fmt.Printf("%+v", trees)
	return nil
}
