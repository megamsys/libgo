package bills

// import (
// //	"fmt"
// // constants   "github.com/megamsys/libgo/utils"
// "testing"
// 	"gopkg.in/check.v1"
// )

// func Test(t *testing.T) { check.TestingT(t) }
//
// func (s *S) TestOnboard(c *check.C) {
// 	w := &whmcsBiller{}
//   o := &BillOpts{
//     AccountId: "tour@megam.io",
//   }
//   m := make(map[string]string)
//   m[constants.DOMAIN] = ""
//   m[constants.USERNAME] = ""
//   m[constants.PASSWORD] = ""
//   m[constants.VERTICE_EMAIL] = ""
//   m[constants.VERTICE_APIKEY] = ""
//   m[constants.SCYLLAHOST] = "192.168.0.116"
//   m[constants.SCYLLAKEYSPACE] = "test"
//   err := w.Onboard(o,m)
//
// 	c.Assert(err, check.IsNil)
// }
//
// func (s *S) TestDeduct(c *check.C) {
// 	w := &whmcsBiller{}
// 	o := &BillOpts{
//     AccountId: "info@megam.io",
// 		AssemblyName: "ASM123"
//   }
//   m := make(map[string]string)
// 	mp[constants.USERNAME] = "test@megam.io"
// 	mp[constants.WHMCS_PASSWORD] = "asdf"
// 	mp[constants.WHMCS_APIKEY] = "1234"
// 	mp[constants.DOMAIN] = "http://192.168.0.100"
// 	mp[constants.PIGGYBANKS] = "scylladb,whmcs"
//   err := w.Onboard(o,mp)
// 	c.Assert(err, check.IsNil)
// }
