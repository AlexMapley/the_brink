package map

// CreateMap will return a blank grid
func CreateMap(xMax int, yMax int) map[Tile]rune {
	grid := map[Tile]rune{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			grid[tile] = '.'
		}
	}
	return grid
}

// UpdateMap will set the position and rune for all existent parties
func (world *World) UpdateMap() {

	// Regerate map from scratch
	// TODO: Optimize me
	grid := map[Tile]rune{}
	for y := 0; y < map.YMax; y++ {
		for x := 0; x < map.XMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			grid[tile] = '.'
		}
	}

	// Move parties
	for _, party := range map.Parties {
		coordinates := Tile {
			X: party.X,
			Y: party.Y,
		}
		grid[coordinates] = party.Rune
	}

	map.Tiles = grid

}