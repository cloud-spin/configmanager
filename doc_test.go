package configmanager_test

import (
	"fmt"

	"github.com/cloud-spin/configmanager"
)

type Configs struct {
	Config1 string
}

func Example() {
	configs := &Configs{}
	cm := configmanager.New()
	cm.LoadConfigs("testdata/config_manager_valid_configs.json", configs)
	fmt.Println(configs.Config1)
	// Output: value1
}
