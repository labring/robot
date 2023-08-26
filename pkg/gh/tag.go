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
	"github.com/labring/robot/pkg/utils"
	"github.com/pkg/errors"
)

func Tag(tag, branch string) error {
	if err := setPreGithub(); err != nil {
		return err
	}
	var firstShell any
	if branch == "" {
		firstShell = utils.SkipShell("")
	} else {
		firstShell = fmt.Sprintf(gitCheck, branch)
	}
	if ok, err := checkRemoteTagExists(tag); err != nil {
		return err
	} else {
		if !ok {
			shells := []any{
				firstShell,
				fmt.Sprintf(gitNewTag, tag),
				fmt.Sprintf(gitPushRemote, tag),
			}
			if err = utils.ExecShellForAny()(shells); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("tag already exists")
}
