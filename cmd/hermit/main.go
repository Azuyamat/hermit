package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/azuyamat/hermit/internal/ast"
	"github.com/azuyamat/hermit/internal/executor"
	"github.com/azuyamat/hermit/internal/lexer"
	"github.com/azuyamat/hermit/internal/parser"
	"github.com/azuyamat/hermit/internal/token"
	"github.com/azuyamat/hermit/internal/types"
)

func main() {
	debugFlag := flag.Bool("debug", false, "print lexer tokens and parse AST")
	flag.Parse()

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

		start := time.Now()
		l := lexer.New(input)

		if *debugFlag {
			printTokens(input)
		}

		p := parser.New(l)
		program, err := p.Parse()
		if err != nil {
			printError(err)
			continue
		}

		if *debugFlag {
			printAST(program)
		}

		exec := executor.New()
		err = exec.Execute(program)
		if err != nil && !types.IsErrExitCode(err) {
			printError(err)
		}
		elapsed := time.Since(start)
		fmt.Printf("Execution time: %.4f ms\n", elapsed.Seconds()*1000)
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "%s", ColorRed.Errorf("[ERROR] %v\n", err))
}

func printTokens(input string) {
	fmt.Println("\n=== LEXER TOKENS ===")
	l := lexer.New(input)
	for {
		tok, err := l.NextToken()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Lexer error: %v\n", err)
			break
		}
		fmt.Printf("Type: %-20s Literal: %-15s Line: %d Column: %d\n",
			tok.Type, fmt.Sprintf("%q", tok.Literal), tok.LineNumber, tok.ColumnNumber)
		if tok.Type == token.EOF {
			break
		}
	}
	fmt.Println("====================")
	fmt.Println()
}

func printAST(program *ast.Program) {
	fmt.Println("\n=== PARSE AST ===")
	fmt.Println(program.String())
	fmt.Println("=================")
	fmt.Println()
}
