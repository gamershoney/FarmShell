package game

func InitGameLoop() {

}

type Crop struct {
	Name   string  `json:"crop_name"`
	Amount int     `json:"amount"`
	Value  float64 `json:"crop_value"`
}

type Chunk struct {
}

type model struct {
	Playername string           `json:"username"`
	Crops      map[string]*Crop `json:"player_crops"`
}

func gameModel() model {

}
