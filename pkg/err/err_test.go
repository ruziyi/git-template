package err

import "testing"

func TestErrorList_Error(t *testing.T) {
	var el ErrorList
	if el != nil {
		t.Fatal(el)
	}
}
