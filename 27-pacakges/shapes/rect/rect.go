package rect

import "fmt"

type Rect struct {
	L, B float32
}

// display cannot be accessed outside of this package
func (r *Rect) display() {
	fmt.Printf("Provided Length:%f Bredth:%f", r.L, r.B)
}

// Go does not have access specifiers/modifiers
// any type that starts with Uppercase is simialr to Public/ unrestricted or unexported
// if any type or variable starts with lower case then it is restricted or exported
