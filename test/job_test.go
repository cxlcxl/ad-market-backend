package test

import (
	_ "market/bootstrap"
	"market/cronJobs/jobs/app/logic"
	"testing"
)

func TestApp(t *testing.T) {
	err := logic.NewAppQueryLogic().AppQuery()
	t.Log(err)
}
