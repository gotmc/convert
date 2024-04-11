// Copyright (c) 2023–2024 The convert developers. All rights reserved.
// Project site: https://github.com/gotmc/convert
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package convert

import (
	"testing"
)

func TestStringFloats(t *testing.T) {
	testCases := []struct {
		given    string
		sep      string
		expected []float64
	}{
		{
			given:    "0.001,0.002,0.003",
			sep:      ",",
			expected: []float64{0.001, 0.002, 0.003},
		},
		{
			given:    "10.0,20.0,30.0,40.0",
			sep:      ",",
			expected: []float64{10.0, 20.0, 30.0, 40.0},
		},
	}
	for _, tc := range testCases {
		calcs, err := StringToFloats(tc.given, tc.sep)
		if err != nil {
			t.Errorf("error parsing slice of strings into floats: %s", err)
		}
		for i, calc := range calcs {
			if calc != tc.expected[i] {
				t.Errorf(
					"given %s / index %d expected = %f / calculated = %f",
					tc.given,
					i,
					tc.expected[i],
					calc,
				)
			}
		}
	}
}

func TestStringNFloats(t *testing.T) {
	testCases := []struct {
		given       string
		sep         string
		numExpected int
		expected    []float64
	}{
		{
			given:       "0.001,0.002,0.003",
			sep:         ",",
			numExpected: 3,
			expected:    []float64{0.001, 0.002, 0.003},
		},
		{
			given:       "10.0,20.0,30.0,40.0",
			sep:         ",",
			numExpected: 4,
			expected:    []float64{10.0, 20.0, 30.0, 40.0},
		},
	}
	for _, tc := range testCases {
		calcs, err := StringToNFloats(tc.given, tc.sep, tc.numExpected)
		if err != nil {
			t.Errorf("error parsing slice of strings into floats: %s", err)
		}
		for i, calc := range calcs {
			if calc != tc.expected[i] {
				t.Errorf(
					"given %s / index %d expected = %f / calculated = %f",
					tc.given,
					i,
					tc.expected[i],
					calc,
				)
			}
		}
	}
}

func TestStripDoubleQuotes(t *testing.T) {
	testCases := []struct {
		given    string
		expected string
	}{
		{
			given:    "\"0.001,0.002\"",
			expected: "0.001,0.002",
		},
		{
			given:    "\"0.001,0.002\"\n",
			expected: "0.001,0.002",
		},
		{
			given:    "\"foo,bash\"\n",
			expected: "foo,bash",
		},
	}
	for _, tc := range testCases {
		got := StripDoubleQuotes(tc.given)
		if got != tc.expected {
			t.Errorf(
				"given %s / expected = %s / got = %s",
				tc.given,
				tc.expected,
				got,
			)
		}
	}
}
