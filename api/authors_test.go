package api

import (
	"net/http"
	"net/url"
	"testing"
)

func testValidateListParams(t *testing.T) {
	u := url.URL{}
	u.RawQuery = "name=oi&limit=10&offset=0"
	req := http.Request{URL: &u}
	params := req.URL.Query()

	cases := []struct {
		queryParam string
		limit      int
		offset     int
		names      []string
		ok         bool
	}{
		{
			queryParam: "name=oi&limit=10&offset=0",
			limit:      10,
			offset:     0,
			names:      []string{"oi"},
			ok:         true,
		},
	}

	for _, c := range cases {
		req.URL.RawQuery = c.queryParam
		limit, offset, names, err := validateListParams(&req)
		if err != nil {
			t.Errorf("error")
		}
		if (limit != c.limit || offset != c.offset || assertSlices(names, c.names)) == c.ok {
			t.Errorf("error")
		}
	}

}

func assertSlices(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
