// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"os"
)

var _ = Describe(
	"Verifying an expand env file reader's functionality",
	Label("expand_env_file_reader"),
	func() {
		instanceCreator := &txtInstanceCreator{}
		instanceCreators := []InstanceCreator{instanceCreator}

		When("an initialization is going", func() {
			Context("and instance creators were not provided", func() {
				It("should return an error", func() {
					// when
					reader, err := NewExpandEnvFileReader([]InstanceCreator{})

					// then
					Expect(err).ToNot(BeNil())
					Expect(reader).To(BeNil())
				})
			})

			Context("and instance creators were provided", func() {
				It("should return a new instance of reader", func() {
					// when
					reader, err := NewExpandEnvFileReader(instanceCreators)

					// then
					Expect(err).To(BeNil())
					Expect(reader).ToNot(BeNil())
				})
			})
		})

		When("a reading of a file is going", func() {
			reader, _ := NewExpandEnvFileReader(instanceCreators)

			Context("and empty file extension was provided", func() {
				It("should return an error", func() {
					// when
					config, err := reader.Read("unknown_file_format")

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())
				})
			})

			Context("and no instance creator was found for the file with such extension", func() {
				It("should return an error", func() {
					// when
					config, err := reader.Read("unknown_file.format")

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())
				})
			})

			Context("and was not possible to open the file", func() {
				It("should return an error", func() {
					// when
					config, err := reader.Read("unknown_file.txt")

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())
				})
			})

			Context("and an env variable was declared but not existed", func() {
				It("should return an error", func() {
					// given
					tmpFileContent := "${UNKNOWN_ENV_VARIABLE}"
					tmpFilepath := "empty_file.txt"

					// when / then
					err := os.WriteFile(tmpFilepath, []byte(tmpFileContent), 777)
					Expect(err).To(BeNil())

					config, err := reader.Read(tmpFilepath)

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})

			Context("and was not possible to create an instance of read content", func() {
				It("should return an error", func() {
					// given
					tmpFileContent := "${UNKNOWN_ENV_VARIABLE:with_default_value}"
					tmpFilepath := "empty_file.txt"

					// when / then
					instanceCreator.On("Create", mock.Anything).
						Times(1).
						Return(nil, fmt.Errorf("unknown error"))

					err := os.WriteFile(tmpFilepath, []byte(tmpFileContent), 777)
					Expect(err).To(BeNil())

					config, err := reader.Read(tmpFilepath)

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})

			Context("and no error was occurred", func() {
				It("should return a new instance of an appropriate configuration", func() {
					// given
					tmpFileContent := "${UNKNOWN_ENV_VARIABLE:with_default_value}"
					tmpFilepath := "empty_file.txt"

					// when / then
					instanceCreator.On("Create", mock.Anything).
						Times(1).
						Return(&simpleConfiguration{}, nil)

					err := os.WriteFile(tmpFilepath, []byte(tmpFileContent), 777)
					Expect(err).To(BeNil())

					config, err := reader.Read(tmpFilepath)

					// then
					Expect(err).To(BeNil())
					Expect(config).ToNot(BeNil())

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})
		})
	},
)
