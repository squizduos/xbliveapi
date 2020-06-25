package xbliveapi

import "testing"

func TestTitles(t *testing.T) {
	c := testClient(t)
	titles, err := c.Titles(testUserXID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", titles[0])
}
