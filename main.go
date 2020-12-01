package main

import (
	"fmt"
	"log"
	"os"
	"scripts/scripts"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

var gm scripts.Game
var cherryN int
var enemiesN int

func init() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong number of parameters")
		os.Exit(3)
	}

	var err error
	cherryN, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Cherry must be a number")
		os.Exit(3)
	}

	enemiesN, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Enemies must be number")
		os.Exit(3)
	}

	if cherryN == 0 || enemiesN == 0 {
		fmt.Println("At least one cherry and one enemy")
		os.Exit(3)
	}

	gm = scripts.NewGame(cherryN, enemiesN)
}

// Game ebiten, code from api
type Game struct {
}

// Update the thread, code from api
func (g *Game) Update(screen *ebiten.Image) error {
	if err := gm.Update(); err != nil {
		return err
	}
	return nil
}

// Draw the image, code from api
func (g *Game) Draw(screen *ebiten.Image) {
	if err := gm.Draw(screen); err != nil {
		fmt.Println(err)
	}
}

// Layout : Function which executes when it needs to reajust, code from api
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Gosnakes")

	//code from api
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
