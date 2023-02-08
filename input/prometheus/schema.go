/*
Copyright [2018] [jc3wish]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prometheus

import (
	"fmt"
	"time"

	inputDriver "github.com/brokercap/Bifrost/input/driver"
)

func (c *InputPrometheus) GetSchemaList() (data []string, err error) {
	return
}

func (c *InputPrometheus) GetSchemaTableList(schema string) (tableList []inputDriver.TableList, err error) {
	return
}

func (c *InputPrometheus) GetSchemaTableFieldList(schema string, table string) (FieldList []inputDriver.TableFieldInfo, err error) {
	return make([]inputDriver.TableFieldInfo, 0), nil
}

func (c *InputPrometheus) CheckPrivilege() (err error) {
	return
}

func (c *InputPrometheus) CheckUri(CheckPrivilege bool) (CheckUriResult inputDriver.CheckUriResult, err error) {
	var httpCode int
	_, httpCode, err = c.GetPrometheusBody()
	if err != nil {
		c.err = err
		return
	}
	if httpCode < 200 || httpCode >= 300 {
		err = fmt.Errorf("http code:%d", httpCode)
		return
	}
	result := inputDriver.CheckUriResult{
		BinlogFile:     DefaultBinlogFileName,
		BinlogPosition: 0,
		Gtid:           fmt.Sprint(time.Now().Unix()),
		ServerId:       1,
		BinlogFormat:   "row",
		BinlogRowImage: "full",
	}
	return result, nil
}

// 获取队列最新的位点

func (c *InputPrometheus) GetCurrentPosition() (p *inputDriver.PluginPosition, err error) {
	err = fmt.Errorf("not supported")
	return
}

func (c *InputPrometheus) GetVersion() (Version string, err error) {
	return "", nil
}
