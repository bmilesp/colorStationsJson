package colorStationsJson

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/fatih/color"
)

type Formatter struct {
	KeyColor        *color.Color
	StringColor     *color.Color
	BoolColor       *color.Color
	NumberColor     *color.Color
	NullColor       *color.Color
	StringMaxLength int
	Indent          int
	DisabledColor   bool
	RawStrings      bool
}

func NewFormatter() *Formatter {
	return &Formatter{
		KeyColor:        color.New(color.FgWhite),
		StringColor:     color.New(color.FgGreen),
		BoolColor:       color.New(color.FgYellow),
		NumberColor:     color.New(color.FgCyan),
		NullColor:       color.New(color.FgMagenta),
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          0,
		RawStrings:      false,
	}
}

func (f *Formatter) sprintfColor(c *color.Color, format string, args ...interface{}) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprintf(format, args...)
	}
	return c.SprintfFunc()(format, args...)
}

func (f *Formatter) sprintColor(c *color.Color, s string) string {
	if f.DisabledColor || c == nil {
		return fmt.Sprint(s)
	}
	return c.SprintFunc()(s)
}

func (f *Formatter) writeIndent(buf *bytes.Buffer, depth int) {
	buf.WriteString(strings.Repeat(" ", f.Indent*depth))
}

func (f *Formatter) writeObjSep(buf *bytes.Buffer) {
	if f.Indent != 0 {
		buf.WriteByte('\n')
	} else {
		buf.WriteByte(' ')
	}
}

func (f *Formatter) Marshal(jsonObj interface{}) string { //([]byte, error) {

	return f.marshalValue(jsonObj, 0)
	//return buffer.Bytes(), nil
}

func (f *Formatter) marshalMap(m map[string]interface{}, depth int) string {

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	var tble = table.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("100")))
	for _, key := range keys {

		tble.Row(key)
		switch m[key].(type) {

		case map[string]interface{}:
			tble.Row(f.marshalValue(m[key], depth+1))
		}

	}
	return tble.Render()
}

func (f *Formatter) marshalValue(val interface{}, depth int) string {
	switch v := val.(type) {
	case map[string]interface{}:
		var tble = table.New().
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")))
		tble.Row(f.marshalMap(v, depth))
		return tble.Render()
	}
	return ""

}

// Marshal JSON data with default options
func Marshal(jsonObj interface{}) string { //([]byte, error) {
	return NewFormatter().Marshal(jsonObj)
}
