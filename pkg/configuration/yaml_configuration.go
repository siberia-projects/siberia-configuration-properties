// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

import (
	"bytes"
	"github.com/goccy/go-yaml"
	"strings"
)

// YamlConfiguration is an instance of Configuration
// which represents YAML configuration
type YamlConfiguration struct {
	content []byte
}

// NewYamlConfiguration creates a new YamlConfiguration
func NewYamlConfiguration(content []byte) Configuration {
	return &YamlConfiguration{content: content}
}

func (configuration *YamlConfiguration) WriteTo(properties Properties) error {
	prefix := properties.GetPrefix()
	if prefix == "" {
		return yaml.Unmarshal(configuration.content, properties)
	}

	path := configuration.getPropertiesPath(properties)
	contentReader := bytes.NewReader(configuration.content)

	return path.Read(contentReader, properties)
}

// getPropertiesPath gets a path to Properties starting point
//
// e.g.
//
// prefix = my.custom.properties
// path = $.my.custom.properties
func (configuration *YamlConfiguration) getPropertiesPath(properties Properties) *yaml.Path {
	prefix := properties.GetPrefix()
	prefixParts := strings.Split(prefix, propertiesLevelsPrefixSeparator)

	pathBuilder := &yaml.PathBuilder{}
	pathBuilder = pathBuilder.Root()

	for _, prefixPart := range prefixParts {
		pathBuilder = pathBuilder.Child(prefixPart)
	}

	return pathBuilder.Build()
}
