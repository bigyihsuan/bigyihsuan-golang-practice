package characters

import "fmt"

type character interface {
	fmt.Stringer
	TakeDamage(int)
}
