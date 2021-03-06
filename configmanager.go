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
	"encoding/json"
	"io/ioutil"
)

// ConfigManager manages configurations.
type ConfigManager interface {
	LoadConfigs(file string, configs interface{}) interface{}
}

// ConfigManagerImpl manages configurations.
type ConfigManagerImpl struct {
}

// New creates a new instance of ConfigManager.
func New() ConfigManager {
	return &ConfigManagerImpl{}
}

// LoadConfigs loads the configs from the specified file.
func (m *ConfigManagerImpl) LoadConfigs(file string, configs interface{}) interface{} {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(content, configs); err != nil {
		panic(err)
	}
	return configs
}
