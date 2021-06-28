package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/danielfmelo/travel_finder/services"
)

const exitCommand = "exit"

type Command struct {
	finderService services.Finder
}

func (c *Command) Start() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("please enter the route: (example: GRU-FLN) \n")
		scanner.Scan()
		text := scanner.Text()
		if isExitCmd(text) {
			fmt.Println("exiting from command interface...")
			break
		}
		ori, dest, err := extractOriDest(text)
		if err != nil {
			fmt.Println("error to validate route format!")
			continue
		}
		cheapestRoute, err := c.finderService.GetSmallestPriceAndRoute(ori, dest)
		if err != nil {
			fmt.Println("error to find best price!")
			continue
		}
		fmt.Printf("%s > $%d\n", cheapestRoute.Path, cheapestRoute.Value)
	}
}

func extractOriDest(text string) (string, string, error) {
	splited := strings.Split(text, "-")
	if len(splited) != 2 {
		return "", "", errors.New("error to validate input")
	}
	return splited[0], splited[1], nil
}

func isExitCmd(cmd string) bool {
	if cmd == exitCommand {
		return true
	}
	return false
}

func New(finderService services.Finder) *Command {
	return &Command{
		finderService: finderService,
	}
}
