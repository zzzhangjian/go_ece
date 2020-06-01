// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package tb_chain_transactions

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table tb_chain_transactions.
type Entity struct {
    Txid             string      `orm:"txid,primary"     json:"txid"`             // 交易ID                   
    BlockNumber      int64       `orm:"block_number"     json:"block_number"`     // 区块高度                 
    TxType           string      `orm:"tx_type"          json:"tx_type"`          // 交易类型 omni;erc20;fee  
    Amount           float64     `orm:"amount"           json:"amount"`           // 总额                     
    Sendingaddress   string      `orm:"sendingaddress"   json:"sendingaddress"`   // 发送方                   
    Referenceaddress string      `orm:"referenceaddress" json:"referenceaddress"` // 接收方                   
    Confirmations    int         `orm:"confirmations"    json:"confirmations"`    // 交易的确认数             
    Fee              float64     `orm:"fee"              json:"fee"`              // 交易手续费               
    Valid            int         `orm:"valid"            json:"valid"`            // 交易是否有效             
    Positioninblock  int         `orm:"positioninblock"  json:"positioninblock"`  // 交易在区块内的序号       
    CreateAt         *gtime.Time `orm:"create_at"        json:"create_at"`        //                          
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