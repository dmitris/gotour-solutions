package stringers

import "fmt"

// https://go.dev/tour/methods/18
// Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
//
// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
