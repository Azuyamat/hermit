package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/azuyamat/hermit/internal/executor"
	"github.com/azuyamat/hermit/internal/lexer"
	"github.com/azuyamat/hermit/internal/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hermit Shell v0.1.0")
	fmt.Println("Type 'exit' to quit.")

	promptConfig := PromptConfig{
		ShowCwd: true,
	}

	for {
		fmt.Print(getPrompt(promptConfig))
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "" {
			continue
		}

		l := lexer.New(input)
		p := parser.New(l)
		program, err := p.Parse()
		if err != nil {
			printError(err)
			continue
		}

		executor := executor.New()
		err = executor.Execute(program)
		if err != nil {
			printError(err)
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s", ColorRed.Errorf("[ERROR] %v\n", err))
}
