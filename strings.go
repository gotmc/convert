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
	s = strings.TrimPrefix(s, "\"")
	return strings.TrimSuffix(s, "\"")
}

// StringToNFloats uses the given separator to split a string into the
// expected number of floats.
func StringToNFloats(s, sep string, numExpected int) ([]float64, error) {
	slice := strings.Split(s, sep)
	if len(slice) != numExpected {
		return nil, fmt.Errorf(
			"error: didn't split into number expected given string: %s", s,
		)
	}
	nums := make([]float64, numExpected)
	for i, s := range slice {
		num, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to float64", s)
		}
		nums[i] = num
	}
	return nums, nil
}

// StringToFloats uses the given separator to split a string into an unknown
// number of floats.
func StringToFloats(s, sep string) ([]float64, error) {
	slice := strings.Split(s, sep)
	if len(slice) < 1 {
		return nil, fmt.Errorf(
			"error splitting the given string: %s", s,
		)
	}
	nums := make([]float64, len(slice))
	for i, s := range slice {
		num, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to float64", s)
		}
		nums[i] = num
	}
	return nums, nil
}
