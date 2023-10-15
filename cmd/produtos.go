/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// produtosCmd represents the produtos command
var produtosCmd = &cobra.Command{
	Use:   "produtos",
	Short: "Acessa as ações do CRUD",
	Long:  `Acessa as ações do CRUD`,
}

func init() {
	rootCmd.AddCommand(produtosCmd)

}
