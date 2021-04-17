package main

import (
	"fmt"
	"strings"
	"errors"
)

var (
	digitosValidos = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	digitos        = "123456789"
	linhas         = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	colunas           = digitosValidos
	quadrados         = cruzamento(linhas, colunas)
	listaCaracteres   = getListaCaracteres()
	encontraCelula    = getEncontraCelula()
	checagemElementos = getChecagemElementos()
)

func main() {
	//fmt.Println(colunas)
	//fmt.Println(quadrados)
	//fmt.Println(listaCaracteres)
	//fmt.Println(encontraCelula)
	//fmt.Println(checagemElementos)
}

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

//Transforma um elemento em uma chave e pega todos os elementos que não podem ter o mesmo valor(peers).
func getChecagemElementos() map[string][]string {

	elementosTamanho := 20

	peersFinal := make(map[string][]string)
	for _, s := range quadrados {
		peer := make(map[string]interface{}, elementosTamanho)
		for _, unidade := range encontraCelula[s] {
			for _, chave := range unidade {
				if s != chave {
					peer[chave] = 1
				}
			}
		}

		separacaoPeers := []string{}
		for elemento := range peer {
			separacaoPeers = append(separacaoPeers, elemento)
		}
		peersFinal[s] = separacaoPeers

	}

	return peersFinal
}

// Converte o grid em um dicionário de {quadrado: caractere} com '.' para as casas vazias.
func valoresGrid(grid string) (map[string]string, error){
	caracteres := []string{}
	tamanhoGrid := len(grid)
	tamanhoCaracteres := len(caracteres)

	for iterador := 0; iterador < tamanhoGrid; iterador++{
		verifica := string(grid[iterador : iterador+1])
		if strings.Contains(digitos, verifica) || verifica == "0" || verifica == "."{
			caracteres = append(caracteres, verifica)
		}
	}

	if tamanhoCaracteres != 81 {
		return nil, errors.New("Tamanho do grid não é 81")
	}

	novoGrid := make(map[string]string)
	iterador := 0

	for _, s := range quadrados {
		novoGrid[s] = caracteres[iterador]
		iterador++
	}

	return novoGrid, nil
}
