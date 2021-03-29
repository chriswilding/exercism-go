package tree

import (
	"errors"
	"sort"
)

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))

	for i, record := range records {
		if record.ID != i || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("error")
		}
		node := &Node{ID: record.ID}
		nodes[i] = node
		if i != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, node)
		}
	}

	return nodes[0], nil
}
