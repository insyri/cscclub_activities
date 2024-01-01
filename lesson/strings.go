package lesson

import (
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"strings"
)

func printPrompts(prompts []string) {
	for i := range prompts {
		if prompts[i] == "" {
			continue
		}

		n := "\n"
		if i == 0 {
			n = ""
		}
		fmt.Println(n + wordwrap.WrapString(prompts[i], WordWrapLimit))
	}
}

func orderedList(prompts []string) string {
	var s strings.Builder
	for i := range prompts {
		// after wrapping, add indents to any places where it wrapped
		x := wordwrap.WrapString(
			fmt.Sprintf("    %d. %s", i+1, prompts[i]),
			WordWrapLimit)
		x = strings.ReplaceAll(x, "\n", "\n       ") + "\n"
		s.WriteString(x)
	}
	return strings.TrimSuffix(s.String(), "\n")
}

func conditional(f bool, prompt string) string {
	if f {
		return prompt
	}
	return ""
}
