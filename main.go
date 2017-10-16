package main

import (
	"fmt"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	r := levenshtein.RatioForStrings([]rune("障害対策"), []rune("今回は障害対策をしたいいいいいいいいいいいいいい"), levenshtein.DefaultOptions)
	r2 := levenshtein.RatioForStrings([]rune("障害対策"), []rune("よおおおおおおお"), levenshtein.DefaultOptions)
	fmt.Println("RatioForStrings", r, r2)
	e := levenshtein.EditScriptForStrings([]rune("障害対策"), []rune("今回は障害対策をしたいいいいいいいいいいいいいい"), levenshtein.DefaultOptions)
	e2 := levenshtein.EditScriptForStrings([]rune("障害対策"), []rune("よおおおおおおお"), levenshtein.DefaultOptions)
	fmt.Println("EditScriptForStrings", e, e2)
	d := levenshtein.DistanceForStrings([]rune("障害対策"), []rune("今回は障害対策をしたいいいいいいいいいいいいいい"), levenshtein.DefaultOptions)
	d2 := levenshtein.DistanceForStrings([]rune("障害対策"), []rune("よおおおおおおお"), levenshtein.DefaultOptions)
	fmt.Println("DistanceForStrings", d, d2)
}
