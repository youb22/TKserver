// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGPS = "gps"

// GPS mapped from table <gps>
type GPS struct {
	ID    int32   `gorm:"column:id;primaryKey" json:"id"`
	Lat   float64 `gorm:"column:lat;not null" json:"lat"`
	Long  float64 `gorm:"column:long;not null" json:"long"`
	Speed int16   `gorm:"column:speed;not null" json:"speed"`
}

// TableName Gp's table name
func (*GPS) TableName() string {
	return TableNameGPS
}
