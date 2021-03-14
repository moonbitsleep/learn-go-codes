package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// 方向，上，左，下，右
var dirs = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int{
	// 新建路径二维数组
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 队列
	Q := []point{start}

	for len(Q) > 0 {
		// 取队头，并拿掉
		cur := Q[0]
		Q = Q[1:]

		// 发现终点
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}

	return steps
}


func main() {
	//dir,_ := os.Getwd()
	//fmt.Println("当前路径：",dir) // /Users/ravilian/GolangProjects/src/learning-go
	maze := readMaze("algorithm/BFS/maze/maze.in")
	print2DimMatrix(maze)

	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})

	print2DimMatrix(steps)
}






func print2DimMatrix(maze [][]int) {
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%-3d", val)
		}
		fmt.Println()
	}
	fmt.Println("--------------------")
}



func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}

	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)

	// 创建行
	maze := make([][]int, row)

	for i := range maze {
		//针对每行创建列
		maze[i] = make([]int, col)
		for j := range maze[i] {
			// 如果不取地址，无法获取到值
			_, _ = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}
