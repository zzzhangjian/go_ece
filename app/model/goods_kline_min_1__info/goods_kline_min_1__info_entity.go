// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package goods_kline_min_1__info

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table goods_kline_min1_info.
type Entity struct {
    Id              int64       `orm:"id,primary"        json:"id"`                //                   
    Code            string      `orm:"code"              json:"code"`              // 代码              
    Period          string      `orm:"period"            json:"period"`            // 期数 20180606     
    Volume          float64     `orm:"volume"            json:"volume"`            // 成交量            
    Price           string      `orm:"price"             json:"price"`             // 当前价            
    OpeningPrice    float64     `orm:"opening_price"     json:"opening_price"`     // 前一刻开盘价      
    ClosingPrice    float64     `orm:"closing_price"     json:"closing_price"`     // 前一刻收盘价      
    PreClosingPrice string      `orm:"pre_closing_price" json:"pre_closing_price"` // 昨收盘价          
    HighestPrice    float64     `orm:"highest_price"     json:"highest_price"`     // 最高价            
    LowestPrice     float64     `orm:"lowest_price"      json:"lowest_price"`      // 最低价            
    DateYmd         *gtime.Time `orm:"date_ymd"          json:"date_ymd"`          // 日期              
    Date            *gtime.Time `orm:"date"              json:"date"`              // 日期时间          
    CreateTime      *gtime.Time `orm:"create_time"       json:"create_time"`       // 创建时间          
    IsDeleted       uint        `orm:"is_deleted"        json:"is_deleted"`        // 1:删除，0:未删除  
    Timestamp       *gtime.Time `orm:"timestamp"         json:"timestamp"`         //                   
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