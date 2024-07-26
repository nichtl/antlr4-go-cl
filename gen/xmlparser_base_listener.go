// Code generated from /Users/xujian8/GolandProjects/go-antlr4-cl/XMLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // XMLParser

import "github.com/antlr4-go/antlr/v4"

// BaseXMLParserListener is a complete listener for a parse tree produced by XMLParser.
type BaseXMLParserListener struct{}

var _ XMLParserListener = &BaseXMLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXMLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXMLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXMLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXMLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseXMLParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseXMLParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterElement is called when production element is entered.
func (s *BaseXMLParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseXMLParserListener) ExitElement(ctx *ElementContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseXMLParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseXMLParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterContent is called when production content is entered.
func (s *BaseXMLParserListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseXMLParserListener) ExitContent(ctx *ContentContext) {}
