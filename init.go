package main

import (
	"flag"
	"strconv"
	"time"
)

// region could be a server, or a request to a server.
type region struct {
	LocationID  string
	NextClosest *region
}

// regions could be requests from a region, or a region hosting a
// server which handles incoming requests.
var regions []*region

// activeRegions are regional servers which have active connections.
var activeRegions = make(map[*region]chan *region)

var totalConns int
var count int
var request_rate int
var response_rate int
var maxConns int
var start time.Time = time.Now()

// init initializes the requestRegions ([]r), and the r[x].NextClosest property
// which indicates the next geographically close region.
func init() {
	flag.IntVar(&count, "regions", 1234, "The number of clusters to create")
	flag.IntVar(&request_rate, "req_rate", 1234, "Sends a request once every (x) milliseconds")
	flag.IntVar(&response_rate, "res_rate", 1234, "Sends a response once every (x) milliseconds")
	flag.IntVar(&maxConns, "max_conns", 1234, "Max connections per node")
	flag.Parse()

	regions = mkRegions(count)
	for i, r := range regions {
		if i+1 == len(regions) {
			r.NextClosest = regions[0]
			return
		}
		r.NextClosest = regions[i+1]
	}
}

func mkRegions(l int) (rs []*region) {
	for i := 1; i <= l; i++ {
		rs = append(rs, &region{"cluster-" + strconv.Itoa(i+100), &region{}})
	}
	return
}
