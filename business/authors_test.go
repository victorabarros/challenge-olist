package business

import "testing"

func testDeduplucation(t *testing.T) {
	cases := []struct {
		arr  []string
		want []string
		ok   bool
	}{
		{
			arr:  []string{"a", "a", "b", "c"},
			want: []string{"a", "b", "c"},
			ok:   true,
		},
		{
			arr:  []string{"a", "a", "a"},
			want: []string{"a"},
			ok:   true,
		},
		{
			arr:  []string{"a", "a", "b", "c"},
			want: []string{"a"},
			ok:   false,
		},
	}

	for _, c := range cases {
		ans := deduplicate(c.arr)
		if (assertSlices(ans, c.want)) != c.ok {
			t.Errorf("Deduplicate of %s must be %s", c.arr, c.want)
		}
	}
}

func testReadCsv(t *testing.T) {
	// TODO check if authors.csv exists
	want := []string{
		"Luciano Ramalho",
		"Osvaldo Santana Neto",
		"David Beazley",
		"Chetan Giridhar",
		"Brian K. Jones",
		"novo",
	}

	ans := readNames("Authors.csv")
	if !assertSlices(want, ans) {
		t.Errorf("Csv differents")
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
