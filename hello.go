package main // Principal pacote da nossa aplicação

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Adicionar URLs para avaliar neste slice.
	sliceSites := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
	}

	exibirIntroducao()

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
