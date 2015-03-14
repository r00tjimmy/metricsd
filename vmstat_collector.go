package main

import "github.com/Sirupsen/logrus"
import "github.com/c9s/goprocinfo/linux"

type VmstatCollector struct{}

func (c *VmstatCollector) Collect() (IntMetricMap, error) {
	stat, err := linux.ReadVMStat("/proc/vmstat")
	if err != nil {
		logrus.Fatal("stat read fail")
		return nil, err
	}

	return IntMetricMap{
		"paging_in": stat.PagePagein,
		"pagingout": stat.PagePageout,
		"swap_in":   stat.PageSwapin,
		"swap_out":  stat.PageSwapout,
	}, nil
}

func (c *VmstatCollector) Report() (MetricMapSlice, error) {
	var report MetricMapSlice
	values, _ := c.Collect()

	if values != nil {
		for k, v := range values {
			report = append(report, MetricMap{
				"target_type": "rate",
				"type":        k,
				"unit":        "Page",
				"result":      v,
			})
		}
	}

	return report, nil
}
