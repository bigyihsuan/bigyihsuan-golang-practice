package subfolder

import "fmt"

func TestPrint() {
	fmt.Println("test print")
}

func TestPrivatePrint() {
	privatePrint()
}

func privatePrint() {
	fmt.Println("private print")
}
