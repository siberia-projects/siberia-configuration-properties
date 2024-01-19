// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

import "github.com/stretchr/testify/mock"

type simpleProperties struct {
	Name string
}

func (configuration *simpleProperties) GetPrefix() string {
	return "simple.prefix"
}

type notSimpleProperties struct {
	Data            map[string]string `yaml:"data"`
	AdditionalField string            `yaml:"additionalField"`
}

func (configuration *notSimpleProperties) GetPrefix() string {
	return "notSimple.prefix"
}

type emptyPrefixProperties struct {
	Empty struct {
		Prefix struct {
			Data map[string]string `yaml:"data"`
		}
	}
}

func (configuration *emptyPrefixProperties) GetPrefix() string {
	return ""
}

type simpleConfiguration struct {
	mock.Mock
}

func (configuration *simpleConfiguration) WriteTo(properties Properties) error {
	arguments := configuration.Called(properties)
	return arguments.Error(0)
}

type txtInstanceCreator struct {
	mock.Mock
}

func (creator *txtInstanceCreator) Supports(format string) bool {
	return format == "txt"
}

func (creator *txtInstanceCreator) Create(content []byte) (Configuration, error) {
	arguments := creator.Called(content)

	config := arguments.Get(0)
	if config != nil {
		return config.(*simpleConfiguration), nil
	}

	return nil, arguments.Error(1)
}
