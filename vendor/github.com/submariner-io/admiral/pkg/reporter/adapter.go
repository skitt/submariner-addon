/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package reporter

import (
	"unicode"

	"github.com/pkg/errors"
)

type Adapter struct {
	Basic
}

func (a *Adapter) Error(err error, message string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	err = errors.Wrapf(err, message, args...)

	capitalizeFirst := func(str string) string {
		for i, v := range str {
			return string(unicode.ToUpper(v)) + str[i+1:]
		}

		return ""
	}

	a.Basic.Failure(capitalizeFirst(err.Error()))
	a.Basic.End()

	return err
}
