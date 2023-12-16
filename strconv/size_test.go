/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package strconv_test

import (
	"reflect"
	"testing"

	"github.com/dirakkar/khovrakh/strconv"
)

var ParseSizeTests = map[string]struct {
	input string
	want  int
	err   error
}{
	"b":  {input: "4b", want: 4},
	"kb": {input: "4kb", want: 4 << 10},
	"mb": {input: "4mb", want: 4 << 20},
	"gb": {input: "4gb", want: 4 << 30},
	"tb": {input: "4tb", want: 4 << 40},

	"empty":   {input: "", err: strconv.ErrSyntax},
	"digit":   {input: "123", err: strconv.ErrSyntax},
	"letter":  {input: "abc", err: strconv.ErrSyntax},
	"invalid": {input: "1kb3", err: strconv.ErrSyntax},
}

func TestParseSize(t *testing.T) {
	for name, test := range ParseSizeTests {
		t.Run(name, func(t *testing.T) {
			got, err := strconv.ParseSize(test.input)
			if err != test.err {
				t.Fatalf("error '%s', does not match '%s'", err, test.err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("result '%v' does not match '%v'", got, test.want)
			}
		})
	}
}
