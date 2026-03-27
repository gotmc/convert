// Copyright (c) 2023–2024 The convert developers. All rights reserved.
// Project site: https://github.com/gotmc/convert
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package convert

import (
	"testing"
)

func TestStringToFloats(t *testing.T) {
	testCases := []struct {
		name     string
		given    string
		sep      string
		expected []float64
	}{
		{
			name:     "three comma-separated decimals",
			given:    "0.001,0.002,0.003",
			sep:      ",",
			expected: []float64{0.001, 0.002, 0.003},
		},
		{
			name:     "four comma-separated floats",
			given:    "10.0,20.0,30.0,40.0",
			sep:      ",",
			expected: []float64{10.0, 20.0, 30.0, 40.0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calcs, err := StringToFloats(tc.given, tc.sep)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
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
		})
	}
}

func TestStringToFloats_errors(t *testing.T) {
	testCases := []struct {
		name  string
		given string
		sep   string
	}{
		{
			name:  "non-numeric value",
			given: "1.0,abc,3.0",
			sep:   ",",
		},
		{
			name:  "empty string",
			given: "",
			sep:   ",",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := StringToFloats(tc.given, tc.sep)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

func TestStringToNFloats(t *testing.T) {
	testCases := []struct {
		name        string
		given       string
		sep         string
		numExpected int
		expected    []float64
	}{
		{
			name:        "three comma-separated decimals",
			given:       "0.001,0.002,0.003",
			sep:         ",",
			numExpected: 3,
			expected:    []float64{0.001, 0.002, 0.003},
		},
		{
			name:        "four comma-separated floats",
			given:       "10.0,20.0,30.0,40.0",
			sep:         ",",
			numExpected: 4,
			expected:    []float64{10.0, 20.0, 30.0, 40.0},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calcs, err := StringToNFloats(tc.given, tc.sep, tc.numExpected)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
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
		})
	}
}

func TestStringToNFloats_errors(t *testing.T) {
	testCases := []struct {
		name        string
		given       string
		sep         string
		numExpected int
	}{
		{
			name:        "wrong count",
			given:       "1.0,2.0,3.0",
			sep:         ",",
			numExpected: 2,
		},
		{
			name:        "non-numeric value",
			given:       "1.0,abc",
			sep:         ",",
			numExpected: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := StringToNFloats(tc.given, tc.sep, tc.numExpected)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

func TestStripDoubleQuotes(t *testing.T) {
	testCases := []struct {
		name     string
		given    string
		expected string
	}{
		{
			name:     "quoted numeric values",
			given:    "\"0.001,0.002\"",
			expected: "0.001,0.002",
		},
		{
			name:     "quoted with trailing newline",
			given:    "\"0.001,0.002\"\n",
			expected: "0.001,0.002",
		},
		{
			name:     "quoted strings with trailing newline",
			given:    "\"foo,bash\"\n",
			expected: "foo,bash",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := StripDoubleQuotes(tc.given)
			if got != tc.expected {
				t.Errorf(
					"given %s / expected = %s / got = %s",
					tc.given,
					tc.expected,
					got,
				)
			}
		})
	}
}
