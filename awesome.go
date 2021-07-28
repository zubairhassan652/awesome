package awesome

import (
	"fmt"
	"github.com/zubairhassan652/awesome/morestrings"
	"rsc.io/quote"
)

// This method calls two method one from our own created package and other from quote package.
func CallToSubPackages()  {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(quote.Go())
}
