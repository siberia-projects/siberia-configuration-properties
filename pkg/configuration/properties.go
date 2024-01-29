// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

const (
	propertiesLevelsPrefixSeparator = "."
)

// Properties is an interface which describes
// a separated part of a Configuration
type Properties interface {
	// GetPrefix returns a prefix in a Configuration
	// which defines a starting point of Properties
	GetPrefix() string
}
