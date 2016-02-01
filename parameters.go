package headerquality

import "fmt"

// Parameters - slice of sortable Parameters.
type Parameters []Parameter

// Returns the total number of items in the slice. Implemented to satisfy
// sort.Interface.
func (al Parameters) Len() int { return len(al) }

// Swaps the items at position i and j. Implemented to satisfy sort.Interface.
func (al Parameters) Swap(i, j int) { al[i], al[j] = al[j], al[i] }

// Determines whether or not the item at position i is "less than" the item
// at position j. Implemented to satisfy sort.Interface.
func (al Parameters) Less(i, j int) bool { return al[i].Quality > al[j].Quality }

// Returns the parsed factors in a human readable fashion.
func (al Parameters) String() (output string) {
	for i, factor := range al {
		output += fmt.Sprintf("%s (%1.1f)", factor.Factor, factor.Quality)
		if i != len(al)-1 {
			output += ", "
		}
	}

	if len(output) == 0 {
		return "[]"
	}

	return output
}
