package models

import "github.com/cqroot/openstack-swift-dashboard/databases"

func InitModels() {
	databases.DB.AutoMigrate(&Target{})
	databases.DB.AutoMigrate(&Disk{})
	databases.DB.AutoMigrate(&TotalDisk{})
}
