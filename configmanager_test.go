// Copyright (c) 2018 cloud-spin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package configmanager

import (
	"testing"
)

type testConfigs struct {
	Config1 string
	Config2 int
}

func TestLoadConfigsWithValidFileShouldLoadConfigs(t *testing.T) {
	configManager := NewConfigManager()
	configs := &testConfigs{}
	configManager.LoadConfigs("testdata/config_manager_valid_configs.json", configs)
	if configs.Config1 != "value1" {
		t.Errorf("Expected: '%s'; Got: '%s'", "value1", configs.Config1)
	}
	if configs.Config2 != 2 {
		t.Errorf("Expected: '%d'; Got: '%d'", 2, configs.Config2)
	}
}

func TestLoadConfigsWithInvalidFileShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic as the configs file doesn't exist")
		}
	}()

	configManager := NewConfigManager()
	configs := &testConfigs{}
	configManager.LoadConfigs("testdata/config_manager_invalid_configs.json", configs)
}

func TestLoadConfigsWithNonExistentFileShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic as the configs file doesn't exist")
		}
	}()

	configManager := NewConfigManager()
	configs := &testConfigs{}
	configManager.LoadConfigs("testdata/config_manager_nonexistent_configs.json", configs)
}
