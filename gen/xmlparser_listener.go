// Code generated from /Users/xujian8/GolandProjects/go-antlr4-cl/XMLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // XMLParser

import "github.com/antlr4-go/antlr/v4"

// XMLParserListener is a complete listener for a parse tree produced by XMLParser.
type XMLParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)
}
