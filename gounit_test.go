package gounit

import "testing"

func TestAssertEqual(t *testing.T) {
	t.Run("3と3を比較するとtrueを返す", func(t *testing.T) {
		if AssertEqual(3, 3) == false {
			t.Error("not equals")
		}
	})

	t.Run("1と2を比較するとfalseを返す", func(t *testing.T) {
		if AssertEqual(1, 2) == false {
			t.Fatal("unexpected equal")
		}
		t.Log("SUCCESS")
	})
}
