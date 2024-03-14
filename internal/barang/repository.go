package barang

import (
	"database/sql"
	"fmt"
)

// BarangRepository provides CRUD operations for Barang.
type BarangRepository interface {
	CreateBarang(barang *Barang) error
	GetBarangByID(id int) (*Barang, error)
	UpdateBarang(barang *Barang) error
	DeleteBarang(id int) error
	GetAllBarang() ([]*Barang, error)
}

// MySQLBarangRepo is a MySQL implementation of BarangRepository.
type MySQLBarangRepo struct {
	DB *sql.DB
}

// NewMySQLBarangRepo initializes a new instance of MySQLBarangRepo.
func NewMySQLBarangRepo(db *sql.DB) *MySQLBarangRepo {
	return &MySQLBarangRepo{DB: db}
}

// CreateBarang creates a new barang.
func (r *MySQLBarangRepo) CreateBarang(barang *Barang) error {
	_, err := r.DB.Exec("INSERT INTO barang (name, price) VALUES (?, ?)", barang.Name, barang.Price)
	if err != nil {
		return fmt.Errorf("failed to create barang: %v", err)
	}
	return nil
}

// GetAllBarang retrieves all barangs from the database.
func (r *MySQLBarangRepo) GetAllBarang() ([]*Barang, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM barang")
	if err != nil {
		return nil, fmt.Errorf("failed to get all barang: %v", err)
	}
	defer rows.Close()

	var barangs []*Barang
	for rows.Next() {
		var b Barang
		err := rows.Scan(&b.ID, &b.Name, &b.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan barang: %v", err)
		}
		barangs = append(barangs, &b)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration of rows: %v", err)
	}

	return barangs, nil
}

// GetBarangByID retrieves a barang by its ID.
func (r *MySQLBarangRepo) GetBarangByID(id int) (*Barang, error) {
	var b Barang
	err := r.DB.QueryRow("SELECT id, name, price FROM barang WHERE id = ?", id).Scan(&b.ID, &b.Name, &b.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get barang: %v", err)
	}
	return &b, nil
}

// UpdateBarang updates an existing barang.
func (r *MySQLBarangRepo) UpdateBarang(barang *Barang) error {
	_, err := r.DB.Exec("UPDATE barang SET name = ?, price = ? WHERE id = ?", barang.Name, barang.Price, barang.ID)
	if err != nil {
		return fmt.Errorf("failed to update barang: %v", err)
	}
	return nil
}

// DeleteBarang deletes a barang by its ID.
func (r *MySQLBarangRepo) DeleteBarang(id int) error {
	_, err := r.DB.Exec("DELETE FROM barang WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete barang: %v", err)
	}
	return nil
}
