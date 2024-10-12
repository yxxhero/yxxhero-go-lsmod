package lsmod_test

import (
	"testing"

	"github.com/yxxhero/go-lsmod"
)

func TestLsMod(t *testing.T) {
	if _, err := lsmod.LsMod(""); err != nil {
		t.Error(err)
	}
}
