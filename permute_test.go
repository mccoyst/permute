// © 2014 Steve McCoy. Licensed under the MIT license. See LICENSE for details.

package permute

import "fmt"

type ints []int

func (n ints) Len() int { return len(n) }
func (n ints) Less(i, j int) bool { return n[i] < n[j] }
func (n ints) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func ExamplePermute() {
	n := ints{ 1, 2, 3 }

	for {
		fmt.Println(n)
		if !Next(n) {
			break
		}
	}
	fmt.Println(n)

	// output: [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 1 2]
	// [3 2 1]
	// [3 2 1]
}
