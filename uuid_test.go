package utils

import "testing"

func TestUUID(t *testing.T) {
	var id UUID = Rand()
	t.Log(id.Hex())
	t.Log(id.Hex())
	t.Log(id.Hex())

	id1, _ := FromStr("1870747d-b26c-4507-9518-1ca62bc66e5d")
	id2 := MustFromStr("1870747db26c450795181ca62bc66e5d")
	t.Log(id1 == id2) // true
	t.Log("succeed")
}
