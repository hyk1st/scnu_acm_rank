package model

import (
	"testing"
)

func TestGetUserCompetitions(t *testing.T) {
	GetDB()
	_, _ = GetUserCompetitions()
}
