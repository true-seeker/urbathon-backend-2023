//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type News struct {
	ID             int32 `sql:"primary_key"`
	Title          string
	Body           string
	Date           *time.Time
	CategoryID     *int32
	PhotoURL       *string
	UserID         *int32
	OrganizationID *int32
	PollID         *int32
	Address        *string
	Latitude       *float64
	Longitude      *float64
}
