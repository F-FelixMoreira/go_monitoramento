package main // Principal pacote da nossa aplicação

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	comandoLido := lerComando()
	exibirMenu()

	switch comandoLido {
	case 0:
		fmt.Println("Saindo do programa")
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Sair do programa")
		sairDoPrograma()
	default:
		fmt.Println("Opção inválida. Tente novamente.")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Felix"
	versao := 1.1
	fmt.Println("Ola,", nome)
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	url := "https://httpbin.org/status/200" // ou /404
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("Status do request: OK")
		fmt.Println(resp)
	} else {
		fmt.Println("Status do request: Não OK")
	}

}

//https://httpbin.org/status/200
//https://httpbin.org/status/404
