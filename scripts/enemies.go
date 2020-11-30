package scripts

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/hajimehoshi/ebiten"
)

type Enemies struct {
	game     *Game
	long     int
	dir      string
	points   int
	parts    [][]float64
	behavior chan int
	head     ebiten.Image
	body     ebiten.Image
	src      rand.Source
}

func newEnemie(g *Game) *Enemies {
	e := Enemies{
		game:   g,
		long:   0,
		dir:    "left",
		points: 0,
	}
	e.behavior = make(chan int)
	e.src = rand.NewSource(time.Now().UnixNano())
	rndm := rand.New(e.src)
	partX := float64(rndm.Intn(30) * 20)
	partY := float64(rndm.Intn(30) * 20)
	e.parts = append(e.parts, []float64{partX, partY})
	headImg, _, _ := ebitenutil.NewImageFromFile("Resources/enemyhead.png")
	bodyImg, _, _ := ebitenutil.NewImageFromFile("Resources/enemybody.png")
	e.head = *headImg
	e.body = *bodyImg
	return &e
}

func (e *Enemies) Direction(dotTime int) error {
	if dotTime == 1 {
		rndm := rand.New(e.src)

		action := rndm.Intn(4)
		changingDirection := rndm.Intn(3)
		posX, posY := e.getHeadPos() //position of the snakes head
		if changingDirection == 0 {
			switch action { //checks the boundings of the map
			case 0:
				if posX < 560 && e.dir != "left" {
					e.dir = "right"
				} else {
					e.dir = "left"
				}
				return nil
			case 1:
				if posY < 560 && e.dir != "up" {
					e.dir = "down"
				} else {
					e.dir = "up"
				}
				return nil
			case 2:
				if posY > 20 && e.dir != "down" {
					e.dir = "up"
				} else {
					e.dir = "down"
				}
				return nil
			case 3:
				if posX > 20 && e.dir != "right" {
					e.dir = "left"
				} else {
					e.dir = "right"
				}
				return nil
			}
		}
		if posX >= 560 { //moves the enemy to avoid getting out of bounds
			e.dir = "left"
			return nil
		}
		if posX == 20 {
			e.dir = "right"
			return nil
		}
		if posY == 560 {
			e.dir = "up"
			return nil
		}
		if posY == 20 {
			e.dir = "down"
			return nil
		}
	}

	if dotTime == 1 { //checks collision between enemy and player
		xPos, yPos := e.game.player.getHeadPos()
		if e.collisionWithPlayer(xPos, yPos) {
			//e.game.End()
		}
	}
	return nil

}

func (e *Enemies) Behavior() error {
	for {
		dotTime := <-e.behavior
		e.Direction(dotTime)
	}
	return nil
}

func (e *Enemies) getHeadPos() (float64, float64) {
	return e.parts[0][0], e.parts[0][1]
}

func (e *Enemies) getBodyPos(x int) (float64, float64) {
	return e.parts[x+1][0], e.parts[x+1][1]
}

func (e *Enemies) moveHead(xPos, yPos float64) {
	newX := e.parts[0][0] + xPos
	newY := e.parts[0][1] + yPos
	e.parts = append([][]float64{[]float64{newX, newY}}, e.parts...)
	e.parts = e.parts[:e.long+1]
}

func (e *Enemies) collisionWithPlayer(xCoor, yCoor float64) bool {
	for i := 1; i < len(e.parts); i++ {
		if xCoor == e.parts[i][0] && yCoor == e.parts[i][1] {
			return true
		}
	}
	return false
}
