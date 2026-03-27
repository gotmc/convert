// Copyright (c) 2023–2024 The convert developers. All rights reserved.
// Project site: https://github.com/gotmc/convert
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// StripDoubleQuotes will strip double quotes at the beginning and end of a
// string.
func StripDoubleQuotes(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "\"")
	return strings.TrimSuffix(s, "\"")
}

// StringToNFloats uses the given separator to split a string into the
// expected number of floats.
func StringToNFloats(s, sep string, numExpected int) ([]float64, error) {
	floats, err := StringToFloats(s, sep)
	if err != nil {
		return nil, err
	}
	if len(floats) != numExpected {
		return nil, fmt.Errorf(
			"got %d floats, want %d for string: %s", len(floats), numExpected, s,
		)
	}
	return floats, nil
}

// StringToFloats uses the given separator to split a string into floats.
func StringToFloats(s, sep string) ([]float64, error) {
	parts := strings.Split(s, sep)
	nums := make([]float64, len(parts))
	for i, part := range parts {
		num, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, fmt.Errorf("converting %q to float64: %w", part, err)
		}
		nums[i] = num
	}
	return nums, nil
}
