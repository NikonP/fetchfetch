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
	"fmt"
	"os"

	"github.com/NikonP/fetchfetch/fetch"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(80)
	t.SetStyle(table.StyleColoredGreenWhiteOnBlack)

	t.AppendHeader(table.Row{
		"Fetch",
		"Description",
		"Is installed",
	})

	installedCount := 0
	totalFetchesCount := len(fetch.Fetches)

	for _, fetch := range fetch.Fetches {
		isInstalled := ""

		if fetch.FetchExists() {
			installedCount += 1
			isInstalled = text.FgGreen.Sprint("[+]")
			t.AppendRow(table.Row{
				fetch.Name,
				fetch.Description,
				isInstalled,
			})
		}

	}

	t.AppendFooter(table.Row{
		"",
		"",
		fmt.Sprintf(
			"%d/%d",
			installedCount,
			totalFetchesCount,
		),
	})

	t.AppendSeparator()

	t.Render()

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedRowLength(80)
	t.SetStyle(table.StyleColoredGreenWhiteOnBlack)
}
