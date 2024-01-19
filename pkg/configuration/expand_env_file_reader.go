// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

import (
	"fmt"
	"github.com/siberia-projects/siberia-env/pkg/env"
	"io"
	"os"
	"strings"
)

const (
	filePathAndExtensionSeparator = "."
)

// ExpandEnvFileReader is an implementation of Reader
// which reads a Configuration from an os.File and replaces
// all "${ENV_VARIABLE_NAME}" mentions with real variable values
type ExpandEnvFileReader struct {
	instanceCreators []InstanceCreator
}

// NewExpandEnvFileReader creates a new ExpandEnvFileReader
func NewExpandEnvFileReader(instanceCreators []InstanceCreator) (Reader, error) {
	if len(instanceCreators) < 1 {
		return nil, fmt.Errorf("unable to create a reader without instance creators")
	}

	return &ExpandEnvFileReader{instanceCreators: instanceCreators}, nil
}

func (reader *ExpandEnvFileReader) Read(filepath string) (Configuration, error) {
	fileFormat, err := reader.getFileFormat(filepath)
	if err != nil {
		return nil, err
	}

	instanceCreator, err := reader.getInstanceCreator(fileFormat)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open a configuration file: %s", err.Error())
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read a configuration file: %s", err.Error())
	}

	contentWithExpandedEnv, err := env.ExpandEnvIn(content)
	if err != nil {
		return nil, fmt.Errorf("unable to expand a configuration: %s", err.Error())
	}

	return instanceCreator.Create(contentWithExpandedEnv)
}

func (reader *ExpandEnvFileReader) getFileFormat(filepath string) (string, error) {
	filepathParts := strings.Split(filepath, filePathAndExtensionSeparator)
	if len(filepathParts) < 2 {
		return "", fmt.Errorf("files without extensions are not supported")
	}

	return filepathParts[1], nil
}

func (reader *ExpandEnvFileReader) getInstanceCreator(fileFormat string) (InstanceCreator, error) {
	for _, instanceCreator := range reader.instanceCreators {
		if instanceCreator.Supports(fileFormat) {
			return instanceCreator, nil
		}
	}

	return nil, fmt.Errorf("unknown \"%s\" format", fileFormat)
}
