package main

import (
	"testing"

	"github.com/sensu/sensu-plugin-sdk/sensu"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
}

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	event := corev2.FixtureEvent("entity1", "check1")

	// Valid defaults
	plugin.WarnLoad1 = 2.75
	plugin.WarnLoad5 = 2.5
	plugin.WarnLoad15 = 2.0
	plugin.CritLoad1 = 3.5
	plugin.CritLoad5 = 3.25
	plugin.CritLoad15 = 3.0

	i, e := checkArgs(event)
	assert.Equal(sensu.CheckStateOK, i)
	assert.NoError(e)

	// crit-load1 < warn-load1
	plugin.CritLoad1 = 1.0
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
	plugin.CritLoad1 = 3.5

	// crit-load5 < warn-load5
	plugin.CritLoad5 = 1.0
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
	plugin.CritLoad5 = 3.25

	// crit-load15 < warn-load15
	plugin.CritLoad15 = 1.0
	i, e = checkArgs(event)
	assert.Equal(sensu.CheckStateWarning, i)
	assert.Error(e)
	plugin.CritLoad15 = 3.0
}
