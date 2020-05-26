package main

import (
	"fmt"
	
	"github.com/golangspell/golangspell-mongodb/cmd"
	_ "github.com/golangspell/golangspell-mongodb/config"
	_ "github.com/golangspell/golangspell-mongodb/gateway/template"
	_ "github.com/golangspell/golangspell-mongodb/gateway/customlog"

)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
