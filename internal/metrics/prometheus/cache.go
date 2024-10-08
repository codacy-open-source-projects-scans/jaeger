// Copyright (c) 2017 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package prometheus

import (
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// vectorCache is used to avoid creating Prometheus vectors with the same set of labels more than once.
type vectorCache struct {
	registerer prometheus.Registerer
	lock       sync.Mutex
	cVecs      map[string]*prometheus.CounterVec
	gVecs      map[string]*prometheus.GaugeVec
	hVecs      map[string]*prometheus.HistogramVec
}

func newVectorCache(registerer prometheus.Registerer) *vectorCache {
	return &vectorCache{
		registerer: registerer,
		cVecs:      make(map[string]*prometheus.CounterVec),
		gVecs:      make(map[string]*prometheus.GaugeVec),
		hVecs:      make(map[string]*prometheus.HistogramVec),
	}
}

func (c *vectorCache) getOrMakeCounterVec(opts prometheus.CounterOpts, labelNames []string) *prometheus.CounterVec {
	c.lock.Lock()
	defer c.lock.Unlock()

	cacheKey := c.getCacheKey(opts.Name, labelNames)
	cv, cvExists := c.cVecs[cacheKey]
	if !cvExists {
		cv = prometheus.NewCounterVec(opts, labelNames)
		c.registerer.MustRegister(cv)
		c.cVecs[cacheKey] = cv
	}
	return cv
}

func (c *vectorCache) getOrMakeGaugeVec(opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	c.lock.Lock()
	defer c.lock.Unlock()

	cacheKey := c.getCacheKey(opts.Name, labelNames)
	gv, gvExists := c.gVecs[cacheKey]
	if !gvExists {
		gv = prometheus.NewGaugeVec(opts, labelNames)
		c.registerer.MustRegister(gv)
		c.gVecs[cacheKey] = gv
	}
	return gv
}

func (c *vectorCache) getOrMakeHistogramVec(opts prometheus.HistogramOpts, labelNames []string) *prometheus.HistogramVec {
	c.lock.Lock()
	defer c.lock.Unlock()

	cacheKey := c.getCacheKey(opts.Name, labelNames)
	hv, hvExists := c.hVecs[cacheKey]
	if !hvExists {
		hv = prometheus.NewHistogramVec(opts, labelNames)
		c.registerer.MustRegister(hv)
		c.hVecs[cacheKey] = hv
	}
	return hv
}

func (*vectorCache) getCacheKey(name string, labels []string) string {
	return strings.Join(append([]string{name}, labels...), "||")
}
