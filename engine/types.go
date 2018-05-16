package engine

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type ParseFunc func(data []byte) ParseResult

type Request struct {
	Url       string
	ParseFunc ParseFunc
	FetchFunc func(string) ([]byte, error)
}

type Item struct {
	Id       string
	Url      string
	Type     string
	Payload  interface{}
	Action   string
	ParentId string
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
