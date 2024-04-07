package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/karthiknayak6/snipe/database"
	"github.com/karthiknayak6/snipe/helpers"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Long: `List all snippets from the 'snippets' collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmd.Help()
			return
		}
		lan, _ := cmd.Flags().GetString("lan")
		head, _ := cmd.Flags().GetBool("head")

		filter := bson.M{}	
		if lan != "" {
			filter = bson.M{"lan": lan}
		} 
		cur, err := database.Db.Collection("snippets").Find(context.TODO(), filter)
		if err != nil {
			log.Panic(err)
		}
		defer cur.Close(context.TODO())

		var snippets []Snippet 

		for cur.Next(context.TODO()) {
			var snippet Snippet
			err := cur.Decode(&snippet)
			if err != nil {
				fmt.Println("Error: ", err)
				return 
			}
			snippets = append(snippets, snippet)
		}

		if err := cur.Err(); err != nil {
			log.Panic(err)
		}

		for _, snippet := range snippets {
			fmt.Printf("%v | %s | %s\n\n", snippet.ID, snippet.Lan, snippet.Title)
			if head {
				if len(snippet.Code) < 350 {
					highlightedCode, err := helpers.HighlightSyntax(snippet.Lan, snippet.Code)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(highlightedCode, ".......")
				} else {

					highlightedCode, err := helpers.HighlightSyntax(snippet.Lan, snippet.Code[:350])
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println(highlightedCode, ".......")
					fmt.Println(	)
				}
			}

		}
	},
}


func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("lan", "l", "", "Filter snippets by language")
	listCmd.Flags().BoolP("head", "y", false, "Display first few characters of code snippet")
}
