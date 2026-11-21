package palette

import (
	"webtyp.com/crud"
	"webtyp.com/model"
)

func normalizeHex(h string) string {
	b := []byte(h)
	for i := range b {
		if b[i] >= 'a' && b[i] <= 'f' {
			b[i] -= 'a' - 'A'
		}
	}
	return string(b)
}

func (p *PaletteColor) Create(a crud.CurrentUser, tx crud.Tx, m *Module) (crud.RuleAnswer, error) {
	p.Hex = normalizeHex(p.Hex)
	dup := false
	err := tx.ReadAll(func() model.Model { return &PaletteColor{} }, crud.Query{}, func(row model.Model) {
		if row.(*PaletteColor).Name == p.Name {
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
