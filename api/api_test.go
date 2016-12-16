package api

import (
  "io/ioutil"
   "fmt"
  	"gopkg.in/check.v1"
)

// func (s *S) TestGetUser(c *check.C) {
//   a := ApiArgs{
//     Email: s.Email,
//     Url: "http://192.168.10.109:9000/v2",
//     Path: "/accounts/"+s.Email,
//     Api_Key: s.Api_Key,
//     Master_Key: "",
//     Password: "",
//     Org_Id: "",
//   }
//   cl := NewClient(a)
//   // _, _ = cl.Post(s.Assembly)
//   response, err := cl.Get()
//   c.Assert(err, check.IsNil)
//   fmt.Println("************************")
//   fmt.Printf("%#v",response.Body)
//   htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
//  	if err != nil {
//  		fmt.Println("******Error",err)
//  	}
//  	fmt.Println("Success  :",string(htmlData)) //<-- here !
//   fmt.Println(err)
//   c.Assert(cl, check.NotNil)
//   c.Assert(err, check.IsNil)
// }

func (s *S) TestGetAssembly(c *check.C) {
  a := ApiArgs{
    Email: s.Email,
    Url: "http://apidb.megam.io:9000/v2",
    Path: "/assembly/ASM5285833184590940525",
  //  Api_Key: s.Api_Key,
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
    Org_Id: "",
  }
  cl := NewClient(a)
  response, err := cl.Get()
  c.Assert(err, check.IsNil)
  htmlData, err := ioutil.ReadAll(response.Body) //<--- here!
 	if err != nil {
 		fmt.Println("******Error",err)
 	}
 	fmt.Println("Success  :",string(htmlData)) //<-- here !
  fmt.Println(err)
  c.Assert(cl, check.NotNil)
  c.Assert(err, check.IsNil)
}
