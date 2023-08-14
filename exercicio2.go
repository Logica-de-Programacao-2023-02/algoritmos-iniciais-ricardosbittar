package algoritmos_iniciais_ricardosbittar
import "fmt"
func main() {
var tempF int
var tempC int
var resposta string
float := 1.8

fmt.Println("Qual a temperatura que você tem para converter?")
fmt.Scanln(&resposta)
 if resposta == "celsius"{
	 fmt.Println("Digite a quantidade de CELSIUS você quer converter para FAHRENHEIT: ")
	 fmt.Scanln(&tempC)
	  convTempF := tempC * float + 32
	 fmt.Println("A temperatura em Fahrenheit de " tempC " é:", convTempF)


 }
 if resposta == "fahrenheit"{
	 fmt.Scanln(&resposta)
	 fmt.Println("Digite a quantidade de FAHRENHEIT você quer converter para CELSIUS: ")
	 fmt.Println(&tempF)
	 convTempC := tempF / float - 32
	 fmt.Println("A temperatura em CELSIUS de" tempF " é:", convTempC)

 }else{
	 fmt.Println("Unidade de medida inválido")
 }



}

