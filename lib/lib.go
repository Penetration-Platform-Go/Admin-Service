package lib

// Warshall algorithm
func Warshall(network [][]bool) {
	count := len(network)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if network[j][i] {
				for k := 0; k < count; k++ {
					network[j][k] = network[j][k] || network[i][k]
				}
			}
		}
	}
}

// InitMatrix for operation
func InitMatrix(network [][]bool) {
	count := len(network)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			network[i][j] = network[i][j] || network[j][i]
		}
	}
}

// FindChildGraphNumber Algorithm
func FindChildGraphNumber(network [][]bool) int {
	count := len(network)
	number := 0
	var graphType map[int]int
	graphType = make(map[int]int)
	for i := 0; i < count; i++ {
		_, exist := graphType[i]
		if !exist {
			number++
			for j := 0; j < count; j++ {
				if network[i][j] {
					graphType[j] = number
				}
			}
		}
	}
	return number
}

// DFS Algorithm
func DFS(network [][]bool, path []bool, begin, end int) []bool {
	count := len(network)
	for i := 0; i < count; i++ {
		if i == end {
			break
		}
		if i == begin {
			continue
		}
		if path[i] {
			continue
		}
		if network[begin][i] {
			path[i] = true
			DFS(network, path, i, end)
			path[i] = false
		}
	}
	return path
}

var max = 1000

// Dijkstra Algorithm
func Dijkstra(network [][]bool) [][]int {
	count := len(network)
	var path [][]int
	path = make([][]int, count)
	for i := 0; i < count; i++ {
		path[i] = make([]int, count)
	}
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i == j {
				path[i][j] = 0
				continue
			} else if network[i][j] {
				path[i][j] = 1
			} else {
				path[i][j] = max
			}
		}
	}
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			for k := 0; k < count; k++ {
				if path[k][j] > path[k][i]+path[i][j] {
					path[k][j] = path[k][i] + path[i][j]
					path[j][k] = path[k][j]
				}
			}
		}
	}
	return path
}
