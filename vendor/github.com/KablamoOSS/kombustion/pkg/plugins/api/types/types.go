package types

import (
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// PluginResult result of a parserFunc
type PluginResult struct {
	Data   kombustionTypes.TemplateObject
	Errors []error
}

// Config provides Kombustion with information about your plugin
type Config struct {
	Name    string
	Version string
	Prefix  string
	Help    Help
}

// Help - a set of available documentation fields
type Help struct {
	// The name of the plugin
	Name string
	// A short description of what the plugin does
	Description string

	// Help information for all the types this pplugin provides
	TypeMappings []TypeMapping

	// Examples/Snippets of how this plugin can be used
	Snippets []string
}

// TypeMapping - recursive list of types with its associated config object
type TypeMapping struct {
	Name        string
	Description string
	Config      interface{}
}
