package main

import "testing"

func TestParseGoVersion(t *testing.T) {
	var tests = []struct {
		s string
		v GoVersion
	}{
		{
			s: "go1.5.7",
			v: GoVersion{Major: 1, Minor: 5, Patch: 7},
		},
		{
			s: "go71.23.44",
			v: GoVersion{Major: 71, Minor: 23, Patch: 44},
		},
		{
			s: "1.2.3",
		},
		{
			s: "go1.2",
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			v := ParseGoVersion(test.s)
			if v != test.v {
				t.Fatalf("wrong version, wanted %v, got %v", test.v, v)
			}
		})
	}
}

func TestGoVersionAtLeast(t *testing.T) {
	var tests = []struct {
		v1, v2 GoVersion
		res    bool
	}{
		{
			v1:  GoVersion{Major: 1, Minor: 9, Patch: 2},
			v2:  GoVersion{Major: 1, Minor: 0, Patch: 0},
			res: true,
		},
		{
			v1:  GoVersion{Major: 1, Minor: 5, Patch: 7},
			v2:  GoVersion{Major: 71, Minor: 23, Patch: 44},
			res: false,
		},
		{
			v1:  GoVersion{},
			v2:  GoVersion{Major: 1, Minor: 7, Patch: 2},
			res: true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := test.v1.AtLeast(test.v2)
			if res != test.res {
				t.Fatalf("wrong result, want %v, got %v", test.res, res)
			}
		})
	}
}
