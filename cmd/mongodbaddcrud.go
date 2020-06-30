package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-mongodb/usecase"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["mongodbaddcrud"] = runmongodbaddcrud
}

func runmongodbaddcrud(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(`The command mongodbaddcrud requires 1 argument
Args:
domainentity: The name (camel case) of the new domain entity which will be generated with the respective controller, usecase and repository

Syntax: 
golangspell mongodbaddcrud [YourNewEntityName]

Examples:
golangspell mongodbaddcrud Product`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered
	err := usecase.RendermongodbaddcrudTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	fmt.Println("mongodbaddcrud executed!")
	fmt.Printf("CRUD created for the domain entity %s\n", strcase.ToCamel(args[0]))
}
