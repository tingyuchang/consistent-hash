package consistent_hash

import (
	"hash/crc32"
	"sort"
)

type Ring struct {
	Nodes Nodes
}

func (r *Ring) AddNode(id string) {
	node := &Node{
		Id: id,
	}
	r.Nodes = append(r.Nodes, node)
	sort.Sort(r.Nodes)
}

func (r *Ring) Get(key string) string {
	search := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.ChecksumIEEE([]byte(key))
	}

	i := sort.Search(r.Nodes.Len(), search)

	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}

type Nodes []*Node
type Node struct {
	Id     string
	HashId uint32
}

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].HashId < n[j].HashId
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

