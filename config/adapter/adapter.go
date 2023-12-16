/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package adapter

type Adapter interface {
	Peek(s string) bool
	Unmarshal(data []byte, v any) error
}
