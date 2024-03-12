package shell

import (
	"strings"

	"github.com/c-bata/go-prompt"

	"github.com/knackwurstking/picow-led/internal/command"
)

func completer(d prompt.Document) []prompt.Suggest {
	sub := strings.TrimLeft(d.Text, " ")
	s := []prompt.Suggest{
		{Text: "config"},
		{Text: "info"},
		{Text: "led"},
		{Text: "motion"},
		{Text: "exit"},
		{Text: "quit"},
	}

	for group, groupData := range command.Data {
		if strings.HasPrefix(sub, string(group)+" ") {
			s, sub = complete(sub[len(string(group))+1:], groupData)
			break
		}
	}

	return prompt.FilterHasPrefix(s, sub, true)
}

func complete(sub string, groupData map[command.Type][]command.Command) ([]prompt.Suggest, string) {
	suggestions := make([]prompt.Suggest, 0)
	sub = strings.TrimLeft(sub, " ")

	for commandType, commands := range groupData {
		if !strings.HasPrefix(sub, string(commandType)+" ") {
			suggestions = append(suggestions, prompt.Suggest{Text: string(commandType)})
			continue
		}

		suggestions = make([]prompt.Suggest, 0)
		for _, _command := range commands {
			suggestions = append(suggestions, prompt.Suggest{Text: string(_command)})
		}

		return suggestions, sub[len(commandType)+1:]
	}

	return suggestions, sub
}
