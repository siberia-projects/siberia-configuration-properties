// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

// Separator is an interface which provides
// a functionality to put a part of Configuration
// into passed Properties
type Separator interface {
	// Separate puts a part of Configuration
	// into passed Properties
	//
	// Returns an error when something was wrong
	Separate(config Configuration, properties Properties) error
}
