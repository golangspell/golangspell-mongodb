package usecase

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golangspell/golangspell-mongodb/appcontext"
	"github.com/golangspell/golangspell-mongodb/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
	"github.com/iancoleman/strcase"
)

func addComponentConstantToContext(currentPath string, constantDefinition string) error {
	filePath := fmt.Sprintf("%s%sappcontext%scontext.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to backup the context file. Error: %s", err.Error())
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to read the context file. Error: %s", err.Error())
	}
	code := strings.ReplaceAll(
		string(content),
		"const (\n",
		fmt.Sprintf("const (\n%s\n", constantDefinition))
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to update the context file. Error: %s", err.Error())
	}

	return nil
}

func addImportToMain(currentPath string, importPath string) error {
	filePath := fmt.Sprintf("%s%smain.go", currentPath, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to backup the main file. Error: %s", err.Error())
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to read the main file. Error: %s", err.Error())
	}

	if strings.Contains(string(content), importPath) {
		return nil
	}

	code := strings.ReplaceAll(
		string(content),
		"/config\"\n",
		fmt.Sprintf("/config\"\n_ \"%s\"\n", importPath))
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to update the main file. Error: %s", err.Error())
	}

	return nil
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
	filePath := fmt.Sprintf("%s%scontroller%srouter.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to backup the context file. Error: %s", err.Error())
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to read the context file. Error: %s", err.Error())
	}
	code := strings.ReplaceAll(
		string(content),
		"g.GET(\"/info\", GetInfo)\n",
		fmt.Sprintf("g.GET(\"/info\", GetInfo)\ng.GET(\"/%s\", Get%sList)\ng.POST(\"/%s\", Create%s)\ng.GET(\"/%s/:%sId\", Get%s)\ng.PUT(\"/%s/:%sId\", Update%s)\ng.DELETE(\"/%s/:%sId\", Delete%s)\n",
			strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity),
			strcase.ToKebab(domainEntity), strcase.ToCamel(domainEntity),
			strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity),
			strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity),
			strcase.ToKebab(domainEntity), strcase.ToLowerCamel(domainEntity), strcase.ToCamel(domainEntity)))
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return fmt.Errorf("An error occurred while trying to add the new REST routes. Error: %s", err.Error())
	}

	return nil
}

//RendermongodbaddcrudTemplate renders the templates defined to the mongodbaddcrud command with the proper variables
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
		fmt.Sprintf("%sRepository = \"%sRepository\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sCreateUsecase = \"%sCreateUsecase\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sGetAllUsecase = \"%sGetAllUsecase\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sGetByIDUsecase = \"%sGetByIDUsecase\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sUpdateUsecase = \"%sUpdateUsecase\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addComponentConstantToContext(
		currentPath,
		fmt.Sprintf("%sDeleteUsecase = \"%sDeleteUsecase\"", domainEntity, domainEntity))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = addImportToMain(
		currentPath,
		fmt.Sprintf("%s/gateway/mongodb", moduleName))
	if err != nil {
		fmt.Printf("An error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}

	err = addImportToMain(
		currentPath,
		fmt.Sprintf("%s/usecase", moduleName))
	if err != nil {
		fmt.Printf("An error occurred while trying to add the import to main. Error: %s\n", err.Error())
		return err
	}

	err = renameTemplateFileNames(currentPath, domainEntity)
	if err != nil {
		fmt.Printf("An error occurred while trying to rename the rendered template files. Error: %s\n", err.Error())
		return err
	}

	err = addNewRoutes(currentPath, domainEntity)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
