package palette

import (
	"webtyp.com/lang"
	"webtyp.com/model"
	"webtyp.com/router"
	"webtyp.com/view"
)

// Item renders a palette colour row with its own live hex swatch on the badge.
func (m *PaletteColor) Item() model.Item {
	return model.Item{
		ID:          m.Id,
		Label:       m.Name,
		Description: m.Hex,
		Tone:        m.Hex,
	}
}

// NewView builds the Presenter for managing the database colour palette.
func NewView(caller router.Caller, opts ...view.Option) view.Presenter {
	b := view.NewCallerLister(view.NewRecords().Caller(caller),
		view.Ops{Module: ModelName, Page: view.Page{First: view.All}},
		func() model.TableSlice { return &PaletteColorList{} })
	allOpts := append([]view.Option{view.WithTitle(lang.Translate("Colors").String())}, opts...)
	return view.New(b, &PaletteColor{}, allOpts...)
}
