package main

import "fmt"

func main() {
	//byte is alias for uint8 (utf8)
	//rune is alias for int32 (utf32)
	//there is no char type, use byte(ascii) and rune(unicode)
	//rune is default

	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	r := 'a' //rune is alias for int32 (utf32)
	fmt.Printf("%v, %T", r, r)
}
