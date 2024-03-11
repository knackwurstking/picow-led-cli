package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest

	switch text := strings.TrimLeft(d.Text, " "); {
	case strings.HasPrefix(text, "config "):
		s = completerSuggestConfig(d.Text[7:])
	case strings.HasPrefix(text, "info "):
		s = completerSuggestInfo(d.Text[5:])
	case strings.HasPrefix(text, "led "):
		s = completerSuggestLED(d.Text[4:])
	case strings.HasPrefix(text, "motion "):
		s = completerSuggestMotion(d.Text[7:])
	default:
		s = []prompt.Suggest{
			{Text: "config"},
			{Text: "info"},
			{Text: "led"},
			{Text: "motion"},
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func completerSuggestConfig(text string) []prompt.Suggest {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		s = []prompt.Suggest{
			{Text: "led"},
			{Text: "motion"},
			{Text: "motion-timeout"},
			{Text: "pwm-range"},
		}
	case strings.HasPrefix(t, "get "):
		s = []prompt.Suggest{
			{Text: "led"},
			{Text: "motion"},
			{Text: "motion-timeout"},
			{Text: "pwm-range"},
		}
	default:
		s = []prompt.Suggest{
			{Text: "set"},
			{Text: "get"},
		}
	}

	return s
}

func completerSuggestInfo(text string) []prompt.Suggest {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		s = []prompt.Suggest{}
	default:
		s = []prompt.Suggest{}
	}

	return s
}

func completerSuggestLED(text string) []prompt.Suggest {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		s = []prompt.Suggest{}
	default:
		s = []prompt.Suggest{}
	}

	return s
}

func completerSuggestMotion(text string) []prompt.Suggest {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		s = []prompt.Suggest{}
	default:
		s = []prompt.Suggest{}
	}

	return s
}
