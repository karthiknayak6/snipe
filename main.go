/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/karthiknayak6/snipe/cmd"
	"github.com/karthiknayak6/snipe/database"
)

func main() {
	database.CreateConnection()
	cmd.Execute()
	database.TerminateConnection()
}
