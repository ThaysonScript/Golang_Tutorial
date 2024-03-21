package main

import "fmt"

func main() {
	// Go tem três funções para impressao
    fmt.Print()		// imprimir sem \n no fim
    fmt.Println()	// imprimir com \n no fim
    fmt.Printf()	// imprimir com formatacao


	/*
			Os seguintes verbos podem ser usados ​​com todos os tipos de dados:
			Verb 	Description
			%v 		Prints the value in the default format
			%#v 	Prints the value in Go-syntax format
			%T 		Prints the type of the value
			%% 		Prints the % sign
	*/

	/*

		Os seguintes verbos podem ser usados ​​com o tipo de dados inteiro:
		Verb 	Description
		%b 		Base 2
		%d 		Base 10
		%+d 	Base 10 and always show sign
		%o 		Base 8
		%O 		Base 8, with leading 0o
		%x 		Base 16, lowercase
		%X 		Base 16, uppercase
		%#x 	Base 16, with leading 0x
		%4d 	Pad with spaces (width 4, right justified)
		%-4d 	Pad with spaces (width 4, left justified)
		%04d 	Pad with zeroes (width 4

	*/

	/*
	
			Os seguintes verbos podem ser usados ​​com o tipo de dados string:
			Verb 	Description
			%s 		Prints the value as plain string
			%q 		Prints the value as a double-quoted string
			%8s 	Prints the value as plain string (width 8, right justified)
			%-8s 	Prints the value as plain string (width 8, left justified)
			%x 		Prints the value as hex dump of byte values
			% x 	Prints the value as hex dump with spaces
	
	*/

	/*
	
			O seguinte verbo pode ser usado com o tipo de dados booleano:
			Verb 	Description
			%t 		Value of the boolean operator in true or false format (same as using %v)
	
	*/

	/*
	
			Os seguintes verbos podem ser usados ​​com o tipo de dados float:
			Verb 	Description
			%e 		Scientific notation with 'e' as exponent
			%f 		Decimal point, no exponent
			%.2f 	Default width, precision 2
			%6.2f 	Width 6, precision 2
			%g 		Exponent as needed, only necessary digits
	
	*/
}