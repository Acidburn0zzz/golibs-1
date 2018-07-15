// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright (c) 2018 Sung Pae <self@sungpae.com>
// Distributed under the MIT license.
// http://www.opensource.org/licenses/mit-license.php

package genericbenchmarks

// QuicksortIntSlice sorts a slice of int in place.
// Elements of type int must be comparable by value.
func QuicksortIntSlice(v []int) {
	switch len(v) {
	case 0, 1:
		return
	// Manually sort small slices
	case 2:
		if v[1] < v[0] {
			v[0], v[1] = v[1], v[0]
		}
		return
	case 3:
		if v[1] < v[0] {
			v[0], v[1] = v[1], v[0]
		}
		if v[2] < v[1] {
			v[1], v[2] = v[2], v[1]
		}
		if v[1] < v[0] {
			v[0], v[1] = v[1], v[0]
		}
		return
	}

	i := PartitionIntSlice(v)
	QuicksortIntSlice(v[:i+1])
	QuicksortIntSlice(v[i+1:])
}

// PartitionIntSlice partitions a slice of int in place
// such that every element 0..index is less than or equal to every element
// index+1..len(v-1). Elements of type int must be comparable by value.
func PartitionIntSlice(v []int) (index int) {
	// Hoare's partitioning with median of first, middle, and last as pivot
	var pivot int

	if len(v) > 16 {
		pivot = MedianOfThreeIntSamples(v)
	} else {
		pivot = v[(len(v)-1)/2]
	}

	i, j := -1, len(v)

	for {
		for {
			i++
			if v[i] >= pivot {
				break
			}
		}

		for {
			j--
			if v[j] <= pivot {
				break
			}
		}

		if i < j {
			v[i], v[j] = v[j], v[i]
		} else {
			return j
		}
	}
}

// MedianOfThreeIntSamples returns the median of the first, middle,
// and last element. Elements of type int must be comparable by value.
func MedianOfThreeIntSamples(v []int) int {
	a := v[0]
	b := v[(len(v)-1)/2]
	c := v[len(v)-1]

	if b < a {
		a, b = b, a
	}
	if c < b {
		b = c
	}
	if b < a {
		b = a
	}

	return b
}
