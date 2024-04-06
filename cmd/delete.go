package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/karthiknayak6/snipe/database"
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
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID format")
			return
		}

		filter := bson.M{"_id": id}
		result, err := database.Db.Collection("snippets").DeleteOne(context.TODO(), filter)
		if err != nil {
			fmt.Println(err)
			return
		}

		if result.DeletedCount == 0 {
			fmt.Println("No document found to delete")
			return
		}

		fmt.Println("Successfully deleted the snippet")
	},
}


func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
