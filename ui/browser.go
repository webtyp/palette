package ui

import (
	"webtyp.com/model"
	pal "webtyp.com/palette"
	"webtyp.com/view"
)

// Screens declares the Colors management screen and wires live CSS custom-property
// synchronization so edits in the database take effect immediately without recompiling.
func Screens(d view.BrowserDeps) []view.Screen {
	pres := pal.NewView(d.Caller, view.WithDescription(descColors))
	pres.OnChange(func() {
		applyPaletteCSS(pres.Items())
	})

	return []view.Screen{{
		ID:         ID,
		Label:      Label,
		Icon:       ID,
		Permission: model.ResourceOf(&pal.PaletteColorModel),
		Presenter:  pres,
	}}
}
