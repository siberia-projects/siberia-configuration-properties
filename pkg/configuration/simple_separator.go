// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

// SimpleSeparator a simple implementation of Separator
type SimpleSeparator struct {
}

// NewSimpleSeparator creates a new SimpleSeparator
func NewSimpleSeparator() Separator {
	return &SimpleSeparator{}
}

func (separator *SimpleSeparator) Separate(config Configuration, properties Properties) error {
	return config.WriteTo(properties)
}
