// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Verifying a simple separator's functionality",
	Label("simple_separator"),
	func() {
		separator := NewSimpleSeparator()

		When("a separation is going", func() {
			Context("and properties could be separated from a configuration", func() {
				It("should successfully fill an empty properties instance", func() {
					// given
					properties := &simpleProperties{}
					configurationInstance := &simpleConfiguration{}

					// when / then
					configurationInstance.On("WriteTo", properties).
						Times(1).
						Return(nil)

					err := separator.Separate(configurationInstance, properties)
					Expect(err).To(BeNil())
				})
			})

			Context("and properties could not be separated from a configuration", func() {
				It("should not fill an empty properties and return an error", func() {
					// given
					properties := &simpleProperties{}
					configurationInstance := &simpleConfiguration{}

					// when / then
					configurationInstance.On("WriteTo", properties).
						Times(1).
						Return(fmt.Errorf("unknown error"))

					err := separator.Separate(configurationInstance, properties)
					Expect(err).ToNot(BeNil())
				})
			})
		})
	},
)
