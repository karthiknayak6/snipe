/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/karthiknayak6/snipe/database"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all snippets",
	Long: `List all snippets from the 'snippets' collection.`,
	Run: func(cmd *cobra.Command, args []string) {
		filter := bson.M{}
		cur, err := database.Db.Collection("snippets").Find(context.TODO(), filter)
		if err != nil {
			log.Panic(err)
		}
		defer cur.Close(context.TODO())

		var snippets []Snippet // Assuming Snippet is your struct to represent a snippet

		for cur.Next(context.TODO()) {
			var snippet Snippet
			err := cur.Decode(&snippet)
			if err != nil {
				log.Panic(err)
			}
			snippets = append(snippets, snippet)
		}

		if err := cur.Err(); err != nil {
			log.Panic(err)
		}

		for _, snippet := range snippets {
			fmt.Printf("ID: %v\nLan: %s\nTitle: %s\n\n", snippet.ID	, snippet.Lan, snippet.Title)
		}
	},
}


func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
