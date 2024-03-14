package transaksi

import (
	"database/sql"
	"fmt"
)

// TransaksiRepository provides CRUD operations for Transaksi.
type TransaksiRepository interface {
	CreateTransaksi(transaksi *Transaksi) error
	GetTransaksiByID(id int) (*Transaksi, error)
	UpdateTransaksi(transaksi *Transaksi) error
	DeleteTransaksi(id int) error
	GetAllTransaksi() ([]*Transaksi, error)
}

// MySQLTransaksiRepo is a MySQL implementation of TransaksiRepository.
type MySQLTransaksiRepo struct {
	DB *sql.DB
}

// NewMySQLTransaksiRepo initializes a new instance of MySQLTransaksiRepo.
func NewMySQLTransaksiRepo(db *sql.DB) *MySQLTransaksiRepo {
	return &MySQLTransaksiRepo{DB: db}
}

// CreateTransaksi creates a new transaksi.
func (r *MySQLTransaksiRepo) CreateTransaksi(transaksi *Transaksi) error {
	_, err := r.DB.Exec("INSERT INTO transaksi (barang_id, quantity, total) VALUES (?, ?, ?)", transaksi.BarangID, transaksi.Quantity, transaksi.Total)
	if err != nil {
		return fmt.Errorf("failed to create transaksi: %v", err)
	}
	return nil
}

// GetAllTransaksi retrieves all transaksis from the database.
func (r *MySQLTransaksiRepo) GetAllTransaksi() ([]*Transaksi, error) {
	rows, err := r.DB.Query("SELECT id, barang_id, quantity, total FROM transaksi")
	if err != nil {
		return nil, fmt.Errorf("failed to get all transaksi: %v", err)
	}
	defer rows.Close()

	var transaksis []*Transaksi
	for rows.Next() {
		var t Transaksi
		err := rows.Scan(&t.ID, &t.BarangID, &t.Quantity, &t.Total)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaksi: %v", err)
		}
		transaksis = append(transaksis, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration of rows: %v", err)
	}

	return transaksis, nil
}

// GetTransaksiByID retrieves a transaksi by its ID.
func (r *MySQLTransaksiRepo) GetTransaksiByID(id int) (*Transaksi, error) {
	var t Transaksi
	err := r.DB.QueryRow("SELECT id, barang_id, quantity, total FROM transaksi WHERE id = ?", id).Scan(&t.ID, &t.BarangID, &t.Quantity, &t.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get transaksi: %v", err)
	}
	return &t, nil
}

// UpdateTransaksi updates an existing transaksi.
func (r *MySQLTransaksiRepo) UpdateTransaksi(transaksi *Transaksi) error {
	_, err := r.DB.Exec("UPDATE transaksi SET barang_id = ?, quantity = ?, total = ? WHERE id = ?", transaksi.BarangID, transaksi.Quantity, transaksi.Total, transaksi.ID)
	if err != nil {
		return fmt.Errorf("failed to update transaksi: %v", err)
	}
	return nil
}

// DeleteTransaksi deletes a transaksi by its ID.
func (r *MySQLTransaksiRepo) DeleteTransaksi(id int) error {
	_, err := r.DB.Exec("DELETE FROM transaksi WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete transaksi: %v", err)
	}
	return nil
}
