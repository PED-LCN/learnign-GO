package main


import "fmt"

func main() {
	var opcao int
	for fecharCalculadora := false; !fecharCalculadora; {
		Menu()
		fmt.Scan(&opcao)
		fecharCalculadora = Selecionar(opcao)
	
	}
}