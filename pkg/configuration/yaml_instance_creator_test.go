// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
// Information about license can be found in the LICENSE file

package configuration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Verifying a yaml instance creator's functionality",
	Label("yaml_configuration_instance_creator"),
	func() {
		yamlInstanceCreator := NewYamlInstanceCreator()

		When("a checking on supporting is going", func() {
			Context("and correct format was passed", func() {
				DescribeTable(
					"should return true",
					func(format string) {
						// when
						supports := yamlInstanceCreator.Supports(format)

						// then
						Expect(supports).To(BeTrue())
					},
					Entry("yaml format", supportedYamlFormat),
					Entry("yml format", supportedYmlFormat),
				)
			})

			Context("and incorrect format was passed", func() {
				DescribeTable(
					"should return false",
					func(format string) {
						// when
						supports := yamlInstanceCreator.Supports(format)

						// then
						Expect(supports).To(BeFalse())
					},
					Entry("json format", "json"),
					Entry("txt format", "txt"),
					Entry("empty format", ""),
					Entry("yml_yaml format", "yml_yaml"),
				)
			})
		})

		When("a creation is going", func() {
			Context("and content is passed", func() {
				It("should return a new instance of yaml configuration", func() {
					// given
					content := []byte{1, 2, 3, 4, 5}

					// when
					configurationInstance, err := yamlInstanceCreator.Create(content)

					// then
					Expect(err).To(BeNil())
					Expect(configurationInstance).ToNot(BeNil())

					yamlConfigurationInstance, casted := configurationInstance.(*YamlConfiguration)

					Expect(casted).To(BeTrue())
					Expect(yamlConfigurationInstance.content).To(Equal(content))
				})
			})
		})
	},
)
