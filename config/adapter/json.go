/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package adapter

import "encoding/json"

type JSON struct{}

func (JSON) Peek(s string) bool { return s == ".json" }

func (JSON) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
