package usecase

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golangspell/golangspell-mongodb/appcontext"
	"github.com/golangspell/golangspell-mongodb/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
)

func addEnvironmentVariable(currentPath string) error {
	filePath := fmt.Sprintf("%s%sconfig%senvironment.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to backup the environment file. Error: %s\n", err.Error())
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read the environment file. Error: %s\n", err.Error())
		return err
	}
	code := strings.ReplaceAll(
		string(content),
		"type Config struct {\n",
		"type Config struct {\n//DBConnectionString to connect to Mongo\nDBConnectionString string\n")
	code = strings.ReplaceAll(
		code,
		"func init() {\n",
		"func init() {\n_ = viper.BindEnv(\"DBConnectionString\", \"DB_CONNECTION_STRING\")\n")
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("An error occurred while trying to update the environment file. Error: %s\n", err.Error())
		return err
	}

	return nil
}

func addClientToContext(currentPath string) error {
	filePath := fmt.Sprintf("%s%sappcontext%scontext.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to backup the context file. Error: %s\n", err.Error())
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to read the context file. Error: %s\n", err.Error())
		return err
	}
	code := strings.ReplaceAll(
		string(content),
		"const (\n",
		"const (\nDBClient = \"DBClient\"\n")
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("An error occurred while trying to update the context file. Error: %s\n", err.Error())
		return err
	}

	return nil
}

//RendermongodbinitTemplate renders the templates defined to the mongodbinit command with the proper variables
func RendermongodbinitTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to initialize the MongoDB infrastructure. Error: %s\n", err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"DatabaseName": args[0],
		"ModuleName":   moduleName,
	}

	err = addEnvironmentVariable(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the environment. Error: %s\n", err.Error())
		return err
	}
	err = addClientToContext(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the context. Error: %s\n", err.Error())
		return err
	}
	return renderer.RenderTemplate(spell, "mongodbinit", globalVariables, nil)
}
