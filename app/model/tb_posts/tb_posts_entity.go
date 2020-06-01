// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package tb_posts

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table tb_posts.
type Entity struct {
    Id          uint64      `orm:"id,primary"   json:"id"`           //                   
    PublishName string      `orm:"publish_name" json:"publish_name"` // 发布人            
    PublishTime *gtime.Time `orm:"publish_time" json:"publish_time"` // 发布时间          
    PostType    string      `orm:"post_type"    json:"post_type"`    // BANNER,NOTICE,AD  
    Title       string      `orm:"title"        json:"title"`        // 标题              
    Url         string      `orm:"url"          json:"url"`          // 跳转路径          
    Banner      string      `orm:"banner"       json:"banner"`       // 图片              
    SubContent  string      `orm:"sub_content"  json:"sub_content"`  // 简述              
    Content     string      `orm:"content"      json:"content"`      // 详情              
    UpdateAt    *gtime.Time `orm:"update_at"    json:"update_at"`    //                   
    CreateAt    *gtime.Time `orm:"create_at"    json:"create_at"`    //                   
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