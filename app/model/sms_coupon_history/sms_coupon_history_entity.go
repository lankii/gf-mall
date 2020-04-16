// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package sms_coupon_history

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table sms_coupon_history.
type Entity struct {
    Id             int64       `orm:"id,primary"      json:"id"`              //                                            
    CouponId       int64       `orm:"coupon_id"       json:"coupon_id"`       //                                            
    MemberId       int64       `orm:"member_id"       json:"member_id"`       //                                            
    CouponCode     string      `orm:"coupon_code"     json:"coupon_code"`     //                                            
    MemberNickname string      `orm:"member_nickname" json:"member_nickname"` // 领取人昵称                                 
    GetType        int         `orm:"get_type"        json:"get_type"`        // 获取类型：0->后台赠送；1->主动获取         
    CreateTime     *gtime.Time `orm:"create_time"     json:"create_time"`     //                                            
    UseStatus      int         `orm:"use_status"      json:"use_status"`      // 使用状态：0->未使用；1->已使用；2->已过期  
    UseTime        *gtime.Time `orm:"use_time"        json:"use_time"`        // 使用时间                                   
    OrderId        int64       `orm:"order_id"        json:"order_id"`        // 订单编号                                   
    OrderSn        string      `orm:"order_sn"        json:"order_sn"`        // 订单号码                                   
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