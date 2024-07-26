package go_antlr4_cl

import (
	"crypto/tls"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "go-antlr4-cl/gen"
	"log"
	"net"
	"syscall"
	"testing"
)

func Test3(t *testing.T) {
	input := antlr.NewInputStream("a = 5\nb = 5\na+b*2\n(1+3)*2\nclear\n")
	lexer := parser.NewLabeledExprLexer(input)
	token := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewLabeledExprParser(token)
	tree := p.Prog()
	visitor := NewEvalVisitor()
	r := visitor.Visit(tree)
	fmt.Println(r)

}

func Test2(t *testing.T) {

	//is := antlr.NewInputStream("{1,2,{3,4,5},6}")
	//lexer := parser.NewArrayInitLexer(is)
	//stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//p := parser.NewArrayInitParser(stream)
	//var listener parser.BaseArrayInitListener
	//antlr.ParseTreeWalkerDefault.Walk(&listener, p.Init())

}

func Test1(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	log.Println("fd:, err:", fd, err)
	listener, _ := net.Listen("udp", ":1234")
	for {
		fmt.Println("connection established")
		conn, _ := listener.Accept()
		go process(conn)
	}
}

func up_protocol_tls(conn net.Conn) {
	conn = tls.Server(conn, &tls.Config{})
}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("error closing connection", err)
		}
	}(conn)
	fmt.Println("connection established,{}", conn.RemoteAddr())
	buf := make([]byte, 1024)
	_, _ = conn.Read(buf[:])
	_, _ = conn.Write([]byte("epoll serve"))
}
