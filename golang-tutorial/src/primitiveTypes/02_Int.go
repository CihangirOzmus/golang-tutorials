package main

import "fmt"

func main() {
	var i int   //platform dependent
	var j int8  //-128 127
	var k int16 //-32768 32767
	var l int32 //-2147483648 2147483647
	var m int64 //-9223372036854775808 9223372036854775807

	var ui uint   //platform dependent
	var uj uint8  //0 255
	var uk uint16 //0 65535
	var ul uint32 //0 4294967295
	var um uint64 //0 to big

	fmt.Println(i, j, k, l, m, ui, uj, uk, ul, um)
}
