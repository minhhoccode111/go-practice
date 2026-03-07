package unionfind

type IUnionFind interface {
	Find(int) int
	Union(int, int)
	Connected(int, int) bool
	Count() int
}

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func New(n int) *UnionFind {
	if n <= 0 {
		panic("unionfind: invalid size")
	}

	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}

	for i := range uf.parent {
		uf.parent[i] = i
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	uf.validate(x)

	for x != uf.parent[x] {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}

	return x
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}

	uf.count--

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
		return
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
		return
	}

	uf.rank[rootX]++
	uf.parent[rootY] = rootX
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) validate(x int) {
	if x < 0 || x >= len(uf.parent) {
		panic("unionfind: index out of range")
	}
}
