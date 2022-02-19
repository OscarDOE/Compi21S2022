grammar Ejercicio1;
/*
S ::= ( num , num ) L

L ::= L ( num , num ) 
    | ( num , num ) 
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
    :   un=lista EOF {
                        mostrar($un.n.A + "integer" + $un.n.B )
    }
    ;
lista returns [node.Nodo n]
    :   na=lista PARA num=NUMBER COMA num2=NUMBER PARC {
                                        x1 := $na.n.A + "," + $num.text + ","
                                        x2 := $na.n.B + ")"
                                        $n = node.NewNodo(x, y)
    }
    |   PARA num=NUMBER COMA num2=NUMBER PARC        {
                                    x := "(ultimo" + $num.text + ","
                                    y := "ultimop)" 
                                    $n = node.NewNodo(x,y)
    }
    ;

//Tokens
PARA: '(';
PARC: ')';
COMA: ',';
NUMBER: [0-9]+;
MENOS: '-';
WHITESPACE: [ \r\n\t]+ -> skip;
