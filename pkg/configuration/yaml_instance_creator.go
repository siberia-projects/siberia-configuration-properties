// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

const (
	supportedYamlFormat = "yaml"
	supportedYmlFormat  = "yml"
)

// YamlInstanceCreator is an instance of InstanceCreator
// which creates a YamlConfiguration
type YamlInstanceCreator struct {
}

// NewYamlInstanceCreator creates a new YamlInstanceCreator
func NewYamlInstanceCreator() InstanceCreator {
	return &YamlInstanceCreator{}
}

func (creator *YamlInstanceCreator) Supports(format string) bool {
	return format == supportedYamlFormat || format == supportedYmlFormat
}

func (creator *YamlInstanceCreator) Create(content []byte) (Configuration, error) {
	return NewYamlConfiguration(content), nil
}
