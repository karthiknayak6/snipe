package helpers

import (
	"fmt"
	"slices"
	"strings"
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