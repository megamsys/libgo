package api
import (
   "fmt"
   "net/http"
  	"gopkg.in/check.v1"
)

func (s *S) TestGetUser(c *check.C) {
  cl := NewClient(&http.Client{}, map[string]string{"email": "thilak@b.com", "password": "team4megam"})
  response, err := cl.Get()
  //url and host
  fmt.Println("************************")
  fmt.Println(response)
  fmt.Println(err)
  c.Assert(cl, check.NotNil)
  err = fmt.Errorf("error")
  c.Assert(err, check.IsNil)
}
