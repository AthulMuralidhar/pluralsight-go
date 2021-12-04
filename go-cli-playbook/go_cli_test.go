package main

import "testing"

func Test(t *testing.T) {
	t.Run("add stuff", func(t *testing.T) {
		if 1+1 !=2 {
			t.Fail()
		}
	})
}