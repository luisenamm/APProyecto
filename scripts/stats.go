package scripts

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

type Stats struct {
	game      *Game
	points    int
	maxPoints int
	cherrys   int
	highScore int
}

func newStats(g *Game, max int) *Stats {
	s := Stats{
		game:      g,
		points:    0,
		maxPoints: max,
		highScore: 0,
	}
	return &s
}

func textWriter(text string) (w int, h int) {
	return 7 * len(text), 13
}

//Shows the final results
func (s *Stats) endGameScreen(screen *ebiten.Image) {
	if s.cherrys != s.maxPoints {
		textToShow := "Game Over"
		textWeight, textHeight := textWriter(textToShow)
		screenWeight := screen.Bounds().Dx
		screenHeight := screen.Bounds().Dy
		text.Draw(screen, textToShow, basicfont.Face7x13, screenWeight/2-textWeight/2, screenHeight/2+textHeight/2, color.White)
	} else if s.points == s.maxPoints {
		textToShow := "Congratulations! You Win!"
		textWeight, textHeight := textWriter(textToShow)
		screenWeight := screen.Bounds().Dx
		screenHeight := screen.Bounds().Dy
		text.Draw(screen, textToShow, basicfont.Face7x13, screenWeight/2-textWeight/2, screenHeight/2+textHeight/2, color.White)
	} else {
		textToShow := "Game Over"
		textWeight, textHeight := textWriter(textToShow)
		screenWeight := screen.Bounds().Dx
		screenHeight := screen.Bounds().Dy
		text.Draw(screen, textToShow, basicfont.Face7x13, screenWeight/2-textWeight/2, screenHeight/2+textHeight/2, color.White)
	}
}

func (s *Stats) Draw(screen *ebiten.Image) error {
	text.Draw(screen, "Score: "+strconv.Itoa(s.points), basicfont.Face7x13, 20, 20, color.White)
	if !s.game.playing {
		cherrys := 0
		max := 0
		for i := 0; i < len(s.game.enemies); i++ {
			cherrys += h.game.enemies[i].points
			if max < s.game.enemies[i].points {
				max = s.game.enemies[i].points
			}
		}
		cherrys += s.game.player.points
		if max < s.game.player.points {
			max = s.game.player.points
		}
		s.highScore = max
		s.cherrys = cherrys
		s.endGameScreen(screen)
	}

	return nil
}

func (s *Stats) secondEnd(screen *ebiten.Image) {
	cherrys := 0
	max := 0
	for i := 0; i < len(s.game.enemies); i++ {
		cherrys += s.game.enemies[i].points
		if max < s.game.enemies[i].points {
			max = s.game.enemies[i].points
		}
	}
	cherrys += s.game.player.points
	if max < s.game.player.points {
		max = s.game.player.points
	}
	s.highScore = max
	s.cherrys = cherrys
	s.endGameScreen(screen)
}
