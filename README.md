Siberia Configuration Properties
=================

[![Author](https://img.shields.io/badge/author-@siberia_projects-green.svg)](https://github.com/siberia-projects/siberia-configuration-properties)
[![Source Code](https://img.shields.io/badge/source-siberia/main-blue.svg)](https://github.com/siberia-projects/siberia-configuration-properties)

## What is it?
Siberia-configuration-properties is a library

## How to download?

```console
john@doe-pc:~$ go get github.com/siberia-projects/siberia-configuration-properties
```

## How to use?
Will be described later

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
```
