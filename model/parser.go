package model

import (
	"regexp"
	"strings"
)

type Address struct {
	ZipCode string
	State   string
	City    string
	Street  string
}

type AddressParser interface {
	Parser(addr string) *Address
}

func NewSimpleAddressParser() AddressParser {
	return &simpleAddressParser{}
}

type simpleAddressParser struct {
}

var (
	stateReg = regexp.MustCompile(`.+?(市|縣|島)`)
	cityReg  = regexp.MustCompile(`.+?(區|鎮|鄉|市|山|門|沙|嶼)`)
)

func (p *simpleAddressParser) Parser(addr string) *Address {

	pArr := stateReg.FindAllString(addr, 1)

	if len(pArr) != 1 {
		return &Address{
			Street: addr,
		}
	}

	addr = strings.Replace(addr, pArr[0], "", 1)

	rArr := cityReg.FindAllString(addr, -1)

	if len(rArr) != 1 {
		return &Address{
			State:  pArr[0],
			Street: addr,
		}
	}
	addr = strings.Replace(addr, rArr[0], "", 1)
	code, ok := getCode(pArr[0], rArr[0])
	if !ok {
		return &Address{
			State:  pArr[0],
			Street: rArr[0],
		}
	}

	return &Address{
		ZipCode: code,
		State:   pArr[0],
		City:    rArr[0],
		Street:  addr,
	}
}
