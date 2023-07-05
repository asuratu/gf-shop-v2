// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Migrations is the golang structure for table migrations.
type Migrations struct {
	Id        uint        `json:"id"        description:""`
	Migration string      `json:"migration" description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
