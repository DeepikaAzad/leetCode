package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type point struct {
	x, y int
}

type snake struct {
	head, tail point
}

type ladder struct {
	head, tail point
}

type player struct {
	position point
	name     string
}

var entity interface {
	generate_enitity()
	set_points()
	print()
}

// global varibales
var (
	gridSize     int
	numOfPlayers int
	snakes       []snake
	ladders      []ladder
)

// constants
var (
	minGridSize   int = 1
	minSnakeSize      = 2
	minLadderSize     = 2
	minPlayerSize     = 2
	minDiceNum        = 1
	maxDiceNum        = 6
	diceCount         = 1
	player_turn       = 0
	players       []player
)

func pre_setup() {
	fmt.Printf("Enter the size of the grid: ")
	fmt.Scanf("%d", &gridSize)
	if gridSize < 6 {
		fmt.Printf("grid size should be greater than 5")
		fmt.Scanf("%d", &gridSize)
	}
	// take input of number of players and name
	fmt.Printf("Enter the number of the players: ")
	fmt.Scanf("%d", &numOfPlayers)
	get_the_players_name(numOfPlayers)
	// generate snake and ladder
	generate_entities(gridSize)
	fmt.Println("snakes position: ", snakes)
	fmt.Println("ladders position: ", ladders)
	// print positions
	print()
}

func get_the_players_name(numOfPlayers int) error {
	for i := 0; i < numOfPlayers; i++ {
		fmt.Printf("Enter Player %d name: ", i+1)
		var player_name string
		fmt.Scanf("%s", &player_name)
		for _, player_entity := range players {
			if strings.EqualFold(player_entity.name, player_name) {
				return errors.New("no two players can have same name! Give me different names")
			}
		}
		players = append(players, player{position: point{0, 0}, name: player_name})
	}
	return nil
}

func generate_entities(gridSize int) {
	// generate snake
	number_of_snakes := get_random_num(gridSize * 2 / 3)
	for number_of_snakes < minSnakeSize {
		number_of_snakes = get_random_num(gridSize * 2 / 3)
	}
	for i := 0; i < number_of_snakes; i++ {
		snakes = append(snakes, snake{}.generate_enitity(gridSize))
	}

	// generate ladder
	number_of_ladders := get_random_num(gridSize * 2 / 3)
	for number_of_ladders < minLadderSize {
		number_of_ladders = get_random_num(gridSize * 2 / 3)
	}
	for i := 0; i < number_of_snakes; i++ {
		ladders = append(ladders, ladder{}.generate_enitity(gridSize))
	}
}

func (s snake) generate_enitity(num int) snake {
	head := get_point(num)
	tail := get_point(num)
	for (head.x == tail.x) && (head.y == tail.y) && (head.y == (num-1) && head.x == 0) {
		head = get_point(num)
		tail = get_point(num)
	}
	if tail.y > head.y {
		head, tail = tail, head
	}
	s.set_points(head, tail)
	return s
}

func (l ladder) generate_enitity(num int) ladder {
	head := get_point(num)
	tail := get_point(num)
	for (head.x == tail.x) && (head.y == tail.y) && (tail.x == 0) {
		head = get_point(num)
		tail = get_point(num)
	}
	if tail.y > head.y {
		head, tail = tail, head
	}
	l.set_points(head, tail)
	return l
}

func get_point(num int) *point {
	x := get_random_num(num)
	y := get_random_num(num)
	for y == 0 && x == 0 {
		y = get_random_num(num)
		x = get_random_num(num)
	}
	// fmt.Printf("[DEBUG]: Point generated is %d %d\n", x, y)
	return &point{x, y}
}

func get_random_num(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func (s *snake) set_points(head *point, tail *point) {
	s.head = *head
	s.tail = *tail
}

func (l *ladder) set_points(head *point, tail *point) {
	l.head = *head
	l.tail = *tail
}

func play_game() {
	// totalNum := roll_dice(diceCount)
	status := false
	for !status {
		fmt.Printf("Player %s roll your dice [Press '1' to roll]: ", players[player_turn].name)
		// var number int
		var entered_char int8
		fmt.Scanf("%d", &entered_char)
		for entered_char != 1 {
			fmt.Printf("Try again to roll your dice properly [Press '1' to roll]: ")
			fmt.Scanf("%d", &entered_char)
		}
		players[player_turn].roll_dice()
		fmt.Printf("Player %s position: %v\n", players[player_turn].name, players[player_turn].position)
		players[player_turn].hit_by_snake()
		players[player_turn].got_elevated()
		status = players[player_turn].end_game()
		player_turn = (player_turn + 1) % len(players)
	}
	fmt.Println("GAME END!!!")
}

func (p *player) roll_dice() {
	number := get_random_num(6)
	skippedNum := 0
	for number == 0 {
		number = get_random_num(6)
	}
	fmt.Println("Dice: ", number)
	if p.position.y%2 != 0 {
		if p.position.x-number >= 0 {
			p.position.x = p.position.x - number
		} else {
			if p.position.y+1 <= gridSize-1 {
				p.position.y = p.position.y + 1 // y%2 == 0
				skippedNum = p.position.x
				p.position.x = gridSize - (number - skippedNum)
			} else {
				fmt.Println("OOps no move ", p.name, " at position ", p.position)
			}
		}
	} else {
		if p.position.x+number <= gridSize-1 {
			p.position.x = p.position.x + number
		} else {
			p.position.y = p.position.y + 1 // y%2 != 0
			skippedNum = (gridSize - 1) - p.position.x
			p.position.x = number - skippedNum - 1
		}
	}
}

func (p *player) hit_by_snake() {
	for _, snake := range snakes {
		if snake.head.x == p.position.x && snake.head.y == p.position.y {
			fmt.Println("Hit by snake position ", snake.head, " player ", p.name)
			p.position = snake.tail
		}
	}
}

func (p *player) got_elevated() {
	for _, ladder := range ladders {
		if ladder.tail.x == p.position.x && ladder.tail.y == p.position.y {
			fmt.Println("Got "+p.name+" elevented by ladder position ", ladder.head)
			p.position = ladder.head
		}
	}
}

func (p *player) end_game() bool {
	if p.position.x == gridSize-1 && p.position.y == gridSize-1 {
		fmt.Printf("Player %s has won the game\n", p.name)
		return true
	}
	return false
}

func main() {
	pre_setup()
	play_game()
}
