package loadbalance

import (
	"errors"
	"math/rand"
)

type RandomLB struct {
	curIndex int
	rss      []string
}

func (lb *RandomLB) Add(ips ...string) error {
	if len(ips) == 0 {
		return errors.New("need at least 1 ip")
	}

	lb.rss = append(lb.rss, ips...)
	return nil
}

func (lb *RandomLB) Next() string {
	if len(lb.rss) == 0 {
		return ""
	}

	lb.curIndex = rand.Intn(len(lb.rss))
	return lb.rss[lb.curIndex]
}

func (lb *RandomLB) Get() (string, error) {
	return lb.Next(), nil
}
