package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/golangspell/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["build-config"] = runBuildConfigCommand
}

func runBuildConfigCommand(cmd *cobra.Command, args []string) {
	configBytes, err := json.MarshalIndent(buildSpellConfig(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(configBytes))
}

func buildSpellConfig() domain.Spell {
	return domain.Spell{
		Name: "golangspell-mongodb",
		URL:  "github.com/golangspell/golangspell-mongodb",
		Commands: map[string]*domain.Command{
			"build-config": &domain.Command{
				Name:             "build-config",
				ShortDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool",
				LongDescription: `Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
`,
			},
			"golangspell-mongodb-hello": &domain.Command{
				Name:             "golangspell-mongodb-hello",
				ShortDescription: "The golangspell-mongodb-hello says Hello! using your new Golangspell base structure",
				LongDescription: `The golangspell-mongodb-hello says Hello! using your new Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
You can use this as a template to create your own commands. 
Please notice that ALL your commands must be prefixed with the name of your Spell (golangspell-mongodb). It will avoid name colision with the Spells from other authors 
Args:
name: Your name (required) to be added to the Hello!. Example: Elvis"

Syntax: 
golangspell golangspell-mongodb-hello [name]
`,
				ValidArgs: []string{"name"},
			},
			"mongodbinit": &domain.Command{
				Name:             "mongodbinit",
				ShortDescription: "The mongodbinit initializes the application with the Mongo DB infrastructure",
				LongDescription: `The mongodbinit initializes the application with the Mongo DB infrastructure
Args:
DatabaseName: indicates the name of the database to connect with

Syntax: 
golangspell mongodbinit [DatabaseName]

Examples:
golangspell mongodbinit mydatabase`,
				ValidArgs: []string{"DatabaseName"},
			},
		},
	}
}
