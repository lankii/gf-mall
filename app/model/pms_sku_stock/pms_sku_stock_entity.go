// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package pms_sku_stock

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table pms_sku_stock.
type Entity struct {
    Id             int64   `orm:"id,primary"      json:"id"`              //                         
    ProductId      int64   `orm:"product_id"      json:"product_id"`      //                         
    SkuCode        string  `orm:"sku_code"        json:"sku_code"`        // sku编码                 
    Price          float64 `orm:"price"           json:"price"`           //                         
    Stock          int     `orm:"stock"           json:"stock"`           // 库存                    
    LowStock       int     `orm:"low_stock"       json:"low_stock"`       // 预警库存                
    Pic            string  `orm:"pic"             json:"pic"`             // 展示图片                
    Sale           int     `orm:"sale"            json:"sale"`            // 销量                    
    PromotionPrice float64 `orm:"promotion_price" json:"promotion_price"` // 单品促销价格            
    LockStock      int     `orm:"lock_stock"      json:"lock_stock"`      // 锁定库存                
    SpData         string  `orm:"sp_data"         json:"sp_data"`         // 商品销售属性，json格式  
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