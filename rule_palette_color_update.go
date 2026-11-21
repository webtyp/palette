package palette

import (
	"webtyp.com/crud"
	"webtyp.com/model"
)

func (next *PaletteColor) Update(a crud.CurrentUser, tx crud.Tx, old *PaletteColor, fields []string, m *Module) (crud.RuleAnswer, error) {
	if next.Hex != "" {
		next.Hex = normalizeHex(next.Hex)
	}
	dup := false
	err := tx.ReadAll(func() model.Model { return &PaletteColor{} }, crud.Query{}, func(row model.Model) {
		r := row.(*PaletteColor)
		if r.Name == next.Name && r.Id != old.Id {
			dup = true
		}
	})
	if err != nil {
		return crud.RuleAnswer{}, err
	}
	if dup {
		return crud.RuleAnswer{}, ErrDuplicateName
	}
	return crud.ApplyAndPublish(), nil
}
