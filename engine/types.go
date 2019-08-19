package engine

import "golang.org/x/net/html"

type Request struct {
	Url string
	ParserFunc func(node *html.Node) ParserResult

}
type ParserResult struct {
	Requests []Request
	Items []interface{}
}