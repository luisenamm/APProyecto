package scripts

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//What snakes eat
type Food struct {
	game   *Game
	img    ebiten.Image
	xLimit int
	yLimit int
	xCoord float64
	yCoord float64
	eaten  bool
}

//Creates Food
func newFood(g *Game) *Food {
	f := Food{
		game:   g,
		xLimit: 30,
		yLimit: 30,
		eaten:  false,
	}
	food, _, _ := ebitenutil.NewImageFromFile("Resources/food.png")
	f.img = *food
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	f.xCoord = float64(random.Intn(f.xLimit) * 20)
	f.yCoord = float64(random.Intn(f.yLimit) * 20)

	return &f
}

func (f *Food) Eaten(dotTime int) error {
	if f.eaten == false {
		return nil //If food has been eaten
	}
	return nil
}

//Draw the Food
func (f *Food) Draw(screen *ebiten.Image, dotTime int) error {
	charDO = &ebiten.DrawImageOptions{}
	charDO.GeoM.Translate(f.xCoord, f.yCoord)
	screen.DrawImage(&f.img, charDO) //Every Cherry is drawn here
	return nil
}
