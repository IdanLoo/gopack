package model

import (
	"strings"
)

// PushBody is the body of push event sent by git server
type PushBody struct {
	Ref        string     `json:"ref" binding:"required"`
	Repository Repository `json:"repository" binding:"required"`
}

// Branch to get the branch name of this ref
func (p *PushBody) Branch() string {
	refs := strings.Split(p.Ref, "/")
	return refs[len(refs)-1]
}
