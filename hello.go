package main // Principal pacote da nossa aplicação

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			fmt.Println("Exibição de log inativa.")
			sairDoPrograma()
		case 3:
			fmt.Println("Programa encerrado, até mais :)")
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
	fmt.Println("\n__________ MENU _________\n")
	fmt.Println("\n1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir log")
	fmt.Println("3 - Sair do programa")
}

func sairDoPrograma() {
	os.Exit(0)
}

func iniciarMonitoramento(sites []string) {
	fmt.Println("\n__________________\n")
	fmt.Println("Monitoramento iniciado...")

	for idx, site := range sites {
		fmt.Println("\nMonitorando site nº: ", idx+1)
		testarSite(site)
	}

	fmt.Println("\nMonitoramento encerrado.")
}

func testarSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Houve um erro ocorrido ao tentar ler o site: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("\n O site", site, "está funcional.")
		registrarLog(site, true)
	} else {
		fmt.Println("\n O site", site, "está não está funcional. Status Code:",
			resp.StatusCode)
		registrarLog(site, false)
	}
}

func lerSitesArquivos(endereco_arquivo_sites string) []string {
	var sites []string

	arquivo, err := os.Open(endereco_arquivo_sites)

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
	arquivo.Close()
	return sites
}

func registrarLog(site string, status bool) {
	//horario, site, status, erro
	arquivo, err := os.OpenFile("log_sites_status.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	horario := time.Now().Format("02/01/2006 15:04:05")
	arquivo.WriteString(horario + " | " + site + " | status: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
