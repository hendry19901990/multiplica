package grpc

import (
  "time"
  "context"
  "testing"
  "google.golang.org/grpc"

  "github.com/stretchr/testify/assert"
)

func TestHttpHandlersOK(t *testing.T) {

  for i := 1; i <= 10; i++{
    resp, err := CallServerGrpc(MultiplyRequest{NumberA: int32(i), NumberB: int32(3)})
    if err != nil {
       assert.Nil(t, err)
    }
    assert.Equal(t, int32(i*3), resp.Result, "they should be equal")
  }
}

func CallServerGrpc(req MultiplyRequest) (*MultiplyResponse, error){
   conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
   if err != nil {
     return nil, err
   }
   defer conn.Close()
   c := NewMultiplyServiceClient(conn)

   ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Multiply(ctx, &req)
	if err != nil {
		return nil, err
	}

  return r, nil
}
