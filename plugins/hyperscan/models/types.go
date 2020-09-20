package models

import (
	"sync"

	"github.com/flier/gohs/hyperscan"
)

// with sync for resource lock
type Scratch struct {
	sync.RWMutex
	s *hyperscan.Scratch
}

type Response struct {
	Errno int         `json:errno`
	Msg   string      `json:msg`
	Data  interface{} `json:data`
}

type MatchResp struct {
	Id         int       `json:id`
	From       int       `json:from`
	To         int       `json:to`
	Flags      int       `json:flags`
	Context    string    `json:context`
	RegexLinev RegexLine `json:regexline`
}

type RegexLine struct {
	Expr string
	Data string
}
