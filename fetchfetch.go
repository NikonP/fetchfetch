package main

// WARNING: The code that follows may make you cry:
//           A Safety Pig has been provided below for your benefit
//                              _
//      _._ _..._ .-',     _.._(`))
//     '-. `     '  /-._.-'    ',/
//       )         \            '.
//      / _    _    |             \
//     |  a    a    /              |
//      \   .-.                     ;
//       '-('' ).-'       ,'       ;
//          '-;           |      .'
//            \           \    /
//            | 7  .__  _.-\   \
//            | |  |  ``/  /`  /
//           /,_|  |   /,_/   /
//              /,_/      '`-'

import (
	"os/exec"

	"github.com/fatih/color"
)

var fetches = []string{
	"fetchfetch",
	"neofetch",
	"pfetch",
	"ramfetch",
	"ufetch",
	"nerdfetch",
	"archfetch",
	"cfetch",
	"shutthefetchup",
	"paleofetch",
}

func fetchExists(fetch string) bool {
	_, err := exec.LookPath(fetch)
	return err == nil
}

func main() {
	headerStyle := color.New(color.FgGreen)
	headerStyle.Add(color.Bold)

	fetchStyle := color.New(color.BgCyan)
	fetchStyle.Add(color.FgBlack)

	isInstalledStyle := color.New(color.FgGreen)
	isNotInstalledStyle := color.New(color.FgRed)

	cyan := color.New(color.FgCyan)

	omgStyle := color.New(color.FgRed)
	omgStyle.Add(color.BlinkSlow)

	headerStyle.Print("### FetchFetch ###\n\n")

	installedCount := 0

	for _, fetch := range fetches {
		fetchStyle.Printf("%20s", fetch)

		if fetchExists(fetch) {
			isInstalledStyle.Println(" is installed")
			installedCount += 1
		} else {
			isNotInstalledStyle.Println(" is not installed")
		}
	}

	cyan.Printf("\nFetches installed: %d\n", installedCount)
	if installedCount == 0 {
		omgStyle.Println("OMG no fetches installed!!!")
	}
}
