package graph

type node struct {
	id int

	xpos int
	ypos int
}

type edge struct {
	src int
	dst int

	val int
}

// 邻接表 map[(node id)]map[(node id)](edge value)
var ajlist map[int]map[int]edge = make(map[int]map[int]edge)

// 节点值
var nodes map[int]node = make(map[int]node)

func setEdge(src int, dst int, val int) {
	srcmap, _ := ajlist[src]
	dstmap, _ := ajlist[dst]
	if srcmap == nil {
		srcmap = make(map[int]edge)
	}
	if dstmap == nil {
		dstmap = make(map[int]edge)
	}
	srcmap[dst] = edge{
		src: src,
		dst: dst,
		val: val,
	}
	dstmap[src] = edge{
		src: dst,
		dst: src,
		val: val,
	}
	ajlist[src] = srcmap
	ajlist[dst] = dstmap
}

func setNode(id int, xpos int, ypos int) {
	nodes[id] = node{
		id:   id,
		xpos: xpos,
		ypos: ypos,
	}
}
