package api

import "fmt"

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func (q *Quote) String() string {
	return fmt.Sprintf("%s\n(c) %s\n", q.Quote, q.Author)
}

type Quotes []Quote
