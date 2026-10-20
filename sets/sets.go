/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sets

import (
	"cmp"
	"sort"

	"k8s.io/utils/set"
)

// Set is an alias for set.Set so that code using apimachinery sets.Set can be
// migrated without having to change all the corresponding code.
type Set[E comparable] = set.Set[E]

// Empty is public since it is used by some internal API objects for conversions between external
// string arrays and internal sets, and conversion logic requires public types today.
type Empty = set.Empty

// New creates a new set.
// Deprecated: use set.New instead.
func New[E comparable](items ...E) Set[E] {
	ss := make(Set[E], len(items))
	ss.Insert(items...)
	return ss
}

// KeySet creates a Set[E] from a keys of a map[E](? extends interface{}).
func KeySet[E comparable, A any](theMap map[E]A) Set[E] {
	ret := Set[E]{}
	for key := range theMap {
		ret.Insert(key)
	}
	return ret
}

type sortableSlice[E cmp.Ordered] []E

func (s sortableSlice[E]) Len() int {
	return len(s)
}
func (s sortableSlice[E]) Less(i, j int) bool { return s[i] < s[j] }
func (s sortableSlice[E]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted T slice.
//
// This is a separate function and not a method because not all types supported
// by Generic are ordered and only those can be sorted.
func List[E cmp.Ordered](s Set[E]) []E {
	res := make(sortableSlice[E], 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return res
}
