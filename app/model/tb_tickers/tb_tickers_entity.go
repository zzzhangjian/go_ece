// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package tb_tickers

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table tb_tickers.
type Entity struct {
    Id         uint        `orm:"id,primary"  json:"id"`          // 主键                        
    Symbol     string      `orm:"symbol"      json:"symbol"`      // 交易对                      
    Open       string      `orm:"open"        json:"open"`        // 开盘价                      
    Close      string      `orm:"close"       json:"close"`       // 收盘价                      
    High       string      `orm:"high"        json:"high"`        // 最高价                      
    Low        string      `orm:"low"         json:"low"`         // 最低价                      
    Rate       float64     `orm:"rate"        json:"rate"`        // 长跌幅 close/open           
    Vol        float64     `orm:"vol"         json:"vol"`         // 交易量                      
    Amount     float64     `orm:"amount"      json:"amount"`      // 以基础币种计量的交易量      
    Count      int         `orm:"count"       json:"count"`       // 交易次数（以滚动24小时计）  
    CreateTime *gtime.Time `orm:"create_time" json:"create_time"` // 时间                        
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