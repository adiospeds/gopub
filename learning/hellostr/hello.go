package hellostr

import "strings"

const frenchPrefix = "Bonjour, "
const spanishPrefix = "Hola, "
const englishPrefix = "Hello, "

// Hello - Says hello to the world/person in French, Spanish or English.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return languagePrefix(strings.ToUpper(language)) + name
}

// languagePrefix - Selects a prefix based on the language selected and returns it.
// If unknown / no language is gives it returns the English prefix.
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
