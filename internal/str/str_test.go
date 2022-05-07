package str

import (
	"testing"
)

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	s := UnderscoreToUpperCamelCase("user_name")
	t.Log(s)
	if s != "UserName" {
		t.Error()
	}
}
