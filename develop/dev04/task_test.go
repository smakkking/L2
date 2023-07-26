package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_data struct {
	name string
	arg  []string
	want map[string][]string
}

func TestUnpackString(t *testing.T) {
	testCases := []test_data{
		{
			name: "normal",
			arg:  []string{`Пятка`, `тяпка`, `лом`, `мол`},
			want: map[string][]string{
				`Пятка`: []string{`пятка`, `тяпка`},
				`лом`:   []string{`лом`, `мол`},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := GroupAnagrams(tc.arg)
			assert.Equal(t, tc.want, res)
		})
	}
}
