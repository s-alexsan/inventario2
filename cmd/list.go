/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/s-alexsan/inventario2/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todos os produtos do inventário",
	Long:  `Lista todos os produtos do inventário`,
	Run: func(cmd *cobra.Command, args []string) {
		sql := "SELECT * FROM inventario"
		data.Select(sql)
	},
}

func init() {
	produtosCmd.AddCommand(listCmd)
}
