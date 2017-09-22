// Nonempty is an example of an in-place slice algorithm
package main

// nonempty return a slice holding only the non=empty stirngs.
//the underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
