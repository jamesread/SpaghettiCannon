package store

import (
	"context"
	"database/sql"
	"errors"
)

type Store interface {
	LatestMigration(ctx context.Context) (string, error)
	HasMigration(ctx context.Context, id string) (bool, error)
	ListCvars(ctx context.Context) ([]CvarRow, error)
	FindCvar(ctx context.Context, key string) (*CvarRow, error)
	InsertCvarIfMissing(ctx context.Context, row CvarRow) error
	UpdateCvar(ctx context.Context, key string, valueInt int, valueString string) error
}

type CvarRow struct {
	Key         string
	MainType    string
	Title       string
	Description string
	Category    string
	ValueString string
	Ordinal     int
	ValueInt    int
}

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (m *MySQL) LatestMigration(ctx context.Context) (string, error) {
	var id sql.NullString
	err := m.db.QueryRowContext(ctx,
		`SELECT id FROM migrations ORDER BY applied_at DESC, id DESC LIMIT 1`).Scan(&id)
	if err != nil {
		return "", err
	}
	if !id.Valid {
		return "", nil
	}
	return id.String, nil
}

func (m *MySQL) HasMigration(ctx context.Context, id string) (bool, error) {
	var n int
	err := m.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM migrations WHERE id = ?`, id).Scan(&n)
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func cvarSelectSQL() string {
	return `SELECT cvar_key, COALESCE(cvar_value_int, 0), COALESCE(cvar_value_string, ''), cvar_main_type,
		COALESCE(cvar_title, ''), COALESCE(cvar_description, ''), COALESCE(cvar_category, ''), COALESCE(cvar_ordinal, 0)
		FROM cvars`
}

func scanCvar(s interface{ Scan(...any) error }) (*CvarRow, error) {
	var row CvarRow
	if err := s.Scan(&row.Key, &row.ValueInt, &row.ValueString, &row.MainType, &row.Title, &row.Description,
		&row.Category, &row.Ordinal); err != nil {
		return nil, err
	}
	return &row, nil
}

func (m *MySQL) ListCvars(ctx context.Context) ([]CvarRow, error) {
	rows, err := m.db.QueryContext(ctx, cvarSelectSQL()+" ORDER BY cvar_ordinal, cvar_key")
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	var out []CvarRow
	for rows.Next() {
		row, err := scanCvar(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, *row)
	}
	return out, nil
}

func (m *MySQL) FindCvar(ctx context.Context, key string) (*CvarRow, error) {
	row := m.db.QueryRowContext(ctx, cvarSelectSQL()+" WHERE cvar_key = ?", key)
	c, err := scanCvar(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return c, err
}

func (m *MySQL) InsertCvarIfMissing(ctx context.Context, row CvarRow) error {
	_, err := m.db.ExecContext(ctx,
		`INSERT INTO cvars (cvar_key, cvar_value_int, cvar_value_string, cvar_main_type, cvar_title, cvar_description, cvar_category, cvar_ordinal)
		 VALUES (?, ?, NULLIF(?, ''), ?, ?, ?, ?, ?)
		 ON DUPLICATE KEY UPDATE
		   cvar_title = VALUES(cvar_title),
		   cvar_description = VALUES(cvar_description),
		   cvar_category = VALUES(cvar_category),
		   cvar_ordinal = VALUES(cvar_ordinal)`,
		row.Key, row.ValueInt, row.ValueString, row.MainType, row.Title, row.Description, row.Category, row.Ordinal)
	return err
}

func (m *MySQL) UpdateCvar(ctx context.Context, key string, valueInt int, valueString string) error {
	_, err := m.db.ExecContext(ctx,
		`UPDATE cvars SET cvar_value_int = ?, cvar_value_string = NULLIF(?, '') WHERE cvar_key = ? LIMIT 1`,
		valueInt, valueString, key)
	return err
}
