package loadbalance

import "errors"

type RoundRobinLB struct {
	curIndex int
	rss      []string
}

func (lb *RoundRobinLB) Add(ips ...string) error {
	if len(ips) == 0 {
		return errors.New("need at least 1 ip")
	}

	lb.rss = append(lb.rss, ips...)
	return nil
}

func (lb *RoundRobinLB) Next() string {
	if len(lb.rss) == 0 {
		return ""
	}

	if lb.curIndex == len(lb.rss) {
		lb.curIndex = 0
	}

	addr := lb.rss[lb.curIndex]
	lb.curIndex++
	return addr
}

func (lb *RoundRobinLB) Get() (string, error) {
	return lb.Next(), nil
}