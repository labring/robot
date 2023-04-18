/*
Copyright 2023 cuisongliu@qq.com.

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

package gh

import (
	"fmt"
	"github.com/cuisongliu/logger"
	"github.com/labring-actions/gh-rebot/pkg/config"
	"github.com/labring-actions/gh-rebot/pkg/utils"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/json"
)

type ActionOut struct {
	Conclusion string `json:"conclusion"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	URL        string `json:"url"`
	IsSuccess  bool   `json:"-"`
}

func CheckRelease(tagName string) (*ActionOut, error) {
	workflowOutput, err := utils.RunCommandWithOutput(fmt.Sprintf(gitWorkflowCheck, config.GlobalsConfig.Release.Action, tagName), true)
	if err != nil {
		return nil, err
	}
	var out ActionOut
	if err = json.Unmarshal([]byte(workflowOutput), &out); err != nil {
		return nil, err
	}
	if out.Conclusion == "success" && out.Status == "completed" {
		out.IsSuccess = true
		return &out, nil
	}
	if out.Conclusion == "failure" {
		out.IsSuccess = false
		return &out, nil
	}
	if out.Conclusion == "" {
		logger.Debug("workflow release is in progress, please wait")
		return CheckRelease(tagName)
	}
	return nil, errors.New("workflow Conclusion is error")
}