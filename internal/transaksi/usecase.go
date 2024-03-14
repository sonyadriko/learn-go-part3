package transaksi

import (
	"errors"
)

// TransaksiUseCase provides use cases for Transaksi.
type TransaksiUseCase struct {
	Repo TransaksiRepository
}

// NewTransaksiUseCase initializes a new instance of TransaksiUseCase.
func NewTransaksiUseCase(repo TransaksiRepository) *TransaksiUseCase {
	return &TransaksiUseCase{
		Repo: repo,
	}
}

// GetAllTransaksi retrieves all transaksis.
func (uc *TransaksiUseCase) GetAllTransaksi() ([]*Transaksi, error) {
	transaksis, err := uc.Repo.GetAllTransaksi()
	if err != nil {
		return nil, err
	}
	return transaksis, nil
}

// CreateTransaksi creates a new transaksi.
func (uc *TransaksiUseCase) CreateTransaksi(transaksi *Transaksi) error {
	return uc.Repo.CreateTransaksi(transaksi)
}

// GetTransaksiByID retrieves a transaksi by its ID.
func (uc *TransaksiUseCase) GetTransaksiByID(id int) (*Transaksi, error) {
	t, err := uc.Repo.GetTransaksiByID(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// UpdateTransaksi updates an existing transaksi.
func (uc *TransaksiUseCase) UpdateTransaksi(transaksi *Transaksi) error {
	if _, err := uc.Repo.GetTransaksiByID(transaksi.ID); err != nil {
		return err
	}
	return uc.Repo.UpdateTransaksi(transaksi)
}

// DeleteTransaksi deletes a transaksi by its ID.
func (uc *TransaksiUseCase) DeleteTransaksi(id int) error {
	if _, err := uc.Repo.GetTransaksiByID(id); err != nil {
		return err
	}
	return uc.Repo.DeleteTransaksi(id)
}

// ErrNotFound indicates that the requested resource was not found.
var ErrNotFound = errors.New("not found")
