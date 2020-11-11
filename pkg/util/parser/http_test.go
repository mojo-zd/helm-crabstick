package parser

import (
	"testing"
)

func TestParseUrl(t *testing.T) {
	url, err := ParseURL("http://baidu.com/xxx?test=ll")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url)
}
