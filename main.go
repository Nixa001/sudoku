package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	arguments := os.Args[1:]
	if CheckSaisi(arguments) == true {
		table := [9][9]rune{}
		table = RempliTable(table, arguments)

		if Resolution(&table) == true {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					if x != 8 {
						z01.PrintRune(rune(table[y][x]))
						z01.PrintRune(32)
					} else {
						z01.PrintRune(rune(table[y][x]))
					}
				}
				z01.PrintRune(10)
			}
		} else {
			PrintStr("Error")
		}
	}
}

// La fonction checkSaisi confirme si l'entrée est valide,
// sinon, elle renvoie le message "Error".
func CheckSaisi(args []string) bool {
	if len(args) != 9 {
		PrintStr("Error")
		return false
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			PrintStr("Error")
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if value == '/' || value == '0' {
				PrintStr("Error")
				return false
			} else if value < '.' || value > '9' {
				PrintStr("Error")
				return false
			}
		}
	}
	return true
}

// La fonction rempliTable remplit le tableau de runes sudoku
// avec les valeurs du tableau de chaînes en entrée.
func RempliTable(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// La fonction Points renvoie un message vrai s'il existe encore
// cellules vides, et renvoie false s'il n'y en a pas.
func Points(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// La fonction SiValide itère sur les nombres
// qui peuvent potentiellement être placés dans des cellules vides
// et vérifie s'ils sont des candidats valides
// (ils remplissent chacune des 3 conditions
// pour le placement des nombres dans le jeu sudoku)

func SiValide(table *[9][9]rune, x int, y int, z rune) bool {
	// Verifier la repetition des int sur l'axe X
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}
	// Verifier la repetition des int sur l'axe Y
	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}

	// Verification du carre
	a := x / 3
	b := y / 3

	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// La fonction Resolution est un algorithme de backtracking
// qui utilise la récursion pour vérifier les solutions potentielles
// en commençant par 1.

func Resolution(table *[9][9]rune) bool {
	if !Points(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if SiValide(table, x, y, z) {
						table[y][x] = z
						if Resolution(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}
