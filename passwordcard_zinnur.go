package main

import (
	"fmt"
	"github.com/ActiveState/golor"
	"math/rand"
)

func myrandom(x int) string {
	alphi := []string{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	numeric := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	loweralphi := []string{"a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var line string

	for k := 0; k < 10; k++ {
		line += alphi[rand.Intn(x+5)]
		line += " "
		line += numeric[rand.Intn(9)]
		line += " "
		line += loweralphi[rand.Intn(x+3)]
		line += " "
	}

	return line

}

func main() {
	signals := "■ □ ▲ △ ○ ● ★ ☂ ☀ ☁ ☹ ☺ ♠ ♣ ♥ ♦ ♫ € ¥ £ $ ! ? ¡ * ¿ ⊙ ◐ ◩ �"
	fmt.Println("    " + signals)
	fmt.Println(golor.Colorize("01. "+myrandom(12), golor.Y, -1))
	fmt.Println(golor.Colorize("02. "+myrandom(14), golor.BLUE, -1))
	fmt.Println(golor.Colorize("03. "+myrandom(15), golor.R, -1))
	fmt.Println(golor.Colorize("04. "+myrandom(9), golor.C, -1))
	fmt.Println(golor.Colorize("05. "+myrandom(17), golor.Y, -1))
	fmt.Println(golor.Colorize("06. "+myrandom(7), golor.G, -1))
	fmt.Println(golor.Colorize("07. "+myrandom(13), golor.M, -1))
	fmt.Println(golor.Colorize("08. "+myrandom(20), golor.GRAY, -1))

}