package model

import "github.com/qauzy/trat/pkg/base"

type CreateTask struct {
	Rid string        `json:"rid"`
	Req *base.Request `json:"req"`
	Opt *base.Options `json:"opt"`
}
