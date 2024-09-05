package storage

import "time"

type GaugeItem struct {
	Value float64
	Date  time.Time
}

type CounterItem struct {
	Value int64
	Date  time.Time
}

type MemStorage struct {
	Gauge   map[string]GaugeItem
	Counter map[string][]CounterItem
}

var Storage MemStorage

func (storage *MemStorage) AddCounter(name string, item CounterItem) {
	storage.Counter[name] = append(storage.Counter[name], item)
}

func (storage *MemStorage) AddGauge(name string, item GaugeItem) {
	storage.Gauge[name] = item
}
