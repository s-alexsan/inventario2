package view

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
	Type     string
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if pc.Type == "float" {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New(pc.ErrorMsg)
			}
		} else if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red}}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failer %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %v\n", result)

	return result
}
