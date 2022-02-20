package main

import (
	parser "Goldang/parserx"
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Syntax struct {
	*parser.BasePractica4Listener
}

func NewSyntax() *Syntax {
	return new(Syntax)
}
func interpretar(in string) {
	is := antlr.NewInputStream(in)

	//crear analizador Léxico
	lexer := parser.NewPractica4Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//crear analizador sintáctico
	p := parser.NewPractica4Parser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewSyntax(), p.Start())

}
func main() {
	fmt.Println("HOla")
	file, err := os.Open("entrada2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	text := ""

	for fileScanner.Scan() {
		text += fileScanner.Text()
	}

	interpretar(text)
	fmt.Println("FINITE INCATATEM")

	/*cadena := "3 -3 4 7 -6 2 -1 2 -3 -5 "

	b := strings.Split(cadena, " ")
	fmt.Println(b)

	var s []float64
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
		w, _ := strconv.ParseFloat(b[i], 64)

		s = append(s, w)
	}
	var answer float64 = 0
	for i := 0; i < len(s)-1; i++ {
		if i+2 >= (len(s) - 1) {
			break
		}

		r1 := math.Pow(s[i+2]-s[i], 2)
		r2 := math.Pow(s[i+3]-s[i+1], 2)
		answer = answer + math.Sqrt(r1+r2)
		fmt.Println("anser", answer)
		i = i + 1
		if i >= (len(s) - 1) {
			break
		}
	}
	fmt.Println("la respuesta es: ", answer)*/
}
