/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package practica1;

import java.io.FileInputStream;

/**
 *
 * @author elmco
 */
public class Practica1 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        interpretar("entrada.txt");
        

    }
    private static void interpretar(String path) {
        analizadores.Parser parse;
        try {
            parse = new analizadores.Parser(new analizadores.Lexico(new FileInputStream(path)));
            parse.parse();
        } catch(Exception e) {
            System.out.println("Error fatal en compilación de entrada.");
            System.out.println("Causa: "+e.getCause());
        }
    }
}
