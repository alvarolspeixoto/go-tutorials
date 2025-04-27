package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitorings = 5
const delay = 5

func main() {

	var name string

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	introduce(name)

	for {
		showMenu()
		switchOption(readOption())
	}

}

func introduce(name string) {
	if name == "" {
		name = "Visitante"
	}

	version := 1.1
	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func switchOption(option int) {
	switch option {
	case 1:
		startMonitoring()
	case 2:
		readLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)

	default:
		fmt.Println("Opção inválida")
	}
}

func readOption() int {
	var option int
	_, err := fmt.Scan(&option)

	if err != nil {
		fmt.Println("Opção inválida! Digite um número")
	}

	fmt.Println("O comando escolhido foi", option)

	return option
}

func testSite(site string) {
	fmt.Println("Testando o site: ", site)
	response, err := http.Get(site)

	writeLog(site, err == nil && response.StatusCode == 200)

	if err != nil || response.StatusCode != 200 {
		fmt.Printf("Erro ao fazer a requisição ao site %s. Status Code: %d\n\n", site, response.StatusCode)
		return
	}

	defer response.Body.Close()
	fmt.Printf("Requisição bem-sucedida! (%s) Status Code: %d\n\n", site, response.StatusCode)
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites, err := getSitesFromFile("sites.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)
		}

		fmt.Println("Aguardando 5 segundos para a próxima verificação...")
		time.Sleep(delay * time.Second)
	}

}

func getSitesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	var sites []string

	for err == nil {
		line, readerErr := reader.ReadString('\n')
		err = readerErr

		sites = append(sites, strings.TrimSpace(line))

	}

	return sites, nil
}

func writeLog(log string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	statusMessage := "Online"

	if !status {
		statusMessage = "Offline"
	}

	writer.WriteString(time.Now().Format("02/01/2006 15:04:05") + " | " + log + " | " + statusMessage + "\n")

	writer.Flush()
}

func readLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo de log:", err)
		return
	}

	fmt.Println("Exibindo Logs...")
	fmt.Println(string(file))
}
