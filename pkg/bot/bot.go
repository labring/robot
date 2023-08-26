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

package bot

import (
	"github.com/labring/robot/pkg/types"
	"strings"
)

func GetReleaseComment() string {
	if types.GlobalsBotConfig.GetPrefix() == "/" {
		return "/release"
	}
	return strings.Join([]string{types.GlobalsBotConfig.GetPrefix(), "release"}, types.GlobalsBotConfig.GetSpe())
}

func GetChangelogComment() string {
	if types.GlobalsBotConfig.GetPrefix() == "/" {
		return "/changelog"
	}
	return strings.Join([]string{types.GlobalsBotConfig.GetPrefix(), "changelog"}, types.GlobalsBotConfig.GetSpe())
}

//approve
//lgtm
//hold
//triage
///ok-to-test

//assign cc cc
//cc xx xx xx
//area
//------

//issue
//assign
