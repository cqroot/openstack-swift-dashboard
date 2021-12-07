package models

import "github.com/cqroot/openstack-swift-dashboard/databases"

type Target struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"primaryKey;unique;not null"`
	Endpoint       string `gorm:"unique;not null"`
	ScrapeInterval string `gorm:"default:30m"`
	ScrapeTimeout  string `gorm:"default:30m"`
}

func TargetList() ([]Target, error) {
	var targets []Target
	err := databases.DB.Find(&targets).Error
	return targets, err
}
