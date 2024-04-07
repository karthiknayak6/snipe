package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/karthiknayak6/snipe/database"
	"github.com/karthiknayak6/snipe/helpers"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a snippet",
	Long: `Rewrite the code snippet in the database by its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID format")
			return
		}
		res := database.Db.Collection("snippets").FindOne(context.TODO(), bson.M{"_id": id})
		if res.Err() != nil{
			fmt.Println("No snippet found with ID ", id)
		}
		var snippet Snippet
		res.Decode(&snippet)
		fmt.Printf("%v | %s | %s\n\n", snippet.ID, snippet.Lan, snippet.Title)
		highlightedCode, err := helpers.HighlightSyntax(snippet.Lan, snippet.Code)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(highlightedCode)
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
		filter := bson.M{"_id": id}
		update := bson.M{"$set": bson.M{"code": code}}
		res = database.Db.Collection("snippets").FindOneAndUpdate(context.TODO(), filter, update)
		if res.Err() != nil {
			fmt.Println("Error: ", res.Err())
			return
		}
		fmt.Println("Snippet updated successfully")
		
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

}
