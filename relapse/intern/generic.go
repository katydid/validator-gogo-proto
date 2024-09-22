//  Copyright 2017 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package intern

// filter returns a list of all items in the list that matches the predicate.
func filter[T any](predicate func(T) bool, list []T) []T {
	j := 0
	for i, elem := range list {
		if predicate(elem) {
			if i != j {
				list[j] = list[i]
			}
			j++
		}
	}
	return list[:j]
}

// isAny reports whether the predicate returns true for any of the elements in the given slice.
func isAny[T any](pred func(T) bool, list []T) bool {
	for _, elem := range list {
		if pred(elem) {
			return true
		}
	}
	return false
}

// areAll reports whether the predicate returns true for all of the elements in the given slice.
func areAll[T any](predicate func(T) bool, slice []T) bool {
	for _, elem := range slice {
		if !predicate(elem) {
			return false
		}
	}
	return true
}

// traverse returns a list where each element of the input list has been morphed by the input function or an error.
func traverse[A, B any](f func(A) (B, error), list []A) ([]B, error) {
	out := make([]B, len(list))
	var err error
	for i, elem := range list {
		out[i], err = f(elem)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
