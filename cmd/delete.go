/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/s-alexsan/inventario2/data"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Apaga o produto do inventario pelo ID",
	Long:  `Apaga o produto do inventario pelo ID`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("d")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if id == "" {
			log.Println("informe o ID do produto")
		}

		sql := "DELETE FROM inventario WHERE id = ?"

		data.Query(sql, id)
	},
}

func init() {
	produtosCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().String("d", "", "parametro para apagar o produto do inventario")

}
