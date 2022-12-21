// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
  "sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type LookMLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lookmllexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  channelNames           []string
  modeNames              []string
  literalNames           []string
  symbolicNames          []string
  ruleNames              []string
  predictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func lookmllexerLexerInit() {
  staticData := &lookmllexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.literalNames = []string{
    "", "'model:'", "'{'", "'}'", "'view:'", "'dimension:'", "'measure:'", 
    "'field:'", "'type:'", "'sql:'", "'description:'", "'group_label:'", 
    "'string'", "'number'", "'boolean'", "'date'", "'datetime'", "'time'", 
    "'duration'",
  }
  staticData.symbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "ID", "STRING", "WS",
  }
  staticData.ruleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "T__17", "ID", "STRING", "WS",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 21, 197, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 
	7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 
	9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 
	1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 
	1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 
	13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 
	177, 8, 18, 10, 18, 12, 18, 180, 9, 18, 1, 19, 1, 19, 5, 19, 184, 8, 19, 
	10, 19, 12, 19, 187, 9, 19, 1, 19, 1, 19, 1, 20, 4, 20, 192, 8, 20, 11, 
	20, 12, 20, 193, 1, 20, 1, 20, 1, 185, 0, 21, 1, 1, 3, 2, 5, 3, 7, 4, 9, 
	5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 
	29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 1, 0, 3, 3, 0, 
	65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 
	10, 13, 13, 32, 32, 199, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 
	0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 
	0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 
	1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 
	29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 
	0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 
	0, 3, 50, 1, 0, 0, 0, 5, 52, 1, 0, 0, 0, 7, 54, 1, 0, 0, 0, 9, 60, 1, 0, 
	0, 0, 11, 71, 1, 0, 0, 0, 13, 80, 1, 0, 0, 0, 15, 87, 1, 0, 0, 0, 17, 93, 
	1, 0, 0, 0, 19, 98, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 124, 1, 0, 0, 
	0, 25, 131, 1, 0, 0, 0, 27, 138, 1, 0, 0, 0, 29, 146, 1, 0, 0, 0, 31, 151, 
	1, 0, 0, 0, 33, 160, 1, 0, 0, 0, 35, 165, 1, 0, 0, 0, 37, 174, 1, 0, 0, 
	0, 39, 181, 1, 0, 0, 0, 41, 191, 1, 0, 0, 0, 43, 44, 5, 109, 0, 0, 44, 
	45, 5, 111, 0, 0, 45, 46, 5, 100, 0, 0, 46, 47, 5, 101, 0, 0, 47, 48, 5, 
	108, 0, 0, 48, 49, 5, 58, 0, 0, 49, 2, 1, 0, 0, 0, 50, 51, 5, 123, 0, 0, 
	51, 4, 1, 0, 0, 0, 52, 53, 5, 125, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 
	118, 0, 0, 55, 56, 5, 105, 0, 0, 56, 57, 5, 101, 0, 0, 57, 58, 5, 119, 
	0, 0, 58, 59, 5, 58, 0, 0, 59, 8, 1, 0, 0, 0, 60, 61, 5, 100, 0, 0, 61, 
	62, 5, 105, 0, 0, 62, 63, 5, 109, 0, 0, 63, 64, 5, 101, 0, 0, 64, 65, 5, 
	110, 0, 0, 65, 66, 5, 115, 0, 0, 66, 67, 5, 105, 0, 0, 67, 68, 5, 111, 
	0, 0, 68, 69, 5, 110, 0, 0, 69, 70, 5, 58, 0, 0, 70, 10, 1, 0, 0, 0, 71, 
	72, 5, 109, 0, 0, 72, 73, 5, 101, 0, 0, 73, 74, 5, 97, 0, 0, 74, 75, 5, 
	115, 0, 0, 75, 76, 5, 117, 0, 0, 76, 77, 5, 114, 0, 0, 77, 78, 5, 101, 
	0, 0, 78, 79, 5, 58, 0, 0, 79, 12, 1, 0, 0, 0, 80, 81, 5, 102, 0, 0, 81, 
	82, 5, 105, 0, 0, 82, 83, 5, 101, 0, 0, 83, 84, 5, 108, 0, 0, 84, 85, 5, 
	100, 0, 0, 85, 86, 5, 58, 0, 0, 86, 14, 1, 0, 0, 0, 87, 88, 5, 116, 0, 
	0, 88, 89, 5, 121, 0, 0, 89, 90, 5, 112, 0, 0, 90, 91, 5, 101, 0, 0, 91, 
	92, 5, 58, 0, 0, 92, 16, 1, 0, 0, 0, 93, 94, 5, 115, 0, 0, 94, 95, 5, 113, 
	0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 58, 0, 0, 97, 18, 1, 0, 0, 0, 98, 
	99, 5, 100, 0, 0, 99, 100, 5, 101, 0, 0, 100, 101, 5, 115, 0, 0, 101, 102, 
	5, 99, 0, 0, 102, 103, 5, 114, 0, 0, 103, 104, 5, 105, 0, 0, 104, 105, 
	5, 112, 0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5, 105, 0, 0, 107, 108, 
	5, 111, 0, 0, 108, 109, 5, 110, 0, 0, 109, 110, 5, 58, 0, 0, 110, 20, 1, 
	0, 0, 0, 111, 112, 5, 103, 0, 0, 112, 113, 5, 114, 0, 0, 113, 114, 5, 111, 
	0, 0, 114, 115, 5, 117, 0, 0, 115, 116, 5, 112, 0, 0, 116, 117, 5, 95, 
	0, 0, 117, 118, 5, 108, 0, 0, 118, 119, 5, 97, 0, 0, 119, 120, 5, 98, 0, 
	0, 120, 121, 5, 101, 0, 0, 121, 122, 5, 108, 0, 0, 122, 123, 5, 58, 0, 
	0, 123, 22, 1, 0, 0, 0, 124, 125, 5, 115, 0, 0, 125, 126, 5, 116, 0, 0, 
	126, 127, 5, 114, 0, 0, 127, 128, 5, 105, 0, 0, 128, 129, 5, 110, 0, 0, 
	129, 130, 5, 103, 0, 0, 130, 24, 1, 0, 0, 0, 131, 132, 5, 110, 0, 0, 132, 
	133, 5, 117, 0, 0, 133, 134, 5, 109, 0, 0, 134, 135, 5, 98, 0, 0, 135, 
	136, 5, 101, 0, 0, 136, 137, 5, 114, 0, 0, 137, 26, 1, 0, 0, 0, 138, 139, 
	5, 98, 0, 0, 139, 140, 5, 111, 0, 0, 140, 141, 5, 111, 0, 0, 141, 142, 
	5, 108, 0, 0, 142, 143, 5, 101, 0, 0, 143, 144, 5, 97, 0, 0, 144, 145, 
	5, 110, 0, 0, 145, 28, 1, 0, 0, 0, 146, 147, 5, 100, 0, 0, 147, 148, 5, 
	97, 0, 0, 148, 149, 5, 116, 0, 0, 149, 150, 5, 101, 0, 0, 150, 30, 1, 0, 
	0, 0, 151, 152, 5, 100, 0, 0, 152, 153, 5, 97, 0, 0, 153, 154, 5, 116, 
	0, 0, 154, 155, 5, 101, 0, 0, 155, 156, 5, 116, 0, 0, 156, 157, 5, 105, 
	0, 0, 157, 158, 5, 109, 0, 0, 158, 159, 5, 101, 0, 0, 159, 32, 1, 0, 0, 
	0, 160, 161, 5, 116, 0, 0, 161, 162, 5, 105, 0, 0, 162, 163, 5, 109, 0, 
	0, 163, 164, 5, 101, 0, 0, 164, 34, 1, 0, 0, 0, 165, 166, 5, 100, 0, 0, 
	166, 167, 5, 117, 0, 0, 167, 168, 5, 114, 0, 0, 168, 169, 5, 97, 0, 0, 
	169, 170, 5, 116, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 111, 0, 0, 
	172, 173, 5, 110, 0, 0, 173, 36, 1, 0, 0, 0, 174, 178, 7, 0, 0, 0, 175, 
	177, 7, 1, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 
	1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 38, 1, 0, 0, 0, 180, 178, 1, 0, 
	0, 0, 181, 185, 5, 34, 0, 0, 182, 184, 9, 0, 0, 0, 183, 182, 1, 0, 0, 0, 
	184, 187, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 
	188, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 34, 0, 0, 189, 40, 
	1, 0, 0, 0, 190, 192, 7, 2, 0, 0, 191, 190, 1, 0, 0, 0, 192, 193, 1, 0, 
	0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 
	195, 196, 6, 20, 0, 0, 196, 42, 1, 0, 0, 0, 4, 0, 178, 185, 193, 1, 6, 
	0, 0,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// LookMLLexerInit initializes any static state used to implement LookMLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLookMLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LookMLLexerInit() {
  staticData := &lookmllexerLexerStaticData
  staticData.once.Do(lookmllexerLexerInit)
}

// NewLookMLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLookMLLexer(input antlr.CharStream) *LookMLLexer {
  LookMLLexerInit()
	l := new(LookMLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &lookmllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "LookML.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LookMLLexer tokens.
const (
	LookMLLexerT__0 = 1
	LookMLLexerT__1 = 2
	LookMLLexerT__2 = 3
	LookMLLexerT__3 = 4
	LookMLLexerT__4 = 5
	LookMLLexerT__5 = 6
	LookMLLexerT__6 = 7
	LookMLLexerT__7 = 8
	LookMLLexerT__8 = 9
	LookMLLexerT__9 = 10
	LookMLLexerT__10 = 11
	LookMLLexerT__11 = 12
	LookMLLexerT__12 = 13
	LookMLLexerT__13 = 14
	LookMLLexerT__14 = 15
	LookMLLexerT__15 = 16
	LookMLLexerT__16 = 17
	LookMLLexerT__17 = 18
	LookMLLexerID = 19
	LookMLLexerSTRING = 20
	LookMLLexerWS = 21
)

