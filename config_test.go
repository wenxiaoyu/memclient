package gmc

import (
	"testing"
)

func TestConfig(t *testing.T) {
	if len(McConfig.Host) != 3 {
		t.Fatal("Hosts is not equal 3")
	}
	
}
