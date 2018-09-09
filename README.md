# ConfigManager [![Build Status](https://travis-ci.com/cloud-spin/configmanager.svg?branch=master)](https://travis-ci.com/cloud-spin/configmanager) [![codecov](https://codecov.io/gh/cloud-spin/configmanager/branch/master/graph/badge.svg)](https://codecov.io/gh/cloud-spin/configmanager) [![Go Report Card](https://goreportcard.com/badge/github.com/cloud-spin/configmanager)](https://goreportcard.com/report/github.com/cloud-spin/configmanager)

ConfigManager manages service startup configurations provided as external JSON files. ConfigManager is a very simple library that provides an abstraction around loading and parsing external JSON. The library has 100% code coverage, freeing your service from the need to have extra tests to achieve 100% code coverage.

#### How to Use

```go
package main

import (
  "github.com/cloud-spin/configmanager"
)

type testConfigs struct {
	Config1 string
	Config2 int
}

func main() {
	configManager := NewConfigManager()
	configs := &testConfigs{}
	configManager.LoadConfigs("service_configs.json", configs)
	fmt.Println(configs.Config1)
	fmt.Println(configs.Config2)
}
```

Also refer to the tests at [configmanager_test.go](configmanager_test.go).
