// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package sms_home_advertise

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table sms_home_advertise.
type Entity struct {
    Id         int64       `orm:"id,primary"  json:"id"`          //                                          
    Name       string      `orm:"name"        json:"name"`        //                                          
    Type       int         `orm:"type"        json:"type"`        // 轮播位置：0->PC首页轮播；1->app首页轮播  
    Pic        string      `orm:"pic"         json:"pic"`         //                                          
    StartTime  *gtime.Time `orm:"start_time"  json:"start_time"`  //                                          
    EndTime    *gtime.Time `orm:"end_time"    json:"end_time"`    //                                          
    Status     int         `orm:"status"      json:"status"`      // 上下线状态：0->下线；1->上线             
    ClickCount int         `orm:"click_count" json:"click_count"` // 点击数                                   
    OrderCount int         `orm:"order_count" json:"order_count"` // 下单数                                   
    Url        string      `orm:"url"         json:"url"`         // 链接地址                                 
    Note       string      `orm:"note"        json:"note"`        // 备注                                     
    Sort       int         `orm:"sort"        json:"sort"`        // 排序                                     
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