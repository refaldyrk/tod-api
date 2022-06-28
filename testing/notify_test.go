package testing

import (
	"testing"
	"tod/helper"
)

func TestNotify(t *testing.T) {
	helper.SendNotify("Test", "Test Message")
}
