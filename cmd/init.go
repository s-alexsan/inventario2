/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/s-alexsan/inventario2/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Cria a tabela do banco de dados",
	Long:  `Cria a tabela do banco de dados`,
	Run: func(cmd *cobra.Command, args []string) {
		data.CrateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
