package main

import (
	"fmt"
	"os"
	"sort"
)

func guardar(save []string, nArchivo string) {

	f, err := os.Create(nArchivo)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for _, v := range save {

		f.WriteString(v)
		f.WriteString("\n")

	}

}

func sortAscendente(ordenar []string) {

	sort.Strings(ordenar)
	guardar(ordenar, "asecendente.txt")
}

func sortDescendente(ordenar []string) {

	sort.Sort(sort.Reverse(sort.StringSlice(ordenar)))
	guardar(ordenar, "descendente.txt")
}

func main() {

	var str string
	var s []string

	for true {

		fmt.Print("Dame un string (o * para terminar): ")
		fmt.Scanln(&str)

		if str != "*" {
			s = append(s, str)
		} else {
			break
		}
	}

	sortAscendente(s)
	sortDescendente(s)

}
