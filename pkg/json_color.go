package pkg

import (
	"encoding/json"
	"os"

	"github.com/mattn/go-colorable"
	jsonColor "github.com/neilotoole/jsoncolor"
)

func PrintJSON(value any, isRaw bool) error {
	if isRaw {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		return enc.Encode(value)
	}
	out := colorable.NewColorable(os.Stdout)
	enc := jsonColor.NewEncoder(out)
	enc.SetIndent("", " ")

	clrs := jsonColor.DefaultColors()
	clrs.Bool = jsonColor.Color("\x1b[95m")
	clrs.Number = jsonColor.Color("\x1b[95m")
	clrs.Key = jsonColor.Color("\x1b[96m")
	clrs.String = jsonColor.Color("\x1b[92m")

	enc.SetColors(clrs)

	return enc.Encode(value)
}
