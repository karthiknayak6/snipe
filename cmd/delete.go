package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/karthiknayak6/snipe/database"
	"github.com/karthiknayak6/snipe/helpers"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snippet",
	Long: `Delete a snippet from the database by its ID.
For example:

delete <snippet_id>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		idStr := args[0]
		id, err := strconv.Atoi(idStr); 
		if err != nil {
			fmt.Println("Invalid ID format")
			return
		}
		result := database.Db.Collection("snippets").FindOne(context.TODO(), bson.M{"_id": id})
		if result.Err() != nil {
			fmt.Println("No Snippet is found with this ID ", id)
			return
		}
		var snippet Snippet
		result.Decode(&snippet)
	
		highlightedCode, err := helpers.HighlightSyntax(snippet.Lan, snippet.Code)
		if err != nil {
			fmt.Println(err)
			return	
		}
		fmt.Println(highlightedCode)
		fmt.Printf("%v | %s | %s	\n\nAre you sure you want to delete this snippet? (y/n) ",snippet.ID, snippet.Lan, snippet.Title)		
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')		
		if text == "n\n" || text == "N\n" {
			return
		} else if text != "y\n" && text != "Y\n" {
			fmt.Println("Invalid input")
			return		
		}
		filter := bson.M{"_id": id}
		del, err := database.Db.Collection("snippets").DeleteOne(context.TODO(), filter)
		if err != nil {
			fmt.Println(err)
			return
		}

		if del.DeletedCount == 0 {
			fmt.Println("No document found to delete")
			return
		}

		fmt.Println("\nSuccessfully deleted the snippet")
	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)
}
