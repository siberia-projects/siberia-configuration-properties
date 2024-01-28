// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe(
	"Verifying a configuration's reading functionality",
	Label("configuration_reading"),
	func() {
		When("a reading of yaml configuration file is going", func() {
			Context("and no error has occurred", func() {
				It("should successfully return a new configuration instance in an appropriate format", func() {
					// given
					name := "Alex"
					tmpFileContentPattern := `simple:
  prefix:
    name: ${SIMPLE_CONFIG_NAME:%s}`

					tmpFileContent := fmt.Sprintf(tmpFileContentPattern, name)
					tmpFilepath := "config.yaml"

					// when / then
					err := os.WriteFile(tmpFilepath, []byte(tmpFileContent), 777)
					Expect(err).To(BeNil())

					config, err := ReadConfiguration(tmpFilepath)

					// then
					Expect(err).To(BeNil())
					Expect(config).ToNot(BeNil())

					_, ok := config.(*YamlConfiguration)
					Expect(ok).To(BeTrue())

					properties := &simpleProperties{}
					err = config.WriteTo(properties)

					Expect(err).To(BeNil())
					Expect(properties.Name).To(Equal(name))

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})

			Context("and an the file has an incorrect format", func() {
				It("should return an error", func() {
					// given
					tmpFilepath := "config.unknown_format"

					// when / then
					err := os.WriteFile(tmpFilepath, []byte(""), 777)
					Expect(err).To(BeNil())

					config, err := ReadConfiguration(tmpFilepath)

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})

			Context("and an environment variable is not presented", func() {
				It("should return an error", func() {
					// given
					tmpFileContent := `simple:
  prefix:
    name: ${SIMPLE_CONFIG_NAME}`

					tmpFilepath := "config.yaml"

					// when / then
					err := os.WriteFile(tmpFilepath, []byte(tmpFileContent), 777)
					Expect(err).To(BeNil())

					config, err := ReadConfiguration(tmpFilepath)

					// then
					Expect(err).ToNot(BeNil())
					Expect(config).To(BeNil())

					// cleanup
					err = os.Remove(tmpFilepath)
					Expect(err).To(BeNil())
				})
			})
		})
	},
)
