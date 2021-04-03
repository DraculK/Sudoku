package sudoku

import (
	"fmt"
	"strings"
)

var (
	digitosValidos = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	digitos        = "123456789"
	linhas         = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	colunas         = digitosValidos
	quadrados       = cruzamento(linhas, colunas)
	listaCaracteres = getListaCaracteres()
	encontraCelula  = getEncontraCelula()
)

//Função para juntar as Linhas com as Conlunas, resultando em [A1,A2..I9]
func cruzamento(linha, coluna []string) []string {
	linha_Coluna := []string{}

	for _, valorLinha := range linha {
		for _, valorColuna := range coluna {
			linha_Coluna = append(linha_Coluna, fmt.Sprint(valorLinha, valorColuna))
		}
	}
	return linha_Coluna
}

//Retorna as linhas, as colunas e os quadrados
func getListaCaracteres() [][]string {
	cLista := make([][]string, len(linhas)*3)

	iterador := 0
	for _, valorColuna := range colunas {
		coluna := cruzamento(linhas, []string{valorColuna})
		cLista[iterador] = coluna
		iterador++
	}

	for _, valorLinha := range linhas {
		linha := cruzamento([]string{valorLinha}, colunas)
		cLista[iterador] = linha
		iterador++
	}

	linhasVetor := []string{`A B C`, `D E F`, `G H I`}
	colunasVetor := []string{`1 2 3`, `4 5 6`, `7 8 9`}

	for _, valorLinha := range linhasVetor {
		for _, valorColuna := range colunasVetor {

			quadrado := cruzamento(strings.Fields(valorLinha), strings.Fields(valorColuna))
			cLista[iterador] = quadrado
			iterador++
		}
	}

	return cLista
}

//Função para encontrar valores por meio das chaves do map. Assim, conseguimos encontrar a linha, a coluna e o quadrado da unidade em questão.
func getEncontraCelula() map[string][][]string {
	mapEncontra := make(map[string][][]string)

	for _, s := range quadrados {
		celula := make([][]string, 3)
		iterador := 0
		for _, unidade := range listaCaracteres {
			for _, chave := range unidade {
				if s == chave {
					celula[iterador] = unidade
					mapEncontra[s] = celula
					iterador++
					break
				}
			}

		}
	}

	return mapEncontra
}
