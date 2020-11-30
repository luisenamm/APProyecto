package scripts

//The main scene
type Game struct {
	player    *Player
	snakeChan chan int
	stats         *Stats
	food        []*Food
	numCherries int
	numEnemies  int
	enemies     []*Enemies
	enemiesChan []chan int
	playing     bool
	points      int
	dotTime     int
}
