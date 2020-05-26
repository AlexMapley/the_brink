package console

func NewTownConsole() Console {
	console := Console{}

	// set default options
	actions := make([]string, 5)
	actions[0] = "Character"
	actions[1] = "Map"
	actions[2] = "Patrol the town"
	actions[3] = "Rest"
	actions[4] = "Become Paladin"

	console.Actions = actions

	return console
}
