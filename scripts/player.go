package scripts

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var charDO *ebiten.DrawImageOptions

//Player
type Player struct {
	game     *Game
	long     int
	dir      string
	points   int
	parts    [][]float64
	behavior chan int
	head     ebiten.Image
	body     ebiten.Image
}

func newPlayer(g *Game) *Player {
	p := Player{
		game:   g,
		long:   0,
		dir:    "up",
		points: 0,
	}
	p.behavior = make(chan int)
	p.parts = append(p.parts, []float64{300, 300})
	headImg, _, _ := ebitenutil.NewImageFromFile("Resources/playerhead.png")
	bodyImg, _, _ := ebitenutil.NewImageFromFile("Resources/playerbody.png")
	p.head = *headImg
	p.body = *bodyImg

	return &p
}

func (p *Player) Direction(dotTime int) error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.dir = "up"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.dir = "down"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.dir = "right"
		return nil
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.dir = "left"
		return nil
	}
	if dotTime == 1 {
		xPos, yPos := p.getHeadPos()
		if xPos < 0 || xPos > 580 || yPos < 0 || yPos > 580 || p.collisionWithHimself() {
			//s.game.End()
		}
	}
	return nil
}

func (p *Player) Move(dotTime int) {
	if dotTime == 1 {
		switch p.dir {
		case "up":
			p.moveHead(0, -20)
		case "down":
			p.moveHead(0, +20)
		case "right":
			p.moveHead(+20, 0)
		case "left":
			p.moveHead(-20, 0)
		}
	}
}

func (p *Player) getHeadPos() (float64, float64) {
	return p.parts[0][0], p.parts[0][1]
}
func (p *Player) moveHead(xPos, yPos float64) {
	newX := p.parts[0][0] + xPos
	newY := p.parts[0][1] + yPos
	p.parts = append([][]float64{[]float64{newX, newY}}, p.parts...)
	p.parts = p.parts[:p.long+1]
}

func (p *Player) collisionWithHimself() bool {
	xCoor, yCoor := p.getHeadPos()
	for i := 1; i < len(p.parts); i++ {
		if xCoor == p.parts[i][0] && yCoor == p.parts[i][1] {
			return true
		}
	}
	return false
}
