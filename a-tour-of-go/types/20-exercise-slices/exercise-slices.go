package main

import (
	_ "fmt"
	"golang.org/x/tour/pic"
)

/*
Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is
a slice of dx 8-bit unsigned integers. When you run the program, it will
display your picture, interpreting the integers as grayscale (well, bluescale)
values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y,
and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

func Pic(dx, dy int) [][]uint8 {
	// fmt.Println(dx, dy)
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// image[y][x] = uint8(x ^ y)
			image[y][x] = uint8((x + y) / 2)
			// image[y][x] = uint8(x * y)
			// image[y][x] = uint8(x | y)
			// fmt.Println(image[y][x])
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
	/*
	   IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAAEEElEQVR42uzdQYrrMBCE4ZbJxXW0PkqOkGUWwcPADMwsvPBCxlJ/YH6KQhjzYJ5SqbT1iNhbC5er5rW1FvseiDX5+P4j2OLz+f2DoOlK2g6AdoDt58/i/f7/8YjPX923A6AdwGdBWgZArJ4B/l6v1/FXp9Zbv8p6OwDaAXwWpGUARD2A74n5egBEGYCmZQDEqj3A0fV8nv/Jtfu7/13vbwdAO4DPgrQMgKgH8D0xXw+AKAPQtAyAaB7A78utNw+AKAPQtAyAqAfg8/UAiDIATcsAiFXmAc5emTH8Fe+e3/M7HwBRBqBpGQBRD8Dn6wEQZQCalgEQzQNYb70eAFEGoGkZAFEPwNcD+J8AZQCfC2kZAFEP4P3x7u98AEQZgKZlAEQ9AJ+vB0CUAWhaBkA0D2C99eYBEGUAmpYBEPUAfL4eAFEGoGkZALH6+QCjr94nfvgL3q/v3988AKIMQNMyAKIegM+XARBlAJq2A6AdwO/FrTcPgCgD0LQMgKgH4PP1AIgyAE3LAIjmAcL76d3f+QCIMgBNywCIegA+Xw+AKAPQtAyAaB7AeuvNAyDKADQtAyDqAfh8PQCiDEDTMgBiifMBMmPq9+t7/rWf3w6AdgCfBWkZAFEP4Htivh4AUQagaRkA0TyA35dbbx4AUQagaRkAUQ/A5+sBEGUAmpYBEJ0PEN5P7/7OB0CUAWhaBkDUA/D5egBEGYCmZQBE8wDWW28eAFEGoGkZAFEPwOfrARBlAJqWARAXOR/A5brzZQdAO4DPgrQMgKgH8D0xXw+AKAPQtAyAaB7A78utNw+AKAPQtAyAqAfg8/UAiDIATcsAiM4HCO+nd3/nAyDKADQtAyDqAfh8PQCiDEDTMgCieQDrrTcPgCgD0LQMgKgH4PP1AIgyAE3LAIiLnA+QGcNf8e75Pb/zARBlAJqWARD1AHy+HgBRBqBpGQDRPID11usBEGUAmpYBEPUAfD2A/wlQBvC5kJYBEPUA3h/v/s4HQJQBaFoGQNQD8Pl6AEQZgKZlAETzANZbbx4AUQagaRkAUQ/A5+sBEGUAmpYBEKufDzD66n3ih7/g/fr+/c0DIMoANC0DIOoB+HwZAFEGoGk7ANoB/F7cevMAiDIATcsAiHoAPl8PgCgD0LQMgGgeILyf3v2dD4AoA9C0DICoB+Dz9QCIMgBNywCI5gGst948AKIMQNMyAKIegM/XAyDKADQtAyCWOB8gM6Z+v77nX/v57QBoB/BZkJYBEPUAvifm6wEQZQCalgEQzQP4fbn15gEQZQCalgEQ9QB8vh4AUQagaRkA0fkA4f307u98AEQZgKZlAEQ9AJ+vB0CUAWhaBkA0D2C99eYBEGUAmpYBEPUAfL4eAFEGoGkZAPEafgUAAP//2P/knioA0wcAAAAASUVORK5CYII=
	*/
}
