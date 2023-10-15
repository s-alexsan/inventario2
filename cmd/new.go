/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/s-alexsan/inventario2/data"
	"github.com/s-alexsan/inventario2/view"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Cria um novo produto do inventario",
	Long:  `Cria um novo produto do inventario`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewProduct()
	},
}

func init() {
	produtosCmd.AddCommand(newCmd)
}

func createNewProduct() {
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

	fmt.Println(produto)
	fmt.Println(categoria)
	fmt.Println(preco)
	fmt.Println(quantidade)

	sql := "INSERT INTO inventario(produto, categoria, preco, quantidade) VALUES (?, ?, ?, ?)"
	data.Query(sql, produto, categoria, preco, quantidade)

}
