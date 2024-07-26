// Code generated from /Users/xujian8/GolandProjects/go-antlr4-cl/XMLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // XMLParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by XMLParser.
type XMLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by XMLParser#document.
	VisitDocument(ctx *DocumentContext) interface{}

	// Visit a parse tree produced by XMLParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by XMLParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by XMLParser#content.
	VisitContent(ctx *ContentContext) interface{}
}
