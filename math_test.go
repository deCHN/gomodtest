package hello

import (
   "testing"
   algo "github.com/dechn/algorithm/v2"
)

func TestAddition(t *testing.T) {
   want := 33
   if get := algo.Add(12,21); get != want {
      t.Fatalf("错误.Failed. Want %d but got %d.", want, get)
   }
}
