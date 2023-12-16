/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package strconv

import "errors"

var ErrSyntax = errors.New("khovrakh/strconv: invalid syntax")

func ParseSize(s string) (int, error) {
	var i, size, multiplier int

	for len(s) > i {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			size = (size * 10) + digit
		} else {
			break
		}

		i++
	}

	switch s[i:] {
	case "tb":
		multiplier = 1 << 40
	case "gb":
		multiplier = 1 << 30
	case "mb":
		multiplier = 1 << 20
	case "kb":
		multiplier = 1 << 10
	case "b":
		multiplier = 1
	default:
		return 0, ErrSyntax
	}

	return size * multiplier, nil
}
