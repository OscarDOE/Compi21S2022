grammar Practica4;
/*
S ::= ( num , num ) L

L ::= L ( num , num ) 
    | ( num , num ) 
*/

@parser::header {
    import ("Goldang/Node"

	"strings"
    "math"
    )
}

@parser::members{



    func mostrar(out string){


        b := strings.Split(out, " ")


	var s []float64
	for i := 0; i < len(b); i++ {

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

		i = i + 1
		if i >= (len(s) - 1) {
			break
		}
	}
        fmt.Println(out)
	fmt.Println("la respuesta es: ", answer)

    }
    
}

//Rules
start returns [node.Nodo n]
    :   un=lista EOF {
                        mostrar($un.n.A + $un.n.B )
    }
    ;
lista returns [node.Nodo n]
    :   na=lista PARA num=NUMBER COMA num2=NUMBER PARC {
                                        x := $na.n.A + " " + $num.text + " "+$num2.text
                                        y := $na.n.B 
                                        $n = node.NewNodo(x, y)
    }
    |   PARA num=NUMBER COMA num2=NUMBER PARC        {
                                    x := $num.text + " " + $num2.text
                                    y := "" 
                                    $n = node.NewNodo(x,y)
    }
    ;

//Tokens
PARA: '(';
PARC: ')';
COMA: ',';
NUMBER: ('-')?[0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
