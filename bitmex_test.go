package traderobot

import (
	"github.com/lpisces/traderobot/bitmex"
	"testing"
)

func TestQuotes(t *testing.T) {

	opt := make(map[string]string)
	opt["symbol"] = "XBT"
	opt["reverse"] = "true"
	opt["count"] = "100"
	q, err:= bitmex.Quotes(opt)

	if err != nil {
		t.Fatal(err)
	}

	if len(q) == 0 {
		t.Fatal("no data")
	}

	t.Log(q[0])
}
