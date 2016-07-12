package assert

import (
  "testing"
  "runtime"
  "fmt"
)

var test *testing.T = nil

func Init(t *testing.T) {
  test = t
}

func stack() string {
  stack := make([]byte, 1024)
  runtime.Stack(stack, false)
  return string(stack)
}

func Ok(ok bool) {
  if !ok {
    fmt.Printf("%s\r\nIt should be true\r\n", stack())
    test.Fail()
  }
}
