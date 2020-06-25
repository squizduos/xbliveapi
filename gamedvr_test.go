package xbliveapi

import "testing"

func TestGameDVRs(t *testing.T) {
	c := testClient(t)
	gameDVRs, err := c.GameDVRs(testUserXID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", gameDVRs[0])
}
