// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

// Reader is an interface which provides
// a functionality for reading a Configuration from a file
type Reader interface {
	// Read reads a Configuration from a file using
	// passed filepath
	//
	// Returns an error when something was wrong
	Read(filepath string) (Configuration, error)
}
