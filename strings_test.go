// Copyright (c) 2023 The convert developers. All rights reserved.
// Project site: https://github.com/gotmc/convert
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package convert

import (
	"math"
	"testing"
)

const tolerance = 0.0000000001

func almostEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < tolerance
}

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
				t.Errorf("given %s / index %d expected = %f / calculated = %f", tc.given, i, tc.expected[i], calc)
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
				t.Errorf("given %s / index %d expected = %f / calculated = %f", tc.given, i, tc.expected[i], calc)
			}
		}
	}
}
