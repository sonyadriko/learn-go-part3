package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"tokovan/internal/barang"
	"tokovan/internal/transaksi"
)

func main() {
	// Initialize MySQL database connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	barangRepo := barang.NewMySQLBarangRepo(db)
	transaksiRepo := transaksi.NewMySQLTransaksiRepo(db)

	// Initialize use cases
	barangUC := barang.NewBarangUseCase(barangRepo)
	transaksiUC := transaksi.NewTransaksiUseCase(transaksiRepo)

	// Initialize router
	r := mux.NewRouter()

	// Define API endpoints for Barang
	r.HandleFunc("/barang", createBarangHandler(barangUC)).Methods("POST")
	r.HandleFunc("/barang", getAllBarangHandler(barangUC)).Methods("GET")
	r.HandleFunc("/barang/{id}", getBarangByIDHandler(barangUC)).Methods("GET")
	r.HandleFunc("/barang/{id}", updateBarangHandler(barangUC)).Methods("PUT")
	r.HandleFunc("/barang/{id}", deleteBarangHandler(barangUC)).Methods("DELETE")

	// Define API endpoints for Transaksi
	r.HandleFunc("/transaksi", createTransaksiHandler(transaksiUC)).Methods("POST")
	r.HandleFunc("/transaksi", getAllTransaksiHandler(transaksiUC)).Methods("GET")
	r.HandleFunc("/transaksi/{id}", getTransaksiByIDHandler(transaksiUC)).Methods("GET")
	r.HandleFunc("/transaksi/{id}", updateTransaksiHandler(transaksiUC)).Methods("PUT")
	r.HandleFunc("/transaksi/{id}", deleteTransaksiHandler(transaksiUC)).Methods("DELETE")

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createBarangHandler(usecase *barang.BarangUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b barang.Barang
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.CreateBarang(&b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)
	}
}

func getAllBarangHandler(usecase *barang.BarangUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		barangs, err := usecase.GetAllBarang()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(barangs)
	}
}

func getBarangByIDHandler(usecase *barang.BarangUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b, err := usecase.GetBarangByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(b)
	}
}

func updateBarangHandler(usecase *barang.BarangUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var b barang.Barang
		err = json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		b.ID = id

		if err := usecase.UpdateBarang(&b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(b)
	}
}

func deleteBarangHandler(usecase *barang.BarangUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.DeleteBarang(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func createTransaksiHandler(usecase *transaksi.TransaksiUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t transaksi.Transaksi
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.CreateTransaksi(&t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	}
}

func getAllTransaksiHandler(usecase *transaksi.TransaksiUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transaksis, err := usecase.GetAllTransaksi()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaksis)
	}
}

func getTransaksiByIDHandler(usecase *transaksi.TransaksiUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := usecase.GetTransaksiByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(t)
	}
}

func updateTransaksiHandler(usecase *transaksi.TransaksiUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var t transaksi.Transaksi
		err = json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.ID = id

		if err := usecase.UpdateTransaksi(&t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(t)
	}
}

func deleteTransaksiHandler(usecase *transaksi.TransaksiUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := usecase.DeleteTransaksi(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
