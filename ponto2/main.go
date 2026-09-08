package main

import (
	"fmt"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	text1, err := os.ReadFile("test2_text1.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err)
		return
	}
	text2, err := os.ReadFile("test2_text2.txt")
	if err != nil {
		fmt.Print("Error opening file: ", err)
		return
	}
	fmt.Print(compare(text1, text2))

}

func compare(text1, text2 []byte) bool {
	//Se forem comparados byte a byte usando .Equal() eles não são identicos pois um é
	//UTF-8 e outro ISO-8859-1 (latin1)
	//Apesar de visualmente serem identicos
	//Necessario converter para UTF-8 para verificar se os caracteres visualmente estão corretos

	txt2Convertido, err := charmap.ISO8859_1.NewDecoder().Bytes(text2)
	if err != nil {
		fmt.Println("Erro na conversão de encoding:", err)
		return false
	}

	texto1 := string(text1)
	texto2 := string(txt2Convertido)

	return texto1 == texto2
}
