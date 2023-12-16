/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package hook

import (
	"reflect"

	"github.com/mitchellh/mapstructure"

	"github.com/dirakkar/khovrakh/strconv"
)

func StringToSize() mapstructure.DecodeHookFunc {
	return func(f, t reflect.Type, data any) (any, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		size, err := strconv.ParseSize(data.(string))
		if err != nil {
			return data, nil
		}

		return size, nil
	}
}
