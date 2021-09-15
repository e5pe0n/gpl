package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "Feburary", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // [April May June]
	fmt.Println(summer) // [June July August]

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	fmt.Printf("%x\n", runes) // [48 65 6c 6c 6f 2c 20 4e16 754c]
}
