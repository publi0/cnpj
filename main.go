package main

import (
	"fmt"
	"os"

	"github.com/publi0/cnpj/pkg"
)

func main() {
	err := pkg.RootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
