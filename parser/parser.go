package parser

import "github.com/JoeCortopassi/go_interpreter/ast"
import "github.com/JoeCortopassi/go_interpreter/lexer"
import "github.com/JoeCortopassi/go_interpreter/token"

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// read two tokens so current and peek token are set on init
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

/***
	pseudocode example
***/
func parseProgram() {
	program = newProgramASTNode()

	advanceTokens()

	for currentToken() != EOF_TOKEN {
		statement = null

		if currentToken() == LET_TOKEN {
			statement = parseLetStatement()
		} else if currentToken() == RETURN_TOKEN {
			statement = parseReturnStatement()
		} else if currentToken() == IF_TOKEN {
			statement = parseIfStatement()
		}

		if statement != null {
			program.Statements.push(statement)
		}

		advanceTokens()
	}

	return program

}

func parseLetStatement() {
	advanceTokens()

	identifier = parseIdentifier()

	advanceTokens()

	if currentToken() != EQUAL_TOKEN {
		parseError("no equal sign!")
		return null
	}

	advanceTokens()

	value = parseExpressions()

	variableStatement = newVariableStatementASTNode()
	variableStatement.identifier = identifier
	variableStatement.value = value

	return variableStatement
}

func parseIdentifier() {
	identifier = newIdentifierASTNode()
	identifier.token = currentToken()
	return identifier
}

func parseExpression() {
	if currentToken() == INTEGER_TOKEN {
		if nextToken() == PLUS_TOKEN {
			return parseOperatorExpression()
		} else if nextToken() == SEMICOLON_TOKEN {
			return parseIntegerLiteral()
		}
	} else if currentToken() == LEFT_PAREN {
		return parseGroupedExpression()
	}
}

func parseOperatorExpression() {
	operatorExpression = newOperatorExpression()
	operatorExpression.left = parseIntegerLiteral()
	operatorExpression.operator = currentToken()
	operatorExpression.right = parseExpression()

	return operatorExpression()
}
