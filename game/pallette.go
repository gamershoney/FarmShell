package game

import lipgloss "charm.land/lipgloss/v2"

var LightGreen = lipgloss.Color("#34eb46")
var DarkGreen = lipgloss.Color("#114516")
var DarkYellow = lipgloss.Color("#98a808")
var White = lipgloss.White
var Black = lipgloss.Black
var Brown = lipgloss.Color("#594100")
var TextBlue = lipgloss.Color("#00c3ff")
var FancyRed = lipgloss.Color("#ff0040")
var FancyOrange = lipgloss.Color("#ff8c00")
var FancyGray = lipgloss.Color("#62787F")
var textures = map[BlockName]Appearance{
	Wheat: {
		FG:   White,
		BG:   Brown,
		Char: 'î',
	},
	Grass: {
		FG:   lipgloss.NoColor{},
		BG:   DarkGreen,
		Char: '\x20',
	},
	PlayerBlock: {
		FG:   White,
		BG:   nil,
		Char: '@',
	},
}
