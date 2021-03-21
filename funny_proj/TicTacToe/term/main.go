package main

import "fmt"

type Game struct {
	Chessboard [][]byte
	// 0 未结束 1 X胜 2 O胜 3 平局
	status int
	// 记录步数，奇数为X偶数为O
	step int
}

func NewGame() Game {
	board := make([][]byte, 3)
	for i := 0; i < 3; i++ {
		board[i] = make([]byte, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = '.'
		}
	}
	g := Game{board, 0, 1}
	return g
}

func (g *Game) PrintBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", g.Chessboard[i][j])
		}
		fmt.Println()
	}
}

func (g *Game) CheckPositionSafe(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x > 2 || y > 2 {
		return false
	}
	if g.Chessboard[x][y] != '.' {
		return false
	}
	return true
}

func isOdd(num int) bool {
	if num%2 != 0 {
		return true
	}
	return false
}

func (g *Game) TakeAStep(x, y int) {
	var symbol byte
	if isOdd(g.step) {
		symbol = 'X'
	} else {
		symbol = 'O'
	}

	if g.CheckPositionSafe(x, y) {
		g.Chessboard[x][y] = symbol
	} else {
		fmt.Println("无法落子！")
		return
	}
	g.step += 1
}

func getPosition() (int, int) {
	var x, y int
	fmt.Println("请输入您要落子的位置:")
	_, _ = fmt.Scanf("%d,%d", &x, &y)
	x = x - 1
	y = y - 1
	return x, y
}

func (g *Game) CheckDraw() bool{
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pos := g.Chessboard[i][j]
			if pos == '.' {
				return false
			}
		}
	}
	return true
}

func (g *Game) CheckWon(symbol byte) bool{
	var success = [][][]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},

		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},

		{{0, 0}, {1, 1}, {2, 2}},
		{{2, 0}, {1, 1}, {0, 2}},
	}
	for _, line := range success {
		x1 := line[0][0]
		y1 := line[0][1]
		x2 := line[1][0]
		y2 := line[1][1]
		x3 := line[2][0]
		y3 := line[2][1]
		if g.Chessboard[x1][y1] == symbol &&
			g.Chessboard[x2][y2] == symbol &&
			g.Chessboard[x3][y3] == symbol {
				return true
		}
	}
	return false
}

func main() {
	game := NewGame()

	for game.status == 0 {
		game.PrintBoard()
		x, y := getPosition()
		game.TakeAStep(x, y)
		if game.CheckWon('X') {
			game.PrintBoard()
			fmt.Println("X Won!")
			game.status = 1
		}
		if game.CheckWon('O') {
			game.PrintBoard()
			fmt.Println("O Won!")
			game.status = 2
		}
		if game.CheckDraw() {
			game.PrintBoard()
			fmt.Println("平局！")
			game.status = 3
		}
	}


}
