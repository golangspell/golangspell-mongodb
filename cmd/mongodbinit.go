package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-mongodb/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["mongodbinit"] = runmongodbinit
}

func runmongodbinit(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(`The command mongodbinit requires 1 argument
		Args:
		DatabaseName: indicates the name of the database to connect with
		
		Syntax: 
		golangspell mongodbinit [DatabaseName]
		
		Examples:
		golangspell mongodbinit mydatabase`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered
	err := usecase.RendermongodbinitTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	fmt.Println("Project initialized with Mongo DB!")
}
