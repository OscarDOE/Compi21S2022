// Code generated from Ejercicio1.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Ejercicio1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Ejercicio1Listener is a complete listener for a parse tree produced by Ejercicio1Parser.
type Ejercicio1Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterLista is called when entering the lista production.
	EnterLista(c *ListaContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitLista is called when exiting the lista production.
	ExitLista(c *ListaContext)
}
