grammar Ejercicio1;
/*
S ::= 'int L

L ::= L [ num ]
    | [ num ]
*/

@parser::header {
    import "Goldang/Node"
}

@parser::members{
    func mostrar(out string){
        fmt.Println(out)
    }
}

//Rules
start returns [node.Nodo n]
    :   t='int' un=lista EOF {
                        mostrar($un.n.Cad + "integer" + $un.n.Aux)
    }
    ;
lista returns [node.Nodo n]
    :   na=lista CORI num=NUMBER CORD {
                                        cad := $na.n.Cad + "arreglo(" + $num.text + ","
                                        aux := $na.n.Aux + ")"
                                        $n = node.NewNodo(cad, aux)
    }
    |   CORI num=NUMBER CORD        {
                                    cad := "arreglo(" + $num.text + ","
                                    aux := ")" 
                                    $n = node.NewNodo(cad,aux)
    }
    ;

//Tokens
CORI: '[';
CORD: ']';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
