package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"the_brink/characters"
	"the_brink/console"
	"the_brink/world"

	"github.com/fatih/color"
)

var DayCounter int
var trim string = "-----------------------------------------\n"
var player characters.Player


func main() {
	metaGame := world.MetaGame{
		Day: 1,
	}


	// Create Character
	color.Cyan("What is your name?\n")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	// Choose Character
	color.Cyan("\nWhat Class Do You Pick?\n")
	classConsole := console.NewClassConsole()
	for len(player.Character.Stats.Class) == 0 {
		option := classConsole.ChooseAction()
		if option > 0 && option <= len(classConsole.Actions) {
			switch classConsole.Actions[option-1] {
			case "Rogue":
				player = characters.NewPlayer(name, "rogue")
				break
			case "Warrior":
				player = characters.NewPlayer(name, "warrior")
				break
			case "Wizard":
				player = characters.NewPlayer(name, "wizard")
				break
			default:
				continue
			}
		}
	}

	// main game loop
	for (player.Character.Stats.Health > 0) {
		townConsole := console.NewTownConsole()

		color.Green("%sDay %d in town, what do you?\n%s", trim, metaGame.Day, trim)

		dayLoop:
		for {
			option := townConsole.ChooseAction()

			if option > 0 && option <= len(townConsole.Actions) {
				color.Green("You have chosen option %d, %s", option, townConsole.Actions[option-1])

				switch townConsole.Actions[option-1] {
				
				// Stats
				case "Stats":
					player.Character.Stats.Display()
				
				// Inventory
				case "Inventory":
					player.Inventory.Display()
				
				// Fight
				case "Patrol the town":

					if (DayCounter % 2 == 0) {
						fmt.Println("\n\nA strange bandit appears")
						bandit := characters.NewBandit("Mel", player.Character.Stats.Level)

						player.Character.Duel(&bandit.Character)
					
						// loot bandit if won
						if (player.Character.Stats.Health > 0) {
							player.Inventory.Loot(&enemy.Inventory)
						}
					} else {
						fmt.Println("\n\nAn agry thug appears")
						thug := characters.NewThug("Dougy", player.Character.Stats.Level)

						player.Character.Duel(&thug.Character)
					
						// loot thug if won
						if (player.Character.Stats.Health > 0) {
							player.Inventory.Loot(&enemy.Inventory)
						}
					}

					break dayLoop

				// Rest
				case "Rest":
					player.Character.Rest()
					fmt.Println("Your stats have been restored")
					break dayLoop

				// Level Up
				case "Level Up":
					// level up player and bandit
					player.Character.LevelUp()
					player.Character.Rest()
					break dayLoop
				}
			}
		}

		// Day ends
		metaGame.Day++
	}

	color.Cyan("\n\nGame Over %s, Day %d\n\n\n", player.Character.Stats.Name, metaGame.Day)
	fmt.Println("Your Stats:")
	player.Character.Stats.Display()
	color.Cyan("\n\nOne day later (Day %d), %s is dead.\n\n\n", metaGame.Day, player.Character.Stats.Name)
}
