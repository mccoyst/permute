// © 2014 Steve McCoy. Licensed under the MIT license. See LICENSE for details.

/*
Package permute provides functions to generate in-place permutations on arbitrary sequences.
*/
package permute

// Interface is identical to the standard "sort" package's Interface.
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// NextLex tries to generate the next lexicographical permutation of s and returns true iff it succeeds.
func NextLex(s Interface) bool {
	k := s.Len() - 2
	for ; k >= 0; k-- {
		if s.Less(k, k+1) {
			break
		}
	}

	if k < 0 || !s.Less(k, k+1) {
		return false
	}

	i := s.Len() - 1
	for ; i > k; i-- {
		if s.Less(k, i) {
			break
		}
	}

	if i == k {
		return false
	}

	s.Swap(k, i)
	for j := 1; j < s.Len()-k-1; j++ {
		s.Swap(k+j, s.Len()-j)
	}
	return true
}
