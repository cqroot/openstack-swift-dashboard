package models

import (
	"github.com/cqroot/openstack-swift-dashboard/databases"
	"gorm.io/gorm/clause"
)

type Disk struct {
	Host   string `gorm:"primaryKey;autoIncrement:false"`
	Device string `gorm:"primaryKey;autoIncrement:false"`
	Target uint   `gorm:"not null"`
	Avail  int64  `gorm:"type:bigint"`
	Used   int64  `gorm:"type:bigint"`
	Size   int64  `gorm:"type:bigint"`
	Usage  int64  `gorm:"type:bigint"`
}

func UpdateDisks(disks *[]Disk) {
	databases.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(disks)
}
