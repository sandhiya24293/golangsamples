package main

import (
	"fmt"
	"unicode/utf8"
)


func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%c ",s[i])
	}
}

func printCharsByRune(s string) {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}



func main() {
	//A string in Go is a slice of bytes.
	//Strings can be created by enclosing their contents inside " ".

	name := "Hello World"
	fmt.Println(name)

	//Accessing individual bytes of a string
	//Since a string is a slice of bytes, its possible to access each byte of a string.

	//printBytes
	for i:= 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	//Output===48 65 6c 6c 6f 20 57 6f 72 6c 64.     Unicode UT8-encoded values of "Hello World"
	//len(s) returns the number of bytes in the string and we use a for loop to print those bytes in hexadecimal notation.
	//%x is the format specifier for hexadecimal.

	//printChars
	for i:= 0; i < len(name); i++ {
		fmt.Printf("%c ",name[i])
	}
	//Output=== H e l l o   W o r l d
	//%c format specifier is used to print the characters of the string. The program outputs


	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	//We are trying to print the characters of Señor and it outputs S e Ã ± o r which is wrong.
	// Why does this program break for Señor when its perfectly alright with Hello World.
	// The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
	// We are trying to print characters assuming that each code point will be one byte long which is wrong.
	// In UTF-8 encoding a code point can occupy more than 1 byte.
	// So how do we solve this. This is where rune saves us.

	//A rune is a builtin type in Go and its the alias of int32.
	//rune represents a Unicode code point in Go.
	// It does not matter how many bytes the code point occupies, it can be represented by a rune.

	fmt.Printf("\n")
	printCharsByRune(name)

	//for range loop on a string
	//The above program (printCharsByRune)is a perfect way to iterate over the individual runes of a string.
	// But Go offers us a much easier way(printCharsAndBytes) to do this using the for range loop.

	printCharsAndBytes(name)

	//Constructing string from slice of bytes
	hexbyteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(hexbyteSlice)
	fmt.Println(str)//outputs Café.

	decimalbyteSlice := []byte{67, 97, 102, 195, 169}//decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	strx := string(decimalbyteSlice)
	fmt.Println(strx)//output Café

	//Constructing a string from slice of runes
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	stry := string(runeSlice)
	fmt.Println(stry)//outputs Señor

	//Length of the string
	//The func RuneCountInString(s string) (n int) function of the utf8 package is used to find the length of the string.
	//This method takes a string as argument and returns the number of runes in it.
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	//Strings are immutable
	//Strings are immutable in Go. Once a string is created its not possible to change it.
	//To workaround this string immutability, strings are converted to a slice of runes.
	// Then that slice is mutated with whatever changes needed and converted back to a new string.

	//EXCERCISE
	//** User should enter the value at runtime **//
	//1) Write code to modify a string.
	//2) Write code to compare two string
	//3) Try to find the length of "Señor" by simple 'len' function and also by 'utf8.RuneCountInString'?
}
