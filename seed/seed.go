package seed

import (
	"webtyp.com/crud"
	"webtyp.com/model"
	"webtyp.com/palette"
)

// StarterColor is one initial row seeded into an empty palette_color table
// so a new installation never starts blank; every row remains editable/deletable.
type StarterColor struct {
	ID   string
	Name string
	Hex  string
}

// Defaults returns the initial theme-harmonized starter palette (led by Go cyan
// and WebAssembly violet). Lives in seed/ so it never ships in the app binary.
func Defaults() []StarterColor {
	return []StarterColor{
		{ID: "cyan", Name: "Celeste", Hex: "#00ADD8"},
		{ID: "violet", Name: "Violeta", Hex: "#654FF0"},
		{ID: "teal", Name: "Turquesa", Hex: "#0D9488"},
		{ID: "blue", Name: "Azul", Hex: "#2563EB"},
		{ID: "indigo", Name: "Índigo", Hex: "#4F46E5"},
		{ID: "emerald", Name: "Esmeralda", Hex: "#16A34A"},
		{ID: "amber", Name: "Ámbar", Hex: "#D97706"},
		{ID: "coral", Name: "Coral", Hex: "#EA580C"},
		{ID: "rose", Name: "Rosa", Hex: "#DB2777"},
		{ID: "slate", Name: "Pizarra", Hex: "#475569"},
	}
}

// Load seeds the starter palette rows only when this tenant has no colours yet.
// Idempotent: safe to run repeatedly; never overwrites user edits or deletions
// once at least one colour exists.
func Load(m *palette.Module) error {
	store := m.StorePaletteColor()
	count := 0
	if err := store.ReadAll(func() model.Model { return &palette.PaletteColor{} }, crud.Query{Limit: 1}, func(model.Model) {
		count++
	}); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	for _, c := range Defaults() {
		rec := &palette.PaletteColor{
			Id:       c.ID,
			TenantId: m.TenantID(),
			Name:     c.Name,
			Hex:      c.Hex,
		}
		if err := store.Create(rec); err != nil {
			return err
		}
	}
	return nil
}
