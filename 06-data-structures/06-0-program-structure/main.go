package main

import (
	"mymodule/another_dep"
	"mymodule/another_dep/nested"
	"mymodule/subfolder"
)

func main() {
	subfolder.TestPrint()
	subfolder.TestPrivatePrint()
	another_dep.AnotherDep()
	nested.Nested()
}
