package models

import (
	_ "github.com/lib/pq"
)

type GoalTracker struct {
	Id int `json:"id"`
	Jan []int64 `json:"jan"`;
	Feb []int64 `json:"feb"`;
	Mar []int64 `json:"mar"`;
	Apl []int64 `json:"apl"`;
	May []int64 `json:"may"`;
	Jun []int64 `json:"jun"`;
	Jul []int64 `json:"jul"`;
	Aug []int64 `json:"aug"`;
	Sep []int64 `json:"sep"`;
	Oct []int64 `json:"oct"`;
	Nov []int64 `json:"nov"`;
	Dcm []int64 `json:"dcm"`;
	User_id int `json:"user_id"`
}