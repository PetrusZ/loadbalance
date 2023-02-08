package loadbalance

type WeightRoundRobinLB struct {
	curIndex int
	rss      []*WeightNode
}

type WeightNode struct {
	addr            string
	Weight          int //初始化时对节点约定的权重
	currentWeight   int //节点临时权重，每轮都会变化
	effectiveWeight int //有效权重, 默认与weight相同 , totalWeight = sum(effectiveWeight)  //出现故障就-1
}

//1, currentWeight = currentWeight + effectiveWeight
//2, 选中最大的currentWeight节点为选中节点
//3, currentWeight = currentWeight - totalWeight

func (lb *WeightRoundRobinLB) Add(addr string, weight int) error {
	node := &WeightNode{
		addr:            addr,
		Weight:          weight,
		effectiveWeight: weight,
	}

	lb.rss = append(lb.rss, node)
	return nil
}

func (lb *WeightRoundRobinLB) Next() string {
	var best *WeightNode
	total := 0

	for i := 0; i < len(lb.rss); i++ {
		node := lb.rss[i]
		//1 计算所有有效权重
		total += node.effectiveWeight
		//2 修改当前节点临时权重
		node.currentWeight += node.effectiveWeight
		//3 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1，直到恢复到weight大小
		if node.effectiveWeight < node.Weight {
			node.effectiveWeight++
		}
		//4 选中最大临时权重节点
		if best == nil || node.currentWeight > best.currentWeight {
			best = node
		}
	}

	if best == nil {
		return ""
	}
	//5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (lb *WeightRoundRobinLB) Get() (string, error) {
	return lb.Next(), nil
}

func (lb *WeightRoundRobinLB) Update() {

}
