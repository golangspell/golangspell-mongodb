package usecase

import (
	"fmt"
	"os"

	coreusecase "github.com/golangspell/golangspell-core/usecase"
	"github.com/golangspell/golangspell-mongodb/appcontext"
	"github.com/golangspell/golangspell-mongodb/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
)

func addEnvironmentVariables(currentPath string) error {
	err := coreusecase.GetAddEnvironmentVariable().Execute(currentPath, "DBConnectionString", "string", "`env:\"DB_CONNECTION_STRING\" envDefault:\"\"`")
	if err != nil {
		fmt.Printf("An error occurred while trying to update the environment file. Error: %s\n", err.Error())
		return err
	}
	return coreusecase.GetAddEnvironmentVariable().Execute(currentPath, "DBConnectionCertificateFileName", "string", "`env:\"DB_CONNECTION_CERTIFICATE_FILE_NAME\" envDefault:\"\"`")
}

func addClientToContext(currentPath string) error {
	return coreusecase.GetAddComponentConstantToContext().Execute(currentPath, "DBClient")
}

// RendermongodbinitTemplate renders the templates defined to the mongodbinit command with the proper variables
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

	err = addEnvironmentVariables(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the environment. Error: %s\n", err.Error())
		return err
	}
	err = addClientToContext(currentPath)
	if err != nil {
		fmt.Printf("An error occurred while trying to configure the context. Error: %s\n", err.Error())
		return err
	}
	err = addImportToMain(
		moduleName,
		currentPath,
		fmt.Sprintf("%s/gateway/mongodb", moduleName))
	if err != nil {
		fmt.Printf("an error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}
	return renderer.RenderTemplate(spell, "mongodbinit", globalVariables, nil)
}
