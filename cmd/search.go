package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/karthiknayak6/snipe/database"
	"github.com/karthiknayak6/snipe/helpers"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for snippets",
	Long: `search for snippets using ID or substring`,
	Run: func(cmd *cobra.Command, args []string) {
		lan, _ := cmd.Flags().GetString("lan")
		head, _ := cmd.Flags().GetBool("head")
		if len(args) != 1 {
				cmd.Help()
				return
		}
		
		search_string := args[0]
		search_id, err := strconv.Atoi(search_string)
		var cur *mongo.Cursor
		
		if err != nil {
			//handle title
			filter := bson.M{"title": bson.M{"$regex": search_string, "$options": "i"}}
			if lan != "" {
				filter = bson.M{"title": search_string, "lan": lan}	
			}
			cur, err = database.Db.Collection("snippets").Find(context.TODO(), filter)
		} else {
			//handle ID
			filter := bson.M{"_id":search_id }		
			if lan != "" {
				filter = bson.M{"_id": search_string, "lan": lan}	
			}
			cur, err = database.Db.Collection("snippets").Find(context.TODO(), filter)
		}
		if err != nil {
			fmt.Println("Error: ", err)
			return
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
			fmt.Println("Error: ", err)
			return
		}
		if len(snippets) == 0 {
			fmt.Println("No snippets found")
			return
		}
		if search_id != 0 {
			fmt.Printf("%v | %s | %s\n\n", snippets[0].ID	, snippets[0].Lan, snippets[0].Title)
			highlightedCode, err := helpers.HighlightSyntax(snippets[0].Lan, snippets[0].Code)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(highlightedCode)
			fmt.Println()
			return
		}
		for _, snippet := range snippets {
			fmt.Printf("%v | %s | %s\n\n", snippet.ID	, snippet.Lan, snippet.Title)
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
					fmt.Println()
					return
				}
			}

		
	
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("lan", "l", "", "Filter snippets by language")
	searchCmd.Flags().BoolP("head", "y", false, "Display first few characters of code snippet")
}


