package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	sub := strings.TrimLeft(d.Text, " ")

	switch text := strings.TrimLeft(d.Text, " "); {
	case strings.HasPrefix(text, "config "):
		s, sub = completerSuggestConfig(sub[7:])
	case strings.HasPrefix(text, "info "):
		s, sub = completerSuggestInfo(sub[5:])
	case strings.HasPrefix(text, "led "):
		s, sub = completerSuggestLED(sub[4:])
	case strings.HasPrefix(text, "motion "):
		s, sub = completerSuggestMotion(sub[7:])
	default:
		s = []prompt.Suggest{
			{Text: "config"},
			{Text: "info"},
			{Text: "led"},
			{Text: "motion"},
			{Text: "exit"},
			{Text: "quit"},
		}
	}

	return prompt.FilterHasPrefix(s, sub, true)
}

func completerSuggestConfig(text string) ([]prompt.Suggest, string) {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		text = t[4:]
		s = []prompt.Suggest{
			{Text: "led"},
			{Text: "motion"},
			{Text: "motion-timeout"},
			{Text: "pwm-range"},
		}
	case strings.HasPrefix(t, "get "):
		text = t[4:]
		s = []prompt.Suggest{
			{Text: "led"},
			{Text: "motion"},
			{Text: "motion-timeout"},
			{Text: "pwm-range"},
		}
	default:
		text = t
		s = []prompt.Suggest{
			{Text: "set"},
			{Text: "get"},
		}
	}

	return s, text
}

func completerSuggestInfo(text string) ([]prompt.Suggest, string) {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		text = t[4:]
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		text = t[4:]
		s = []prompt.Suggest{}
	default:
		text = t
		s = []prompt.Suggest{}
	}

	return s, text
}

func completerSuggestLED(text string) ([]prompt.Suggest, string) {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		text = t[4:]
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		text = t[4:]
		s = []prompt.Suggest{}
	default:
		text = t
		s = []prompt.Suggest{}
	}

	return s, text
}

func completerSuggestMotion(text string) ([]prompt.Suggest, string) {
	var s []prompt.Suggest

	switch t := strings.TrimLeft(text, " "); {
	case strings.HasPrefix(t, "set "):
		text = t[4:]
		s = []prompt.Suggest{}
	case strings.HasPrefix(t, "get "):
		text = t[4:]
		s = []prompt.Suggest{}
	default:
		text = t
		s = []prompt.Suggest{}
	}

	return s, text
}
