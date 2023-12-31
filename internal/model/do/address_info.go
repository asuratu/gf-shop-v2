// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddressInfo is the golang structure of table address_info for DAO operations like Where/Data.
type AddressInfo struct {
	g.Meta    `orm:"table:address_info, do:true"`
	Id        interface{} //
	Name      interface{} //
	Pid       interface{} //
	Status    interface{} //
	UpdatedAt *gtime.Time //
}
