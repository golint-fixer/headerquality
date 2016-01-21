package headerquality

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FactorTest struct {
	header   string
	factor   string
	expected Parameters
}

var FactorTests = []FactorTest{
	// Test an empty header
	{"", "Accept-Language", make(Parameters, 0)},

	// Test a single unqualified header value
	{"en-gb", "Accept-Language", Parameters{Parameter{"en-gb", 1}}},

	// Test a single qualified header value
	{"en-gb;q=0.8", "Accept-Language", Parameters{Parameter{"en-gb", 0.8}}},

	// Test multiple unqualified header values
	{"en-gb, nl,en-us", "Accept-Language", Parameters{
		Parameter{"en-gb", 1}, Parameter{"nl", 1}, Parameter{"en-us", 1},
	}},

	// Test multiple qualified header values
	{"en-gb;q=0.2, nl;q=1,en-us;q=0.5", "Accept-Language", Parameters{
		Parameter{"nl", 1}, Parameter{"en-us", 0.5}, Parameter{"en-gb", 0.2},
	}},

	{"en-gb;q=0.2 , nl;q=1, en-us;q=0.5    ", "Accept-Language", Parameters{
		Parameter{"nl", 1}, Parameter{"en-us", 0.5}, Parameter{"en-gb", 0.2},
	}},
}

func TestAcceptFactorTests(t *testing.T) {
	for _, test := range FactorTests {
		request, _ := http.NewRequest("GET", "", nil)
		request.Header.Add(test.factor, test.header)

		got := Factors(test.factor, request)

		assert.Equal(t, test.expected, got)
	}
}
