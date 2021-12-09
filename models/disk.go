package models

import (
	"github.com/cqroot/openstack-swift-dashboard/databases"
	"gorm.io/gorm/clause"
)

type Disk struct {
	Host   string  `gorm:"primaryKey;autoIncrement:false"`
	Device string  `gorm:"primaryKey;autoIncrement:false"`
	Target uint    `gorm:"not null"`
	Avail  int64   `gorm:"type:bigint"`
	Used   int64   `gorm:"type:bigint"`
	Size   int64   `gorm:"type:bigint"`
	Usage  float64 `gorm:"index:"`
}

func DiskList(target uint, limit int, offset int, desc bool) ([]Disk, error) {
	var disks []Disk
	order := "`usage`"
	if desc {
		order = "`usage` desc"
	}
	err := databases.DB.Where("target = ?", target).Order(order).Limit(limit).Offset(offset).Find(&disks).Error
	return disks, err
}

func UpdateDisks(disks *[]Disk) {
	databases.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(disks)
}
