package colorStationsJson

import (
	"fmt"

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

func (f *Formatter) Marshal(jsonObj interface{}) string { //([]byte, error) {
	switch v := jsonObj.(type) {
	case map[string]interface{}:
		return f.marshalMap(v)
	}
	return "Must be type map[string]interface{}"
	//return buffer.Bytes(), nil
}

func (f *Formatter) marshalMap(m map[string]interface{}) string {

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	var tble = table.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("100")))
	for _, key := range keys {

		tble.Row(key)
		switch v := m[key].(type) {

		case map[string]interface{}:
			var tble = table.New().
				BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")))
			tble.Row(f.marshalMap(v))
			return tble.Render()
		case []interface{}:
			if len(v) != 0 {
				return f.marshalArray(v, key)
			}
		}

	}
	return tble.Render()
}

func (f *Formatter) marshalArray(a []interface{}, key string) string {
	var arrayTable = table.New().
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99")))
	arrayTable.Row(key)
	for _, v := range a {

		if v.(map[string]interface{}) != nil {
			arrayTable.Row(f.marshalMap(v.(map[string]interface{})))
		} else {
			fmt.Println("Not a map")
		}

	}
	return arrayTable.Render()

}

func (f *Formatter) marshalValue(val interface{}) string {

	return ""

}

// Marshal JSON data with default options
func Marshal(jsonObj interface{}) string { //([]byte, error) {
	return NewFormatter().Marshal(jsonObj)
}
