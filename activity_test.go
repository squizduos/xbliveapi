package xbliveapi

import "testing"

func TestActivity(t *testing.T) {
	c := testClient(t)
	activities, err := c.ActivityStatuses(testUserXID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", activities[0])
}
