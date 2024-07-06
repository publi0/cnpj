package pkg

import (
	"regexp"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	isRaw        bool
	IsFull       bool
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "yaml", "Output format (yaml or json)")
	RootCmd.PersistentFlags().BoolVarP(&isRaw, "raw", "r", false, "Print raw output")
	RootCmd.PersistentFlags().BoolVarP(&IsFull, "full", "f", false, "Print full output")
}

var RootCmd = &cobra.Command{
	Use:   "cnpj [cnpj]",
	Short: "Consulte o CNPJ de uma empresa.",
	Long:  "Consulte o CNPJ de uma empresa e obtenha informações detalhadas sobre a empresa.",
	Args:  cobra.ExactArgs(1),
	Run:   fetchCNPJ,
}

func fetchCNPJ(cmd *cobra.Command, args []string) {
	var data any
	cnpj := args[0]

	reg := regexp.MustCompile("[^0-9]+")
	cnpj = reg.ReplaceAllString(cnpj, "")

	if cnpj == "" {
		cmd.Println("CNPJ não pode ser vazio.")
		return
	}

	data, err := fetchCNPJData(cmd.Context(), cnpj)
	if err != nil {
		panic(err)
	}

	if !IsFull {
		data = SimplifyCNPJData(data.(CNPJData))
	}

	switch outputFormat {
	case "json":
		PrintJSON(data, isRaw)
	case "yaml":
		PrintYAML(data, isRaw)
	default:
		cmd.Println("Unsupported output format. Please choose 'json' or 'yaml'.")
	}
}
