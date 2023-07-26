package dev02_test

import (
	"testing"

	"github.com/smakkking/wildberries/L2/develop/dev02"
	"github.com/stretchr/testify/assert"
)

type test_data struct {
	name string
	arg  []rune
	want []rune
	err  bool
}

func TestUnpackString(t *testing.T) {
	testCases := []test_data{
		{
			name: "normal unpack",
			arg:  []rune("a2b2"),
			want: []rune("aabb"),
			err:  false,
		},
		{
			name: "normal unpack with one symbol",
			arg:  []rune("a2b2c"),
			want: []rune("aabbc"),
			err:  false,
		},
		{
			name: "normal unpack only one symbol",
			arg:  []rune("abc"),
			want: []rune("abc"),
			err:  false,
		},
		{
			name: "unpack with ecran",
			arg:  []rune("\\45"),
			want: []rune("44444"),
			err:  false,
		},
		{
			name: "unpack with ecran",
			arg:  []rune("qwe\\4\\5"),
			want: []rune("qwe45"),
			err:  false,
		},
		{
			name: "unpack with ecran",
			arg:  []rune("qwe\\\\2"),
			want: []rune("qwe\\\\"),
			err:  false,
		},
		{
			name: "unpack with ecran",
			arg:  []rune("45"),
			want: []rune(""),
			err:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := dev02.UnpackString(tc.arg)
			if err != nil {
				assert.Equal(t, tc.err, true)
			} else {
				assert.Equal(t, tc.want, res)
			}
		})
	}
}
