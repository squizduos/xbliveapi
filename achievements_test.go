package xbliveapi

import "testing"

func TestAchievements(t *testing.T) {
	c := testClient(t)
	achievements, err := c.Achievements(testUserXID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", achievements[0])
}
