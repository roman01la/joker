// This file is generated by generate-std.joke script. Do not edit manually!

// +build !fast_init

package string

import (
	. "github.com/candid82/joker/core"
	"fmt"
	"os"
)

func InternsOrThunks() {
	if VerbosityLevel > 0 {
		fmt.Fprintln(os.Stderr, "Lazily running slow version of string.InternsOrThunks().")
	}
	stringNamespace.ResetMeta(MakeMeta(nil, `Implements simple functions to manipulate strings.`, "1.0"))

	
	stringNamespace.InternVar("blank?", isblank_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`True if s is nil, empty, or contains only whitespace.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("capitalize", capitalize_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts first character of the string to upper-case, all other
  characters to lower-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("ends-with?", isends_with_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s ends with substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("escape", escape_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("cmap"))),
			`Return a new string, using cmap to escape each character ch
  from s as follows:

  If (cmap ch) is nil, append ch to the new string.
  If (cmap ch) is non-nil, append (str (cmap ch)) instead.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("includes?", isincludes_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s includes substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("index-of", index_of_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
			`Return index of value (string or char) in s, optionally searching
  forward from from or nil if not found.`, "1.0"))

	stringNamespace.InternVar("join", join_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("coll")), NewVectorFrom(MakeSymbol("separator"), MakeSymbol("coll"))),
			`Returns a string of all elements in coll, as returned by (seq coll), separated by an optional separator.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("last-index-of", last_index_of_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("value")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("value"), MakeSymbol("from"))),
			`Return last index of value (string or char) in s, optionally
  searching backward from from or nil if not found.`, "1.0"))

	stringNamespace.InternVar("lower-case", lower_case_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts string to all lower-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("pad-left", pad_left_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
			`Returns s padded with pad at the beginning to length n.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("pad-right", pad_right_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("pad"), MakeSymbol("n"))),
			`Returns s padded with pad at the end to length n.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("re-quote", re_quote_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Returns an instance of Regex that matches the string exactly`, "1.0").Plus(MakeKeyword("tag"), String{S: "Regex"}))

	stringNamespace.InternVar("replace", replace_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
			`Replaces all instances of match (String or Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("replace-first", replace_first_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("match"), MakeSymbol("repl"))),
			`Replaces the first instance of match (String or Regex) with string repl in string s.

  If match is Regex, $1, $2, etc. in the replacement string repl are
  substituted with the string that matched the corresponding
  parenthesized group in the pattern.
  `, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("reverse", reverse_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Returns s with its characters reversed.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("split", split_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("sep")), NewVectorFrom(MakeSymbol("s"), MakeSymbol("sep"), MakeSymbol("n"))),
			`Splits string on a string or regular expression. Returns vector of the splits.`, "1.0"))

	stringNamespace.InternVar("split-lines", split_lines_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Splits string on \n or \r\n. Returns vector of the splits.`, "1.0"))

	stringNamespace.InternVar("starts-with?", isstarts_with_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"), MakeSymbol("substr"))),
			`True if s starts with substr.`, "1.0").Plus(MakeKeyword("tag"), String{S: "Boolean"}))

	stringNamespace.InternVar("trim", trim_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from both ends of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-left", trim_left_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the left side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-newline", trim_newline_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes all trailing newline \n or return \r characters from string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trim-right", trim_right_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the right side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("triml", triml_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the left side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("trimr", trimr_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Removes whitespace from the right side of string.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

	stringNamespace.InternVar("upper-case", upper_case_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("s"))),
			`Converts string to all upper-case.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

}
