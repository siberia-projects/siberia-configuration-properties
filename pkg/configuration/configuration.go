// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

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

// ReadConfiguration is a default method to read a Configuration
// by passed path
//
// It reads, replaces env variables and returns parsed content
// as a new Configuration
//
// Returns error when was not possible to open / read a file
//
// This method exists only for simplification of Configuration's
// reading
//
// If you want to have a better control over a configuration,
// just use everything you need manually
func ReadConfiguration(configFilepath string) (Configuration, error) {
	yamlInstanceCreator := NewYamlInstanceCreator()
	instanceCreators := []InstanceCreator{yamlInstanceCreator}

	configurationReader, _ := NewExpandEnvFileReader(instanceCreators)

	configurationInstance, err := configurationReader.Read(configFilepath)
	if err != nil {
		return nil, err
	}

	return configurationInstance, nil
}
