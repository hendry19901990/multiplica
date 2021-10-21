package service

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestMultiplyOK(t *testing.T) {

  for i := 1; i <= 10; i++{
    res := Multiply(int32(i), int32(3))

    assert.Equal(t, int32(i*3), res, "they should be equal")

  }

}
