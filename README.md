Siberia Configuration Properties
=================

[![Author](https://img.shields.io/badge/author-@siberia_projects-green.svg)](https://github.com/siberia-projects/siberia-configuration-properties)
[![Source Code](https://img.shields.io/badge/source-siberia/main-blue.svg)](https://github.com/siberia-projects/siberia-configuration-properties)
![Version](https://img.shields.io/badge/version-v1.0.2-green.svg)
[![Coverage Status](https://coveralls.io/repos/github/siberia-projects/siberia-configuration-properties/badge.svg?branch=main)](https://coveralls.io/github/siberia-projects/siberia-configuration-properties?branch=main)

## What is it?
Siberia-configuration-properties is a library which provides a convenient way of
parsing a configuration file with replacing environment variables with real and separation
a parsed data by small configs

## Why?
When I parse a configuration file I want to see 2 things:
 - I can simply type a declaration of an environment variable in the file 
and I will be sure this variable will be parsed automatically
 - I can use only a part of a configuration

When I dived into the internet for looking something similar, I found everything
but not what I wanted to see: Or it was 1 structure for all configuration, or environment
variables were not parsed, or I was forced to provide additional structures which were useless,
or the stars didn't align well etc

Based on the mentioned above this library has the right to life

## How to download?

```console
john@doe-pc:~$ go get github.com/siberia-projects/siberia-configuration-properties
```

## How to use?
Comparing with other libraries it may look a bit complicated, but it's not

All you need is to create a couple of tools, declare your data model and
parse the config into the model using the tools:
 - Declare a structure in a way it implements the **Properties** interface
(it provides a starting point where your data really begins (could be empty - which literally means
"all"))
 - Create an instance of **ExpandEnvFileReader**
 - Read your configuration file using the reader
 - (Optional) Create a **SimpleSeparator**
 - Create an instance of your data model
 - Call **Configuration.WriteTo(Properties)** method or use the **Separator.Separate(Configuration, Properties)**

## Examples
```yaml
my:
  custom:
    properties:
      headers:
        - key: Cache-Control
          value: ${CACHE_CONTROL:no-cache}
        - key: Content-Type
          value: application/json

not:
  my:
    properties:
      supportsHttps: true
```

```go
package main

import (
	"fmt"
	"github.com/siberia-projects/siberia-configuration-properties/pkg/configuration"
)

const (
	configFilepath = "path/to/config.yaml"
)

type CustomProperties struct {
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}

func (properties *CustomProperties) GetPrefix() string {
	return "my.custom.properties"
}

func (properties *CustomProperties) String() string {
	return fmt.Sprintf("%#v", properties)
}

type NotMyProperties struct {
	SupportsHttps bool
}

func (properties *NotMyProperties) GetPrefix() string {
	return "not.my.properties"
}

func (properties *NotMyProperties) String() string {
	return fmt.Sprintf("%#v", properties)
}

func main() {
	yamlInstanceCreator := configuration.NewYamlInstanceCreator()
	instanceCreators := []configuration.InstanceCreator{yamlInstanceCreator}

	configurationReader, err := configuration.NewExpandEnvFileReader(instanceCreators)
	if err != nil {
		panic(err.Error())
	}

	configurationInstance, err := configurationReader.Read(configFilepath)
	if err != nil {
		panic(err.Error())
	}

	propertiesSeparator := configuration.NewSimpleSeparator()
	customProperties := &CustomProperties{}

	err = propertiesSeparator.Separate(configurationInstance, customProperties)
	if err != nil {
		panic(err.Error())
	}

	customPropertiesString := customProperties.String()

	notMyProperties := &NotMyProperties{}

	err = propertiesSeparator.Separate(configurationInstance, notMyProperties)
	if err != nil {
		panic(err.Error())
	}

	notMyPropertiesString := notMyProperties.String()
	
	println(customPropertiesString)
	println(notMyPropertiesString)
}
```

```text
Result:

&main.CustomProperties{Headers:[]main.Header{main.Header{Key:"Cache-Control", Value:"no-cache"}, main.Header{Key:"Content-Type", Value:"application/json"}}}
&main.NotMyProperties{SupportsHttps:false}
```
