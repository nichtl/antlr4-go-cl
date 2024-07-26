// Code generated from /Users/xujian8/GolandProjects/go-antlr4-cl/XMLParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // XMLParser

import "github.com/antlr4-go/antlr/v4"

type BaseXMLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseXMLParserVisitor) VisitDocument(ctx *DocumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseXMLParserVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}
