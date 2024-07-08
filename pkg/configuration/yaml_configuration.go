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
	"github.com/goccy/go-yaml/ast"
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

func (configuration *YamlConfiguration) ReadString(prefix string) (string, error) {
	path := configuration.getPrefixPath(prefix)
	contentReader := bytes.NewReader(configuration.content)

	node, err := path.ReadNode(contentReader)
	if err != nil {
		return "", err
	}

	if node.Type() == ast.NullType {
		return "", nil
	}

	value := node.String()

	return strings.Trim(value, " "), nil
}

// getPropertiesPath gets a path to Properties starting point
//
// e.g.
//
// prefix = my.custom.properties
// path = $.my.custom.properties
func (configuration *YamlConfiguration) getPropertiesPath(properties Properties) *yaml.Path {
	prefix := properties.GetPrefix()
	return configuration.getPrefixPath(prefix)
}

// getPrefixPath gets a path to a property by prefix
//
// e.g.
//
// prefix = my.custom.properties
// path = $.my.custom.properties
func (configuration *YamlConfiguration) getPrefixPath(prefix string) *yaml.Path {
	prefixParts := strings.Split(prefix, propertiesLevelsPrefixSeparator)

	pathBuilder := &yaml.PathBuilder{}
	pathBuilder = pathBuilder.Root()

	for _, prefixPart := range prefixParts {
		pathBuilder = pathBuilder.Child(prefixPart)
	}

	return pathBuilder.Build()
}
