package hellostr

import "strings"

const (
	frenchPrefix  string = "Bonjour, "
	spanishPrefix        = "Hola, "
	englishPrefix        = "Hello, "
)

// Hello - Says hello to the world/person('name') in ('language') French, Spanish or English.
// Default language is "English" if unknown language is passed
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return languagePrefix(strings.ToUpper(language)) + name
}

// languagePrefix - Selects a prefix based on the language selected and returns it.
// If unknown/no language is given it returns the English prefix.
func languagePrefix(language string) string {
	switch language {
	case "FRENCH":
		return frenchPrefix
	case "SPANISH":
		return spanishPrefix
	case "ENGLISH":
		return englishPrefix
	default:
		return englishPrefix
	}
}

// Bitcoin stuct was created just for the sake of dcoumentation testing.
type Bitcoin struct {
	balance float64
}
