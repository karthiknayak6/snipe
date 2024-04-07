package helpers

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

func PrintSupportedLanguages() {
			fmt.Printf("Select from:\n")
			slices.Sort(Lan_list)
			
			// Find the maximum length of a language name in lan_list
			maxLen := 0
			for _, l := range Lan_list {
				if len(l) > maxLen {
					maxLen = len(l)
				}
			}
			
			// Calculate the indentation size
			indentSize := maxLen + 4  // 4 extra spaces for padding
			
			for k, l := range Lan_list {
				if k%5 == 0 {
					fmt.Println()
				}
				
				// Calculate the number of spaces needed for indentation
				numSpaces := indentSize - len(l)
				
				fmt.Printf("%s%s", l, strings.Repeat(" ", numSpaces))
			}
			fmt.Println()
}


func HighlightSyntax(lan string, code string) (string, error) {
	lexer := lexers.Get(lan)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		return "", err
	}

	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	formatter := formatters.Get("terminal16m")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	var highlightedCode bytes.Buffer
	err = formatter.Format(&highlightedCode, style, iterator)
	if err != nil {
		return "", err
	}
	return highlightedCode.String(), nil
}