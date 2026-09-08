package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	conteudo, err := os.ReadFile("test1.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
	fmt.Println(soma(conteudo))
}

func soma(conteudo []byte) int {

	var numeros []any
	err := json.Unmarshal(conteudo, &numeros)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return 0
	}

	result := 0
	for _, numero := range numeros {
		switch valor := numero.(type) {
		case float64:
			result += int(valor)

		case string:
			num, err := strconv.Atoi(valor)
			if err == nil {
				result += num
			}
		}

	}

	return result
}
