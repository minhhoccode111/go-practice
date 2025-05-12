package tree

// TODO: refactor logic

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

/*

root (ID: 0, parent ID: 0)
|-- child1 (ID: 1, parent ID: 0)
|    |-- grandchild1 (ID: 2, parent ID: 1)
|    +-- grandchild2 (ID: 4, parent ID: 1)
+-- child2 (ID: 3, parent ID: 0)
|    +-- grandchild3 (ID: 6, parent ID: 3)
+-- child3 (ID: 5, parent ID: 0)

{
	0: [1, 3, 5],
	1: [2, 4],
	3: [6],
	2: [],
	4: [],
	5: [],
	6: [],
}

*/

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	dict := map[int][]int{}

	// because children always have id greater than parent
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for _, record := range records {
		// ID must be in range 0 ~ len(records)-1
		if record.ID > len(records)-1 {
			return nil, errors.New("Non continuous ID found")
		}

		// the records is sorted, so every record must be first met when we loop
		if _, ok := dict[record.ID]; ok {
			return nil, errors.New("Duplicate ID found")
		}

		// init children slice
		dict[record.ID] = []int{}

		// record's id must be greater than its parent
		if record.ID < record.Parent {
			return nil, errors.New("Record ID is greater than Parent ID")
		}

		// if we found a record that its parent is not in dict
		if _, ok := dict[record.Parent]; !ok {
			return nil, errors.New("Parent ID does not exist")
		}

		// NOTE: skip if root, can cause stack overflow if add itself to its own children
		if record.ID == 0 {
			continue
		}

		// except root, the record cannot be its own parent
		if record.ID == record.Parent {
			return nil, errors.New("Cycle Found")
		}

		// update children slice of current record's parent
		dict[record.Parent] = append(dict[record.Parent], record.ID)
	}

	// init all nodes and store in a slice
	nodes := make([]Node, len(records))
	for i := range nodes {
		nodes[i] = Node{ID: i}
	}

	// loop through every key-value in dict and build tree
	for id, childIds := range dict {
		// loop through every child id
		for _, childId := range childIds {
			// append child node's pointer to current node children list
			nodes[id].Children = append(nodes[id].Children, &nodes[childId])
		}
	}

	return &nodes[0], nil
}
