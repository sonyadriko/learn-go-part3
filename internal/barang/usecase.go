package barang

import (
	"errors"
)

// BarangUseCase provides use cases for Barang.
type BarangUseCase struct {
	Repo BarangRepository
}

// NewBarangUseCase initializes a new instance of BarangUseCase.
func NewBarangUseCase(repo BarangRepository) *BarangUseCase {
	return &BarangUseCase{
		Repo: repo,
	}
}

// GetAllBarang retrieves all barangs.
func (uc *BarangUseCase) GetAllBarang() ([]*Barang, error) {
	barangs, err := uc.Repo.GetAllBarang()
	if err != nil {
		return nil, err
	}
	return barangs, nil
}

// CreateBarang creates a new barang.
func (uc *BarangUseCase) CreateBarang(barang *Barang) error {
	return uc.Repo.CreateBarang(barang)
}

// GetBarangByID retrieves a barang by its ID.
func (uc *BarangUseCase) GetBarangByID(id int) (*Barang, error) {
	b, err := uc.Repo.GetBarangByID(id)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// UpdateBarang updates an existing barang.
func (uc *BarangUseCase) UpdateBarang(barang *Barang) error {
	if _, err := uc.Repo.GetBarangByID(barang.ID); err != nil {
		return err
	}
	return uc.Repo.UpdateBarang(barang)
}

// DeleteBarang deletes a barang by its ID.
func (uc *BarangUseCase) DeleteBarang(id int) error {
	if _, err := uc.Repo.GetBarangByID(id); err != nil {
		return err
	}
	return uc.Repo.DeleteBarang(id)
}

// ErrNotFound indicates that the requested resource was not found.
var ErrNotFound = errors.New("not found")
