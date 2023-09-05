package gojieba

import (
	"github.com/zhufenghua1998/gojieba/deps/cppjieba"
	"github.com/zhufenghua1998/gojieba/deps/limonp"
	"github.com/zhufenghua1998/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
