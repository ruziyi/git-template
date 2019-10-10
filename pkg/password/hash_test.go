package password

import "testing"

func TestPasswordHash(t *testing.T) {
	str := "123"
	h, err := PasswordHash(str)
	if err != nil {
		t.Fatal()
	}
	if !CheckHash(str, h) {
		t.Log(h, str)
		t.Fatal()
	}
}
