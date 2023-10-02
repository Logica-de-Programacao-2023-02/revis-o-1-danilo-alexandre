package q5

import (
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	// Inicializa uma string vazia para armazenar o resultado.
	resultado := ""

	// Define uma string contendo todas as vogais em minúsculas.
	vogais := "aeiouAEIOU"

	// Itera pelos caracteres da string fornecida.
	for _, char := range s {
		charStr := string(char) // Converte o caractere para string.

		// Verifica se o caractere não é uma vogal.
		if !strings.ContainsAny(charStr, vogais) {
			// Se não for uma vogal, adiciona um "." antes do caractere.
			resultado += "." + charStr
			resultado = strings.ToLower(resultado)
		}
	}

	return resultado
}
