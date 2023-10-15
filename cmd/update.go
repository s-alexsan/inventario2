/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/s-alexsan/inventario2/data"
	"github.com/s-alexsan/inventario2/view"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza os dados do produto",
	Long:  "Atualiza os dados do produto",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		if id == "" {
			log.Println("Informe o id")
			os.Exit(1)
		}

		updateProduct(id)
	},
}

func init() {
	produtosCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().String("id", "", "Código ID do produto para atualizar")
}

func updateProduct(id string) {
	produtoPromptContent := view.PromptContent{
		ErrorMsg: "Informe o produto corretamente.",
		Label:    "Informe o produto: ",
	}

	produto := view.PromptGetInput(produtoPromptContent)

	categoriaPromptContent := view.PromptContent{
		ErrorMsg: "Informe a categoria corretamente.",
		Label:    fmt.Sprintf("Informe a categoria do %v: ", produto),
	}

	categoria := view.PromptGetInput(categoriaPromptContent)

	precoPromptContent := view.PromptContent{
		ErrorMsg: "Número inválido",
		Label:    fmt.Sprintf("Informe o preco do %v: ", produto),
		Type:     "float",
	}

	preco, _ := strconv.ParseFloat(view.PromptGetInput(precoPromptContent), 64)

	quantidadePromptContent := view.PromptContent{
		ErrorMsg: "Número inválido",
		Label:    fmt.Sprintf("Informe a quantidae do %v: ", produto),
		Type:     "float",
	}

	quantidade, _ := strconv.ParseFloat(view.PromptGetInput(quantidadePromptContent), 64)

	// fmt.Println(produto)
	// fmt.Println(categoria)
	// fmt.Println(preco)
	// fmt.Println(quantidade)

	sql := "UPDATE inventario SET produto = ?, categoria = ?, preco = ?, quantidade = ? WHERE id = ?"
	data.Query(sql, produto, categoria, preco, quantidade, id)

}
