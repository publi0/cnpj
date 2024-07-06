package pkg

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/printer"
	"github.com/mattn/go-colorable"
)

const (
	defaultIndent = 4
	escape        = "\x1b"
)

func PrintYAML(value any, isRaw bool) error {
	yamlBytes, err := yaml.MarshalWithOptions(value, yaml.Indent(defaultIndent))
	if err != nil {
		return err
	}

	if isRaw {
		_, err = fmt.Println(string(yamlBytes))
		return err
	}

	return printWithColor(string(yamlBytes))
}

func format(attr color.Attribute) string {
	return fmt.Sprintf("%s[%dm", escape, attr)
}

func printWithColor(output string) error {
	tokens := lexer.Tokenize(output)
	var p printer.Printer
	p.Bool = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiMagenta),
			Suffix: format(color.Reset),
		}
	}
	p.Number = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiMagenta),
			Suffix: format(color.Reset),
		}
	}
	p.MapKey = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiCyan),
			Suffix: format(color.Reset),
		}
	}
	p.Anchor = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiYellow),
			Suffix: format(color.Reset),
		}
	}
	p.Alias = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiYellow),
			Suffix: format(color.Reset),
		}
	}
	p.String = func() *printer.Property {
		return &printer.Property{
			Prefix: format(color.FgHiGreen),
			Suffix: format(color.Reset),
		}
	}
	writer := colorable.NewColorableStdout()
	if _, err := writer.Write([]byte(p.PrintTokens(tokens) + "\n")); err != nil {
		return err
	}
	return nil
}
