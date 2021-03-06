// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package oms_company_address

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table oms_company_address.
type Entity struct {
    Id            int64  `orm:"id,primary"     json:"id"`             //                                 
    AddressName   string `orm:"address_name"   json:"address_name"`   // 地址名称                        
    SendStatus    int    `orm:"send_status"    json:"send_status"`    // 默认发货地址：0->否；1->是      
    ReceiveStatus int    `orm:"receive_status" json:"receive_status"` // 是否默认收货地址：0->否；1->是  
    Name          string `orm:"name"           json:"name"`           // 收发货人姓名                    
    Phone         string `orm:"phone"          json:"phone"`          // 收货人电话                      
    Province      string `orm:"province"       json:"province"`       // 省/直辖市                       
    City          string `orm:"city"           json:"city"`           // 市                              
    Region        string `orm:"region"         json:"region"`         // 区                              
    DetailAddress string `orm:"detail_address" json:"detail_address"` // 详细地址                        
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