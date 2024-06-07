package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	coredomain "github.com/golangspell/golangspell-core/domain"
	coreusecase "github.com/golangspell/golangspell-core/usecase"
	"github.com/golangspell/golangspell-mongodb/appcontext"
	"github.com/golangspell/golangspell-mongodb/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
	"github.com/iancoleman/strcase"
)

func addComponentConstantToContext(currentPath string, componentName string) error {
	return coreusecase.GetAddComponentConstantToContext().Execute(currentPath, componentName)
}

func addImportToMain(moduleName string, currentPath string, importPath string) error {
	return coreusecase.GetAddPackageImportToMain().Execute(moduleName, currentPath, importPath)
}

func renameTemplateFileNames(currentPath string, domainEntity string) error {
	sourcePath := fmt.Sprintf("%s%sdomain%smodel_new_domain.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%smodel_%s.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%scontroller%snew_domain_controller.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_controller.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snew_domain_create_usecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_create_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snew_domain_delete_usecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_delete_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snew_domain_get_all_usecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_get_all_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snew_domain_get_by_id_usecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_get_by_id_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snew_domain_update_usecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_update_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%sgateway%smongodb%snew_domain_repository.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_repository.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	return os.Rename(sourcePath, destinationPath)
}

func addNewRoutes(currentPath string, domainEntity string) error {
	routerFilePath := fmt.Sprintf("%s%scontroller%srouter.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	return new(coredomain.CodeFile).
		ParseFromPath(routerFilePath).
		AddStatementToFunction(
			"MapRoutes",
			fmt.Sprintf("g.GET(\"/%s\", Get%sList)", strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity)),
			func(statementCode string) bool {
				return strings.Contains(statementCode, "GET(\"/info\", GetInfo)")
			}).
		AddStatementToFunction(
			"MapRoutes",
			fmt.Sprintf("g.POST(\"/%s\", Create%s)", strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity)),
			func(statementCode string) bool {
				return strings.Contains(statementCode, fmt.Sprintf("g.GET(\"/%s\", Get%sList)", strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity)))
			}).
		AddStatementToFunction(
			"MapRoutes",
			fmt.Sprintf("g.GET(\"/%s/:%sId\", Get%s)", strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)),
			func(statementCode string) bool {
				return strings.Contains(statementCode, fmt.Sprintf("g.POST(\"/%s\", Create%s)", strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity)))
			}).
		AddStatementToFunction(
			"MapRoutes",
			fmt.Sprintf("g.PUT(\"/%s/:%sId\", Update%s)", strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)),
			func(statementCode string) bool {
				return strings.Contains(statementCode, fmt.Sprintf("g.GET(\"/%s/:%sId\", Get%s)", strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)))
			}).
		AddStatementToFunction(
			"MapRoutes",
			fmt.Sprintf("g.DELETE(\"/%s/:%sId\", Delete%s)", strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)),
			func(statementCode string) bool {
				return strings.Contains(statementCode, fmt.Sprintf("g.PUT(\"/%s/:%sId\", Update%s)", strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)))
			}).Save()
}

// RendermongodbaddcrudTemplate renders the templates defined to the mongodbaddcrud command with the proper variables
func RendermongodbaddcrudTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	domainEntity := strcase.ToCamel(args[0])
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to add a CRUD for the domain %s. Error: %s\n", domainEntity, err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"DomainEntity":           domainEntity,
		"DomainEntityLowerCamel": strcase.ToLowerCamel(args[0]),
		"DomainEntityLower":      strings.ToLower(args[0]),
		"ModuleName":             moduleName,
	}

	err = renderer.RenderTemplate(spell, "mongodbaddcrud", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sRepository", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sCreateUsecase", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sGetAllUsecase", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sGetByIDUsecase", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sUpdateUsecase", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sDeleteUsecase", domainEntity))
	if err != nil {
		fmt.Println(err.Error())
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

	err = addImportToMain(
		moduleName,
		currentPath,
		fmt.Sprintf("%s/usecase", moduleName))
	if err != nil {
		fmt.Printf("an error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}

	err = renameTemplateFileNames(currentPath, domainEntity)
	if err != nil {
		fmt.Printf("an error occurred while trying to rename the rendered template files. Error: %s\n", err.Error())
		return err
	}

	err = addNewRoutes(currentPath, domainEntity)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
