// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package ums_admin

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table ums_admin.
type Entity struct {
    Id         int64       `orm:"id,primary"  json:"id"`          //                                 
    Username   string      `orm:"username"    json:"username"`    //                                 
    Password   string      `orm:"password"    json:"password"`    //                                 
    Icon       string      `orm:"icon"        json:"icon"`        // 头像                            
    Email      string      `orm:"email"       json:"email"`       // 邮箱                            
    NickName   string      `orm:"nick_name"   json:"nick_name"`   // 昵称                            
    Note       string      `orm:"note"        json:"note"`        // 备注信息                        
    CreateTime *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间                        
    LoginTime  *gtime.Time `orm:"login_time"  json:"login_time"`  // 最后登录时间                    
    Status     int         `orm:"status"      json:"status"`      // 帐号启用状态：0->禁用；1->启用  
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}