// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package ums_member_receive_address

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table ums_member_receive_address.
type Entity struct {
    Id            int64  `orm:"id,primary"     json:"id"`             //                 
    MemberId      int64  `orm:"member_id"      json:"member_id"`      //                 
    Name          string `orm:"name"           json:"name"`           // 收货人名称      
    PhoneNumber   string `orm:"phone_number"   json:"phone_number"`   //                 
    DefaultStatus int    `orm:"default_status" json:"default_status"` // 是否为默认      
    PostCode      string `orm:"post_code"      json:"post_code"`      // 邮政编码        
    Province      string `orm:"province"       json:"province"`       // 省份/直辖市     
    City          string `orm:"city"           json:"city"`           // 城市            
    Region        string `orm:"region"         json:"region"`         // 区              
    DetailAddress string `orm:"detail_address" json:"detail_address"` // 详细地址(街道)  
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