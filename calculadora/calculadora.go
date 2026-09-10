package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return float32(a) / float32(b), nil
}

func Fibonnacio(a int) int {
	if a <= 1 {
		return a
	}
	return Fibonnacio(a-1) + Fibonnacio(a-2)
}

func Selecionar(condition int) bool {

	switch condition {
	case 1:
		fmt.Println("Digite os valores da soma :")
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Resultado: %d\n", Add(a, b))
		return false
	case 2:
		fmt.Println("Digite os valores da subtração :")
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Resultado: %d\n", Subtract(a, b))
		return false
	case 3:
		fmt.Println("Digite os valores da multiplicação :")
		var a, b int
		fmt.Scan(&a, &b)
		fmt.Printf("Resultado: %d\n", Multiply(a, b))
		return false
	case 4:
		fmt.Println("Digite os valores da divisão :")
		var a, b int
		fmt.Scan(&a, &b)
		result, err := Divide(a, b)
		if err != nil {
			fmt.Println("Não é possível dividir por zero:", err)
		} else {
			fmt.Printf("Resultado: %.2f\n", result)
		}
		return false
	case 5:
		fmt.Println("Digite o valor para calcular o fibonacci :")
		var a int
		fmt.Scan(&a)
		fmt.Printf("Resultado: %d\n", Fibonnacio(a))
		return false
	case 6:
		fmt.Println("Saindo da calculadora...")
		return true
	default:
		fmt.Println("Opção inválida. Tente novamente.")
		return false
	}
}

func Menu() {
	fmt.Println("Calculadora em Go")
	fmt.Println("------------------")
	fmt.Println("Escolha uma opção:")
	fmt.Println("[1] Soma")
	fmt.Println("[2] Subtração")
	fmt.Println("[3] Multiplicação")
	fmt.Println("[4] Divisão")
	fmt.Println("[5] Fibonacci")
	fmt.Println("[6] Sair")
}

