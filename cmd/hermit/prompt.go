package main

import (
	"os"
	"os/user"
	"strings"
)

type PromptConfig struct {
	ShowCwd bool
}

type Prompt struct {
	Parts []PromptPart
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (p *Prompt) AddPart(part PromptPart) {
	p.Parts = append(p.Parts, part)
}

func (p *Prompt) AddTextPart(text string, color Color) {
	p.AddPart(&TextPart{Text: text, Color: color})
}

type PromptPart interface {
	String() string
}

type TextPart struct {
	Text  string
	Color Color
}

func (t *TextPart) String() string {
	return t.Color.Sprint(t.Text)
}

func (p *Prompt) String() string {
	var result string
	for _, part := range p.Parts {
		result += part.String()
	}
	return result
}

func getPrompt(config PromptConfig) string {
	prompt := &Prompt{}

	user, _ := user.Current()
	cwd, _ := os.Getwd()
	homeDir := user.HomeDir

	if strings.HasPrefix(cwd, homeDir) {
		cwd = "~" + cwd[len(homeDir):]
	} else {
		parts := strings.Split(cwd, string(os.PathSeparator))
		if len(parts) > 3 {
			cwd = "..." + string(os.PathSeparator) + strings.Join(parts[len(parts)-2:], string(os.PathSeparator))
		}
	}

	prompt.AddTextPart(user.Username, ColorGreen)
	prompt.AddTextPart("@", ColorWhite)
	prompt.AddTextPart("hermit", ColorCyan)
	prompt.AddTextPart(":", ColorWhite)
	if config.ShowCwd {
		prompt.AddTextPart(cwd, ColorBlue)
	} else {
		prompt.AddTextPart("~", ColorBlue)
	}
	prompt.AddTextPart("$ ", ColorWhite)

	return prompt.String()
}
