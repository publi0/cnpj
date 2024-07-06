package main

import (
	"cnpj/pkg"
	"fmt"
	"os"
)

func main() {
	err := pkg.RootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
