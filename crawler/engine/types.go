package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
	Content []interface{}
}

type ParseArticleResult struct {
	Content []interface{}
	Title []interface{}
	chapter []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
