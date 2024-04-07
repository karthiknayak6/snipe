package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/karthiknayak6/snipe/database"
	"github.com/karthiknayak6/snipe/helpers"
	"github.com/spf13/cobra"
)

	type Snippet struct {
		ID    int 				 `bson:"_id,omitempty"`
		Lan   string             `bson:"lan,omitempty"`
		Title string             `bson:"title,omitempty"`
		Code  string             `bson:"code,omitempty"`
	}

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new snippet",	
		Long: `Add a new snippet to the snippet manager.
	Specify the language, title, and code snippet.`,
		Run: func(cmd *cobra.Command, args []string) {
			
			if len(args) != 2 {
				cmd.Help()
				return
			}
			lan := args[0]
			title := args[1]
			var found bool
			for _, l := range helpers.Lan_list {
				if l == lan {
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Invalid Programming Language %v\n\n", lan)
				helpers.PrintSupportedLanguages()
				return
			}
		
			
			fmt.Println("Enter your code snippet (press Ctrl+D on a new line to finish):")
			
			var codeBuffer bytes.Buffer
			scanner := bufio.NewScanner(os.Stdin)
			
			for scanner.Scan() {
				line := scanner.Text()
				if line == "Ctrl+D" {
					break
				}
				codeBuffer.WriteString(line)
				codeBuffer.WriteString("\n")
			}
			
			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				return
			}
			
			code := codeBuffer.String()

		

			fmt.Println("Your Code:")
			highlightedCode, err := helpers.HighlightSyntax(lan, code)
			if err != nil {
				panic(err)
			}
			fmt.Println(highlightedCode)


			seq, err := database.GetNextSequence(database.Client, "snippets")
			if err != nil {
				log.Fatal(err)
			}

			snippet := Snippet{ID: seq ,Lan: lan, Title: title, Code: code}
			
			res, err := database.Db.Collection("snippets").InsertOne(context.TODO(), snippet)
			
			if err != nil {
				log.Panic(err)
			}
			fmt.Println("Hurry! Your snippet has been added with ID:", res.InsertedID)
		},
	}


	func init() {
		rootCmd.AddCommand(addCmd)
	}
