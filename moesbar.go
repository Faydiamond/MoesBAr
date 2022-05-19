package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("A continuacion digita tu edad ")
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	age := scaner.Text()
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		//log.Fatal(err)
		println("Por favor digite un numero valido!")
	} else if ageInt >= 21 {
		println("Sos mayor de edad, bienvenido al bar de Moe!")
	} else {
		println("Por tu edad, no puedes ingresar  al bar de Moe!")
	}
	//fmt.Println(ageInt)
}
