package main

import (
	"fmt"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	WarnLoad1  float64
	WarnLoad5  float64
	WarnLoad15 float64
	CritLoad1  float64
	CritLoad5  float64
	CritLoad15 float64
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-load",
			Short:    "Check system load average per CPU core",
			Keyspace: "sensu.io/plugins/check-load/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "warn-load1",
			Argument:  "warn-load1",
			Shorthand: "",
			Default:   float64(2.75),
			Usage:     "Warning threshold for 1-minute per-core load average",
			Value:     &plugin.WarnLoad1,
		},
		{
			Path:      "warn-load5",
			Argument:  "warn-load5",
			Shorthand: "",
			Default:   float64(2.5),
			Usage:     "Warning threshold for 5-minute per-core load average",
			Value:     &plugin.WarnLoad5,
		},
		{
			Path:      "warn-load15",
			Argument:  "warn-load15",
			Shorthand: "",
			Default:   float64(2.0),
			Usage:     "Warning threshold for 15-minute per-core load average",
			Value:     &plugin.WarnLoad15,
		},
		{
			Path:      "crit-load1",
			Argument:  "crit-load1",
			Shorthand: "",
			Default:   float64(3.5),
			Usage:     "Critical threshold for 1-minute per-core load average",
			Value:     &plugin.CritLoad1,
		},
		{
			Path:      "crit-load5",
			Argument:  "crit-load5",
			Shorthand: "",
			Default:   float64(3.25),
			Usage:     "Critical threshold for 5-minute per-core load average",
			Value:     &plugin.CritLoad5,
		},
		{
			Path:      "crit-load15",
			Argument:  "crit-load15",
			Shorthand: "",
			Default:   float64(3.0),
			Usage:     "Critical threshold for 15-minute per-core load average",
			Value:     &plugin.CritLoad15,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if plugin.CritLoad1 < plugin.WarnLoad1 {
		return sensu.CheckStateWarning, fmt.Errorf("--crit-load1 must be >= --warn-load1")
	}
	if plugin.CritLoad5 < plugin.WarnLoad5 {
		return sensu.CheckStateWarning, fmt.Errorf("--crit-load5 must be >= --warn-load5")
	}
	if plugin.CritLoad15 < plugin.WarnLoad15 {
		return sensu.CheckStateWarning, fmt.Errorf("--crit-load15 must be >= --warn-load15")
	}
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	avg, err := load.Avg()
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to read load average: %v", err)
	}

	cpuCount, err := cpu.Counts(true)
	if err != nil || cpuCount == 0 {
		return sensu.CheckStateCritical, fmt.Errorf("failed to determine CPU count: %v", err)
	}

	load1 := avg.Load1 / float64(cpuCount)
	load5 := avg.Load5 / float64(cpuCount)
	load15 := avg.Load15 / float64(cpuCount)

	perfData := fmt.Sprintf("load1=%.2f;%.2f;%.2f load5=%.2f;%.2f;%.2f load15=%.2f;%.2f;%.2f",
		load1, plugin.WarnLoad1, plugin.CritLoad1,
		load5, plugin.WarnLoad5, plugin.CritLoad5,
		load15, plugin.WarnLoad15, plugin.CritLoad15,
	)

	status := "OK"
	state := sensu.CheckStateOK

	if load1 >= plugin.CritLoad1 || load5 >= plugin.CritLoad5 || load15 >= plugin.CritLoad15 {
		status = "Critical"
		state = sensu.CheckStateCritical
	} else if load1 >= plugin.WarnLoad1 || load5 >= plugin.WarnLoad5 || load15 >= plugin.WarnLoad15 {
		status = "Warning"
		state = sensu.CheckStateWarning
	}

	fmt.Printf("%s %s: Per core load average (%d CPU): %.2f, %.2f, %.2f | %s\n",
		plugin.Name, status, cpuCount, load1, load5, load15, perfData)

	return state, nil
}
