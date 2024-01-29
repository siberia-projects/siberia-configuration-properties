// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

// InstanceCreator is an interface which provides
// a functionality to create an instance of a Configuration
type InstanceCreator interface {
	// Supports checks if an InstanceCreator supports
	// a content of passed format
	Supports(format string) bool

	// Create creates an instance of a Configuration
	//
	// Returns an error when something was wrong
	Create(content []byte) (Configuration, error)
}
