package main

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	BraceOpen TokenType = iota
	BraceClose
	BracketOpen
	BracketClose
	String
	Number
	Comma
	Colon
	True
	False
	Null
)

type Token struct {
	tokenType TokenType
	value     string
}

func (T TokenType) String() string {
	return [...]string{"BraceOpen", "BraceClose", "BracketOpen", "BracketClose", "String", "Number", "Comma", "Colon", "True", "False", "Null"}[T]
}

func Tokenizer(inputString string) []Token {
	var Tokens []Token
	for i := 0; i < len(inputString); i++ {

		switch string(inputString[i]) {
		case "\\s":
			continue
		case "[":
			Tokens = append(Tokens, Token{tokenType: BraceOpen, value: string(inputString[i])})
		case "]":
			Tokens = append(Tokens, Token{tokenType: BraceClose, value: string(inputString[i])})
		case "{":
			Tokens = append(Tokens, Token{tokenType: BracketOpen, value: string(inputString[i])})
		case "}":
			Tokens = append(Tokens, Token{tokenType: BracketClose, value: string(inputString[i])})
		case ",":
			Tokens = append(Tokens, Token{tokenType: Comma, value: string(inputString[i])})
		case ":":
			Tokens = append(Tokens, Token{tokenType: Colon, value: string(inputString[i])})
		case "\"":
			{
				j := i + 1
				var value []byte
				value = append(value, inputString[j])
				for string(inputString[j]) != "\"" && j < (len(inputString)-1) {
					j++
					if string(inputString[j]) != "\"" {
						value = append(value, inputString[j])
					}

				}
				Tokens = append(Tokens, Token{tokenType: String, value: string(value)})
				i = j
			}

			//  update to  this
			//case isNumber(value):
			//				tokenType = Number
			//			case isBooleanTrue(value):
			//				tokenType = True
			//			case isBooleanFalse(value):
			//				tokenType = False
			//			case isNull(value):
			//				tokenType = Null
		case "T":
			{
				Tokens = append(Tokens, Token{tokenType: True, value: "T"})
				i = i + 2
			}
		case "F":
			{
				Tokens = append(Tokens, Token{tokenType: False, value: "F"})
				i = i + 3
			}
		case "N":
			{
				Tokens = append(Tokens, Token{tokenType: Null, value: "N"})
				i = i + 2
			}
		default:
			if unicode.IsNumber(rune(inputString[i])) {
				Tokens = append(Tokens, Token{tokenType: Number, value: string(inputString[i])})

			} else {
				fmt.Println("UNKNOWN CHAR")
			}

		}
	}
	return Tokens
}
