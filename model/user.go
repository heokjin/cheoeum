package model

type User struct {
	ID        int            `json:"id" gorm:"primary_key"`
	ClusterID int            `json:"clusterID"`
	UserID    string         `json:"userID"`
}
