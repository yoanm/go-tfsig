package tfsig_test

import (
	"fmt"

	"github.com/yoanm/go-tfsig"
)

func ExampleToTerraformIdentifier() {
	fmt.Printf("a_valid-identifier becomes %s\n", tfsig.ToTerraformIdentifier("a_valid-identifier"))
	fmt.Printf(".github becomes %s\n", tfsig.ToTerraformIdentifier(".github"))
	fmt.Printf("an.identifier becomes %s\n", tfsig.ToTerraformIdentifier("an.identifier"))
	fmt.Printf("0id becomes %s\n", tfsig.ToTerraformIdentifier("0id"))

	// Output:
	// a_valid-identifier becomes a_valid-identifier
	// .github becomes _github
	// an.identifier becomes an-identifier
	// 0id becomes _id
}
