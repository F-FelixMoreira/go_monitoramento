package main // Principal pacote da nossa aplicação

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	exibirIntroducao()
	sliceSites := lerSitesArquivos("sites.txt")

	for {
		exibirMenu()
		comandoLido := lerComando()

		switch comandoLido {
		case 0:
			fmt.Println("Saindo do programa")
		case 1:
			iniciarMonitoramento(sliceSites)
		case 2:
			fmt.Println("Sair do programa")
			sairDoPrograma()
		default:
			fmt.Println("Opção inválida. Tente novamente.")
			os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Felix"
	versao := 1.1
	fmt.Println("\nOla,", nome)
	fmt.Println("Você está utilizando a versão:", versao)
}

func lerComando() int {
	var input int
	fmt.Scan(&input)
	fmt.Println("Comando escolhido foi:", input)

	return input
}

func exibirMenu() {
	fmt.Println("1 -Iniciar monitoramento")
	fmt.Println("2 - Exibir log")
	fmt.Println("3 - Sair do programa")
}

func sairDoPrograma() {
	os.Exit(0)
}

func iniciarMonitoramento(sites []string) {
	fmt.Println("\n__________________\n")
	fmt.Println("Monitorando...")

	for idx, site := range sites {
		fmt.Println("\nMonitorando site nº: ", idx+1)
		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("\nStatus do request: OK")
			fmt.Println(resp)
		} else {
			fmt.Println("\nStatus do request: Não OK")
		}
	}
}

func lerSitesArquivos(endereco string) []string {

	var sites []string

	arquivo, err := os.Open(endereco)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		os.Exit(1)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			fmt.Println("lista de sites:", sites)
			break
		}
	}

	return sites

}
