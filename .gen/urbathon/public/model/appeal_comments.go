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

type AppealComments struct {
	ID       int32 `sql:"primary_key"`
	AppealID *int32
	UserID   *int32
	Date     *time.Time
	Text     *string
}