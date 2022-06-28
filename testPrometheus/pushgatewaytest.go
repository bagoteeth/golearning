package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"time"
)

var (
	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last completion of a DB backup, successful or not.",
	})
	successTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_success_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	duration = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "db_backup_duration_seconds",
		Help: "The duration of the last DB backup in seconds.",
	})
)

func main() {

	registry := prometheus.NewRegistry()
	registry.MustRegister(completionTime, duration, successTime)
	// Note that successTime is not registered.
	pushers := make(map[string]*push.Pusher)
	for i := 0; ; i += 3 {
		fmt.Println(i)
		pushers["a"] = push.New("http://localhost:9091", "db-backup")
		if err := pushers["a"].Push(); err != nil {
			fmt.Printf("0: %s\n", err.Error())
		}
		time.Sleep(10 * time.Second)

		fmt.Println(i + 1)
		pushers["a"].Gatherer(registry)
		if err := pushers["a"].Push(); err != nil {
			fmt.Printf("1: %s\n", err.Error())
		}
		time.Sleep(10 * time.Second)

		fmt.Println(i + 2)
		if err := pushers["a"].Delete(); err != nil {
			fmt.Printf("2: %s\n", err.Error())
		}
		time.Sleep(10 * time.Second)
	}

}
