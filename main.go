/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/s-alexsan/inventario2/cmd"
	"github.com/s-alexsan/inventario2/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
