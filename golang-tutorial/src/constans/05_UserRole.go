package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHead
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Head? %v\n", isHead&roles == isAdmin)
}
