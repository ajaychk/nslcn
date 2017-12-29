package models

import "time"

type commStatusType int

const (
	CommStatusUnknown = iota
	CommStatusRegistered
	CommStatusJoined
	CommStatusUplinkReceived
)

// Node is light device
type Node struct {
	ID         string         `orm:"column(id);pk" json:"id"`
	GatewayID  string         `orm:"column(gwid)" json:"gwid"`
	LightType  int            `orm:"default(0)" json:"light_type"`
	CommStatus commStatusType `orm:"default(0)" json:"comm_status"`
	LastSeen   time.Time      `orm:"auto_now;type(datetime)" json:"last_seen,omitempty"`
}
