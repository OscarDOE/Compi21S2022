// Code generated from Ejercicio1.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 33, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 23, 10, 5, 13, 5, 14,
	5, 24, 3, 6, 6, 6, 28, 10, 6, 13, 6, 14, 6, 29, 3, 6, 3, 6, 2, 2, 7, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 34, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2,
	7, 19, 3, 2, 2, 2, 9, 22, 3, 2, 2, 2, 11, 27, 3, 2, 2, 2, 13, 14, 7, 107,
	2, 2, 14, 15, 7, 112, 2, 2, 15, 16, 7, 118, 2, 2, 16, 4, 3, 2, 2, 2, 17,
	18, 7, 93, 2, 2, 18, 6, 3, 2, 2, 2, 19, 20, 7, 95, 2, 2, 20, 8, 3, 2, 2,
	2, 21, 23, 9, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 22,
	3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 10, 3, 2, 2, 2, 26, 28, 9, 3, 2, 2,
	27, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3,
	2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 8, 6, 2, 2, 32, 12, 3, 2, 2, 2, 5,
	2, 24, 29, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'int'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "CORI", "CORD", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "CORI", "CORD", "NUMBER", "WHITESPACE",
}

type Ejercicio1Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEjercicio1Lexer(input antlr.CharStream) *Ejercicio1Lexer {

	l := new(Ejercicio1Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Ejercicio1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Ejercicio1Lexer tokens.
const (
	Ejercicio1LexerT__0       = 1
	Ejercicio1LexerCORI       = 2
	Ejercicio1LexerCORD       = 3
	Ejercicio1LexerNUMBER     = 4
	Ejercicio1LexerWHITESPACE = 5
)
