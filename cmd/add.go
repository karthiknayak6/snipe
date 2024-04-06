/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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

	// addCmd represents the add command
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a new snippet",	
		Long: `Add a new snippet to the snippet manager.
	Specify the language, title, and code snippet.`,
		Run: func(cmd *cobra.Command, args []string) {
			
			if len(args) != 2 {
				fmt.Println("Invalid number of arguments")
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
				log.Panic(err)
			}
			
			code := codeBuffer.String()

			// Syntax highlighting
			lexer := lexers.Get(lan)
			if lexer == nil {
				lexer = lexers.Fallback
			}
			iterator, err := lexer.Tokenise(nil, code)
			if err != nil {
				log.Panic(err)
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
				log.Panic(err)
			}

			fmt.Println("Highlighted Code:")
			fmt.Println(highlightedCode.String())
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

		// Here you will define your flags and configuration settings.

		// Cobra supports Persistent Flags which will work for this command
		// and all subcommands, e.g.:
		// addCmd.PersistentFlags().String("foo", "", "A help for foo")

		// Cobra supports local flags which will only run when this command
		// is called directly, e.g.:
		// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}
