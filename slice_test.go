package tu

import "testing"

func TestCmpStr(t *testing.T) {
	if CmpStr([]string{"s1"}, []string{"s1", "s2"}) {
		t.Error("ill")
		return
	}
	if CmpStr([]string{"s1", "s1"}, []string{"s1", "s2"}) {
		t.Error("ill")
		return
	}
	if CmpStr([]string{"s1", "s2"}, []string{"s2", "s2"}) {
		t.Error("ill")
		return
	}
	if !CmpStr([]string{"s1", "s2", "s3"}, []string{"s2", "s3", "s1"}) {
		t.Error("ill")
		return
	}
}

func TestCmpStr_Strict(t *testing.T) {
	if CmpStr_Strict([]string{"s1"}, []string{"s1", "s2"}) {
		t.Error("ill")
		return
	}
	if CmpStr_Strict([]string{"s1", "s2"}, []string{"s2", "s1"}) {
		t.Error("ill")
		return
	}
	if !CmpStr_Strict([]string{"s1", "s2", "s3"}, []string{"s1", "s2", "s3"}) {
		t.Error("ill")
		return
	}
}
