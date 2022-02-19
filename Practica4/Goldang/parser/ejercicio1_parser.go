// Code generated from Ejercicio1.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Ejercicio1

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Goldang/Node"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 7, 28, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 23, 10, 3, 12, 3, 14,
	3, 26, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 2, 2, 26, 2, 6, 3, 2, 2, 2, 4,
	11, 3, 2, 2, 2, 6, 7, 7, 3, 2, 2, 7, 8, 5, 4, 3, 2, 8, 9, 7, 2, 2, 3, 9,
	10, 8, 2, 1, 2, 10, 3, 3, 2, 2, 2, 11, 12, 8, 3, 1, 2, 12, 13, 7, 4, 2,
	2, 13, 14, 7, 6, 2, 2, 14, 15, 7, 5, 2, 2, 15, 16, 8, 3, 1, 2, 16, 24,
	3, 2, 2, 2, 17, 18, 12, 4, 2, 2, 18, 19, 7, 4, 2, 2, 19, 20, 7, 6, 2, 2,
	20, 21, 7, 5, 2, 2, 21, 23, 8, 3, 1, 2, 22, 17, 3, 2, 2, 2, 23, 26, 3,
	2, 2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 5, 3, 2, 2, 2, 26,
	24, 3, 2, 2, 2, 3, 24,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'int'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "CORI", "CORD", "NUMBER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "lista",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Ejercicio1Parser struct {
	*antlr.BaseParser
}

func NewEjercicio1Parser(input antlr.TokenStream) *Ejercicio1Parser {
	this := new(Ejercicio1Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ejercicio1.g4"

	return this
}

func mostrar(out string) {
	fmt.Println(out)
}

// Ejercicio1Parser tokens.
const (
	Ejercicio1ParserEOF        = antlr.TokenEOF
	Ejercicio1ParserT__0       = 1
	Ejercicio1ParserCORI       = 2
	Ejercicio1ParserCORD       = 3
	Ejercicio1ParserNUMBER     = 4
	Ejercicio1ParserWHITESPACE = 5
)

// Ejercicio1Parser rules.
const (
	Ejercicio1ParserRULE_start = 0
	Ejercicio1ParserRULE_lista = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t token.
	GetT() antlr.Token

	// SetT sets the t token.
	SetT(antlr.Token)

	// GetUn returns the un rule contexts.
	GetUn() IListaContext

	// SetUn sets the un rule contexts.
	SetUn(IListaContext)

	// GetN returns the n attribute.
	GetN() node.Nodo

	// SetN sets the n attribute.
	SetN(node.Nodo)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      node.Nodo
	t      antlr.Token
	un     IListaContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Ejercicio1ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Ejercicio1ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetT() antlr.Token { return s.t }

func (s *StartContext) SetT(v antlr.Token) { s.t = v }

func (s *StartContext) GetUn() IListaContext { return s.un }

func (s *StartContext) SetUn(v IListaContext) { s.un = v }

func (s *StartContext) GetN() node.Nodo { return s.n }

func (s *StartContext) SetN(v node.Nodo) { s.n = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserEOF, 0)
}

func (s *StartContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Ejercicio1Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Ejercicio1ParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)

		var _m = p.Match(Ejercicio1ParserT__0)

		localctx.(*StartContext).t = _m
	}
	{
		p.SetState(5)

		var _x = p.lista(0)

		localctx.(*StartContext).un = _x
	}
	{
		p.SetState(6)
		p.Match(Ejercicio1ParserEOF)
	}

	mostrar(localctx.(*StartContext).GetUn().GetN().Cad + "integer" + localctx.(*StartContext).GetUn().GetN().Aux)

	return localctx
}

// IListaContext is an interface to support dynamic dispatch.
type IListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// GetNa returns the na rule contexts.
	GetNa() IListaContext

	// SetNa sets the na rule contexts.
	SetNa(IListaContext)

	// GetN returns the n attribute.
	GetN() node.Nodo

	// SetN sets the n attribute.
	SetN(node.Nodo)

	// IsListaContext differentiates from other interfaces.
	IsListaContext()
}

type ListaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      node.Nodo
	na     IListaContext
	num    antlr.Token
}

func NewEmptyListaContext() *ListaContext {
	var p = new(ListaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Ejercicio1ParserRULE_lista
	return p
}

func (*ListaContext) IsListaContext() {}

func NewListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaContext {
	var p = new(ListaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Ejercicio1ParserRULE_lista

	return p
}

func (s *ListaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaContext) GetNum() antlr.Token { return s.num }

func (s *ListaContext) SetNum(v antlr.Token) { s.num = v }

func (s *ListaContext) GetNa() IListaContext { return s.na }

func (s *ListaContext) SetNa(v IListaContext) { s.na = v }

func (s *ListaContext) GetN() node.Nodo { return s.n }

func (s *ListaContext) SetN(v node.Nodo) { s.n = v }

func (s *ListaContext) CORI() antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserCORI, 0)
}

func (s *ListaContext) CORD() antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserCORD, 0)
}

func (s *ListaContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(Ejercicio1ParserNUMBER, 0)
}

func (s *ListaContext) Lista() IListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaContext)
}

func (s *ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.EnterLista(s)
	}
}

func (s *ListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Ejercicio1Listener); ok {
		listenerT.ExitLista(s)
	}
}

func (p *Ejercicio1Parser) Lista() (localctx IListaContext) {
	return p.lista(0)
}

func (p *Ejercicio1Parser) lista(_p int) (localctx IListaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, Ejercicio1ParserRULE_lista, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.Match(Ejercicio1ParserCORI)
	}
	{
		p.SetState(11)

		var _m = p.Match(Ejercicio1ParserNUMBER)

		localctx.(*ListaContext).num = _m
	}
	{
		p.SetState(12)
		p.Match(Ejercicio1ParserCORD)
	}

	cad := "arreglo(" + (func() string {
		if localctx.(*ListaContext).GetNum() == nil {
			return ""
		} else {
			return localctx.(*ListaContext).GetNum().GetText()
		}
	}()) + ","
	aux := ")"
	localctx.(*ListaContext).n = node.NewNodo(cad, aux)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaContext(p, _parentctx, _parentState)
			localctx.(*ListaContext).na = _prevctx
			p.PushNewRecursionContext(localctx, _startState, Ejercicio1ParserRULE_lista)
			p.SetState(15)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(16)
				p.Match(Ejercicio1ParserCORI)
			}
			{
				p.SetState(17)

				var _m = p.Match(Ejercicio1ParserNUMBER)

				localctx.(*ListaContext).num = _m
			}
			{
				p.SetState(18)
				p.Match(Ejercicio1ParserCORD)
			}

			cad := localctx.(*ListaContext).GetNa().GetN().Cad + "arreglo(" + (func() string {
				if localctx.(*ListaContext).GetNum() == nil {
					return ""
				} else {
					return localctx.(*ListaContext).GetNum().GetText()
				}
			}()) + ","
			aux := localctx.(*ListaContext).GetNa().GetN().Aux + ")"
			localctx.(*ListaContext).n = node.NewNodo(cad, aux)

		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Ejercicio1Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ListaContext = nil
		if localctx != nil {
			t = localctx.(*ListaContext)
		}
		return p.Lista_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Ejercicio1Parser) Lista_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
