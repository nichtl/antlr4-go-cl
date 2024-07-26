parser grammar XMLParser;

options {tokenVocab = XMLLexer;}

document : element* EOF;

element : OPEN Name attribute* (CLOSE  content* OPEN SlashName CLOSE | SLASH_CLOSE) ;

attribute : Name EQUALS STRING ;

content : element | TEXT | EntityRef ;

