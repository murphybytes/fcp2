package tufer

import (
  "testing"
  "fmt"
  )

  func TestPack( t *testing.T ) {
    expected := "one two three "
    actual   := Pack( "one", "two", "three" )
    if actual != expected {
      t.Error(  fmt.Sprintf("Failed actual [%s] expected [%s]", actual, expected ) )
    }
  }

  func TestUnpack( t *testing.T ) {
    packege  := "one two three "
    expected := []string{ "one", "two", "three" }
    actual   := Unpack( packege )

    if len(expected) != len(actual) {
      t.Error( "Length mismatch expected", len(expected), "actual", len(actual))
    }

  }
