// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRecord = "record"

// Record mapped from table <record>
type Record struct {
	ID        int32 `gorm:"column:id;primaryKey" json:"id"`
	TrackerID int32 `gorm:"column:tracker_id;not null" json:"tracker_id"`
	GpsID     int32 `gorm:"column:gps_id" json:"gps_id"`
	Timestamp int64 `gorm:"column:timestamp;not null" json:"timestamp"`
	Priority  int16 `gorm:"column:priority;not null" json:"priority"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}
