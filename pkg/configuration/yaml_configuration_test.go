// Copyright (c) 2024 Nikolai Osipov <nao99.dev@gmail.com>
//
// All rights are reserved
//
// Information about a license can be found in a LICENSE file
// in a root of the project

package configuration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Verifying a yaml configuration's functionality",
	Label("yaml_configuration"),
	func() {
		When("a writing into a custom structure is going", func() {
			Context("and a structure has an incorrect prefix", func() {
				It("should not return any error but the structure should be non-touched", func() {
					// given
					content := "simple:\n  prefix:\n    data:\n      alex"
					contentBytes := []byte(content)

					properties := &simpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).To(BeNil())
					Expect(properties.Name).To(BeEmpty())
				})
			})

			Context("and a content by passed prefix is not applicable to a structure", func() {
				It("should return an error", func() {
					// given
					content := "notSimple:\n  prefix:\n    data:\n      alex"
					contentBytes := []byte(content)

					properties := &notSimpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).ToNot(BeNil())
				})
			})

			Context("and a content has an incorrect format", func() {
				It("should return an error", func() {
					// given
					content := "{\"notSimple\": {\"prefix\": {\"data\": \"Alex\"}}}"
					contentBytes := []byte(content)

					properties := &notSimpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).ToNot(BeNil())
				})
			})

			Context("and a structure has an empty prefix", func() {
				It("should successfully fill the structure with all configuration's content", func() {
					// given
					content := "empty:\n  prefix:\n    data:\n      key: value"
					contentBytes := []byte(content)

					properties := &emptyPrefixProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).To(BeNil())
					Expect(properties.Empty.Prefix.Data).ToNot(BeEmpty())
				})
			})

			Context("and a structure has less fields than a content", func() {
				It("should successfully fill the structure with a content's part", func() {
					// given
					content := "simple:\n  prefix:\n    additionalField: value\n    name:\n      alex"
					contentBytes := []byte(content)

					properties := &simpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).To(BeNil())
					Expect(properties.Name).ToNot(BeEmpty())
				})
			})

			Context("and a structure has more fields than a content", func() {
				It("should successfully fill the structure with a content", func() {
					// given
					content := "notSimple:\n  prefix:\n    data:\n      key: value"
					contentBytes := []byte(content)

					properties := &notSimpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).To(BeNil())

					Expect(properties.Data).ToNot(BeEmpty())
					Expect(properties.AdditionalField).To(BeEmpty())
				})
			})

			Context("and a structure has all fields", func() {
				It("should successfully fill the structure with a content", func() {
					// given
					content := "notSimple:\n  prefix:\n    additionalField: value\n    data:\n      key: value"
					contentBytes := []byte(content)

					properties := &notSimpleProperties{}
					config := NewYamlConfiguration(contentBytes)

					// when
					err := config.WriteTo(properties)

					// then
					Expect(err).To(BeNil())

					Expect(properties.Data).ToNot(BeEmpty())
					Expect(properties.AdditionalField).ToNot(BeEmpty())
				})
			})
		})
	},
)
