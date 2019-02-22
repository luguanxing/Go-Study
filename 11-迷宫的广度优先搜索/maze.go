package main

import (
	"fmt"
	"os"
)

// 从文件读入矩阵
func readMaze(filename string) [][]int {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 读行读列
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 初始化二维矩阵并读入数据
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

// 点结构
type point struct {
	row, col int
}

// 定义方向，上左下右
var directions = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 走点的方法
func (p point) move(direction point) point {
	return point{p.row + direction.row, p.col + direction.col}
}

// 检查点是否合法(包括边界、墙、走过)
func (p point) isMoveable(maze [][]int, isVisited [][]int) bool {
	if p.row < 0 || p.row >= len(maze) || p.col < 0 || p.col >= len(maze[0]) {
		return false
	}
	if maze[p.row][p.col] == 1 {
		return false
	}
	if isVisited[p.row][p.col] > 0 {
		return false
	}
	return true
}

// BFS走迷宫
func walkMaze(maze [][]int, start point, end point) (canReachEnd bool, steps [][]int) {
	// 用于标记走过的节点
	visitedSteps := make([][]int, len(maze))
	for i := range visitedSteps {
		visitedSteps[i] = make([]int, len(maze[i]))
	}
	visitedSteps[start.row][start.col] = 1
	// 使用队列进行BFS
	queue := []point{start}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		// 找到终点则退出
		if cur == end {
			return true, visitedSteps
		}
		for _, direction := range directions {
			next := cur.move(direction)
			// 检查节点合法且没走过
			if next.isMoveable(maze, visitedSteps) {
				visitedSteps[next.row][next.col] = visitedSteps[cur.row][cur.col] + 1
				queue = append(queue, next)
			}
		}
	}
	return false, visitedSteps
}

// 打印起点到终点路径，反向打印即可
func printResult(visitedSteps [][]int, end point) []point {
	res := make([]point, 0, len(visitedSteps)*len(visitedSteps[0]))
	cur := end
	for visitedSteps[cur.row][cur.col] != 1 {
		for _, direction := range directions {
			next := cur.move(direction)
			if (0 <= next.row && next.row < len(visitedSteps) && 0 <= next.col && next.col < len(visitedSteps[0])) {
				if (visitedSteps[next.row][next.col] == visitedSteps[cur.row][cur.col]-1) {
					res = append(res, cur)
					cur = next
					//找到上一个点则停止四周环顾(已修改cur)，需要重新开始
					break
				}
			}
		}
	}
	// 加上起点
	res = append(res, cur)
	return res
}

func main() {
	maze := readMaze("maze.in")
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	canReachEnd, visitedSteps := walkMaze(maze, start, end)
	for _, row := range visitedSteps {
		for _, val := range row {
			fmt.Printf("%5d", val)
		}
		fmt.Println()
	}
	if canReachEnd {
		fmt.Println("可以找到终点")
		result := printResult(visitedSteps, end)
		// 翻转slice
		for i,j:=0,len(result)-1; i<j; i,j = i+1,j-1 {
			result[i], result[j] = result[j], result[i]
		}
		// 打印路径
		for i, p := range result {
			if (i != len(result) - 1) {
				fmt.Print(p, " -> ")
			} else {
				fmt.Print(p)
			}
		}
	} else {
		fmt.Println("无法找到终点")
	}
}
