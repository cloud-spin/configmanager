# ConfigManager [![Build Status](https://travis-ci.com/cloud-spin/configmanager.svg?branch=master)](https://travis-ci.com/cloud-spin/configmanager) [![codecov](https://codecov.io/gh/cloud-spin/configmanager/branch/master/graph/badge.svg)](https://codecov.io/gh/cloud-spin/configmanager) [![Go Report Card](https://goreportcard.com/badge/github.com/cloud-spin/configmanager)](https://goreportcard.com/report/github.com/cloud-spin/configmanager)  [![GoDoc](https://godoc.org/github.com/cloud-spin/configmanager?status.svg)](https://godoc.org/github.com/cloud-spin/configmanager)

ConfigManager manages service startup configurations provided as external JSON files. ConfigManager is a very simple library that provides an abstraction around loading and parsing external JSON. The library has 100% code coverage, freeing your service from the need to have extra tests to achieve 100% code coverage.

#### Install

From a configured [Go environment](https://golang.org/doc/install#testing):
```sh
go get -u github.com/cloud-spin/configmanager
```

If you are using dep:
```sh
dep ensure -add github.com/cloud-spin/configmanager
```

#### How to Use

ConfigManager takes a JSON file path as param.

```go
package main

import (
	"fmt"

	"github.com/cloud-spin/configmanager"
)

type testConfigs struct {
	Config1 string
	Config2 int
}

func main() {
	cm := configmanager.NewConfigManager()
	configs := &testConfigs{}
	cm.LoadConfigs("service_configs.json", configs)
	fmt.Println(configs.Config1)
	fmt.Println(configs.Config2)
}
```

Above "service_configs.json" contains below:

```json
{
    "Config1": "value1",
    "Config2": 2
}

```

Output:
```
value1
2
```

Also refer to the tests at [configmanager_test.go](configmanager_test.go).


## License
MIT, see [LICENSE](LICENSE).

"Use, abuse, have fun and contribute back!"


## Contributions
See [Contributions.md](https://github.com/cloud-spin/docs/blob/master/contributing.md).

