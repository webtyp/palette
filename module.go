package palette

import (
	"webtyp.com/crud"
	"webtyp.com/model"
	"webtyp.com/orm"
)

// Module owns the database-backed colour palette for one tenant.
type Module struct {
	db       *orm.DB
	ids      model.IDGenerator
	tenantID string
}

// New validates dependencies and creates the palette module.
func New(db *orm.DB, ids model.IDGenerator, tenantID string) (*Module, error) {
	switch {
	case db == nil:
		return nil, ErrNoDB
	case ids == nil:
		return nil, ErrNoIDs
	case tenantID == "":
		return nil, ErrNoTenant
	}
	return &Module{db: db, ids: ids, tenantID: tenantID}, nil
}

// StorePaletteColor returns the tenant-scoped CRUD store for palette_color.
func (m *Module) StorePaletteColor() crud.Store {
	return orm.TenantStore(m.db, m.tenantID, m.ids)
}

// TenantID returns the tenant this module is scoped to.
func (m *Module) TenantID() string {
	return m.tenantID
}
