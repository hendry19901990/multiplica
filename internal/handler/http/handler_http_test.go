package http

import (
  "fmt"
  "encoding/json"
  "testing"
  "io/ioutil"
  "net/http"

  "multiplica/internal/entity"

  "github.com/stretchr/testify/assert"
)

func TestHttpHandlersOK(t *testing.T) {
  url := "http://localhost:8080/api/v1/multiply/%d/%d"

  for i := 1; i <= 10; i++{
    byt, err := GetHttp(fmt.Sprintf(url, i, 3))
    if err != nil {
       assert.Nil(t, err)
    }

    var resp entity.MultiplyResponse
    if err := json.Unmarshal(*byt, &resp); err != nil {
         assert.Nil(t, err)
     }

    assert.Equal(t, int32(i*3), resp.Result, "they should be equal")
  }

}

func TestHttpHandlersFail(t *testing.T) {
  urla := "http://localhost:8080/api/v1/multiply/%d.0/%d"

  for i := 1; i < 5; i++{
    _, err := GetHttp(fmt.Sprintf(urla, i, 3))
    if err != nil {
       assert.NotNil(t, err)
    }

  }

  urlb := "http://localhost:8080/api/v1/multiply/letter/%d"

  for i := 1; i < 5; i++{
    _, err := GetHttp(fmt.Sprintf(urlb, i, 3))
    if err != nil {
       assert.NotNil(t, err)
    }

  }

}

func GetHttp(url string) (*[]byte, error){
  resp, err := http.Get(url)
   if err != nil {
      return nil, err
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }

   return &body, nil
}
