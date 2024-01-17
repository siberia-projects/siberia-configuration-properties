// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

// Configuration is an interface which describes
// a read configuration
type Configuration interface {
	// WriteTo writes an appropriate data
	// into a Properties
	//
	// Returns an error when something was wrong
	WriteTo(properties Properties) error
}
