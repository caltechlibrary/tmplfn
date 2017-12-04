package tmplfn

import (
	"testing"
)

func TestSlugUnslug(t *testing.T) {
	s := "The Jumbles"
	expected := "the-jumbles"
	r := slug(s)
	if expected != r {
		t.Errorf("expected %q, got %q", expected, r)
	}

	s = "red/blue"
	expected = "red~blue"
	r = slug(s)
	if expected != r {
		t.Errorf("expected %q, got %q", expected, r)
	}
}
