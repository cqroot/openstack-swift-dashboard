package models

import (
	"github.com/cqroot/openstack-swift-dashboard/databases"
	"gorm.io/gorm/clause"
)

type TotalDisk struct {
	Target     uint  `gorm:"type:bigint;primaryKey;autoIncrement:false"`
	TotalAvail int64 `gorm:"type:bigint"`
	TotalUsed  int64 `gorm:"type:bigint"`
	TotalSize  int64 `gorm:"type:bigint"`
	Date       int   `gorm:"type:int;primaryKey;autoIncrement:false"`
}

func UpdateTotalDisk(totalDisk *TotalDisk) {
	databases.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(totalDisk)
}
