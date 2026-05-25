package main
import (
	"fmt"
	"strings"
)

const MAX = 100
type Meja struct {
	NoMeja, Kapasitas, Frekuensi int
	Status string
}

type Pelanggan struct {
	ID int
	Nama, NoHP string
}

type Reservasi struct {
	IDReservasi, IDPelanggan, NoMeja int
	Tanggal, Jam string
}

var dataMeja [MAX]Meja
var dataPelanggan [MAX]Pelanggan
var dataReservasi [MAX]Reservasi
var jumlahMeja int
var jumlahPelanggan int
var jumlahReservasi int

// ===================== MENU =====================

func main() {
	menuUtama()
}

func menuUtama() {
	var pilihan int
	for pilihan != 9 {
		fmt.Println("\n===== APLIKASI RESERVASI RESTORAN =====")
		fmt.Println("1. Tambah Data Meja")
		fmt.Println("2. Tampil Data Meja")
		fmt.Println("3. Edit Data Meja")
		fmt.Println("4. Hapus Data Meja")
		fmt.Println("5. Data Pelanggan")
		fmt.Println("6. Reservasi")
		fmt.Println("7. Cari Meja")
		fmt.Println("8. Statistik")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			tambahMeja()
		} else if pilihan == 2 {
			menuTampilMeja()
		} else if pilihan == 3 {
			editMeja()
		} else if pilihan == 4 {
			hapusMeja()
		} else if pilihan == 5 {
			menuPelanggan()
		} else if pilihan == 6 {
			menuReservasi()
		} else if pilihan == 7 {
			menuCariMeja()
		} else if pilihan == 8 {
			statistik()
		}
	}
}

// ===================== DATA MEJA =====================

func tambahMeja() {
	if jumlahMeja < MAX {
		fmt.Print("Nomor meja: ")
		fmt.Scan(&dataMeja[jumlahMeja].NoMeja)

		fmt.Print("Kapasitas kursi: ")
		fmt.Scan(&dataMeja[jumlahMeja].Kapasitas)

		dataMeja[jumlahMeja].Status = "Kosong"
		dataMeja[jumlahMeja].Frekuensi = 0

		jumlahMeja++
		fmt.Println("Data meja berhasil ditambahkan")
	}
}

func tampilMeja() {
	var i int

	fmt.Println("\n===== DATA MEJA =====")
	for i < jumlahMeja {
		fmt.Println("No Meja     :", dataMeja[i].NoMeja)
		fmt.Println("Kapasitas   :", dataMeja[i].Kapasitas)
		fmt.Println("Status      :", dataMeja[i].Status)
		fmt.Println("Frekuensi   :", dataMeja[i].Frekuensi)
		fmt.Println("---------------------------")
		i++
	}
}

func editMeja() {
	var no int
	var idx int

	fmt.Print("Masukkan nomor meja yang ingin diubah: ")
	fmt.Scan(&no)

	idx = sequentialSearchMeja(no)

	if idx != -1 {
		fmt.Print("Kapasitas baru: ")
		fmt.Scan(&dataMeja[idx].Kapasitas)

		fmt.Print("Status baru: ")
		fmt.Scan(&dataMeja[idx].Status)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Meja tidak ditemukan")
	}
}

func hapusMeja() {
	var no int
	var idx int
	var i int

	fmt.Print("Masukkan nomor meja yang ingin dihapus: ")
	fmt.Scan(&no)

	idx = sequentialSearchMeja(no)

	if idx != -1 {
		i = idx

		for i < jumlahMeja-1 {
			dataMeja[i] = dataMeja[i+1]
			i++
		}

		jumlahMeja--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Meja tidak ditemukan")
	}
}

// ===================== DATA PELANGGAN =====================

func menuPelanggan() {
	var pilih int

	for pilih != 5 {
		fmt.Println("\n===== MENU PELANGGAN =====")
		fmt.Println("1. Tambah Pelanggan")
		fmt.Println("2. Tampil Pelanggan")
		fmt.Println("3. Edit Pelanggan")
		fmt.Println("4. Hapus Pelanggan")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahPelanggan()
		} else if pilih == 2 {
			tampilPelanggan()
		} else if pilih == 3 {
			editPelanggan()
		} else if pilih == 4 {
			hapusPelanggan()
		}
	}
}

func tambahPelanggan() {
	if jumlahPelanggan < MAX {
		fmt.Print("ID Pelanggan: ")
		fmt.Scan(&dataPelanggan[jumlahPelanggan].ID)

		fmt.Print("Nama: ")
		fmt.Scan(&dataPelanggan[jumlahPelanggan].Nama)

		fmt.Print("No HP: ")
		fmt.Scan(&dataPelanggan[jumlahPelanggan].NoHP)

		jumlahPelanggan++
		fmt.Println("Data pelanggan berhasil ditambahkan")
	}
}

func tampilPelanggan() {
	var i int

	fmt.Println("\n===== DATA PELANGGAN =====")

	for i < jumlahPelanggan {
		fmt.Println("ID      :", dataPelanggan[i].ID)
		fmt.Println("Nama    :", dataPelanggan[i].Nama)
		fmt.Println("No HP   :", dataPelanggan[i].NoHP)
		fmt.Println("----------------------")
		i++
	}
}

func editPelanggan() {
	var id int
	var idx int

	fmt.Print("Masukkan ID pelanggan: ")
	fmt.Scan(&id)

	idx = cariPelanggan(id)

	if idx != -1 {
		fmt.Print("Nama baru: ")
		fmt.Scan(&dataPelanggan[idx].Nama)

		fmt.Print("No HP baru: ")
		fmt.Scan(&dataPelanggan[idx].NoHP)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusPelanggan() {
	var id int
	var idx int
	var i int

	fmt.Print("Masukkan ID pelanggan: ")
	fmt.Scan(&id)

	idx = cariPelanggan(id)

	if idx != -1 {
		i = idx

		for i < jumlahPelanggan-1 {
			dataPelanggan[i] = dataPelanggan[i+1]
			i++
		}

		jumlahPelanggan--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func cariPelanggan(id int) int {
	var i int
	var idx int = -1

	for i < jumlahPelanggan {
		if dataPelanggan[i].ID == id {
			idx = i
		}
		i++
	}

	return idx
}

// ===================== RESERVASI =====================

func menuReservasi() {
	var pilih int

	for pilih != 3 {
		fmt.Println("\n===== MENU RESERVASI =====")
		fmt.Println("1. Tambah Reservasi")
		fmt.Println("2. Tampil Reservasi")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahReservasi()
		} else if pilih == 2 {
			tampilReservasi()
		}
	}
}

func tambahReservasi() {
	if jumlahReservasi < MAX {
		fmt.Print("ID Reservasi: ")
		fmt.Scan(&dataReservasi[jumlahReservasi].IDReservasi)

		fmt.Print("ID Pelanggan: ")
		fmt.Scan(&dataReservasi[jumlahReservasi].IDPelanggan)

		fmt.Print("Nomor Meja: ")
		fmt.Scan(&dataReservasi[jumlahReservasi].NoMeja)

		fmt.Print("Tanggal (dd-mm-yyyy): ")
		fmt.Scan(&dataReservasi[jumlahReservasi].Tanggal)

		fmt.Print("Jam: ")
		fmt.Scan(&dataReservasi[jumlahReservasi].Jam)

		updateStatusMeja(dataReservasi[jumlahReservasi].NoMeja)

		jumlahReservasi++

		fmt.Println("Reservasi berhasil ditambahkan")
	}
}

func tampilReservasi() {
	var i int

	fmt.Println("\n===== DATA RESERVASI =====")

	for i < jumlahReservasi {
		fmt.Println("ID Reservasi :", dataReservasi[i].IDReservasi)
		fmt.Println("ID Pelanggan :", dataReservasi[i].IDPelanggan)
		fmt.Println("No Meja      :", dataReservasi[i].NoMeja)
		fmt.Println("Tanggal      :", dataReservasi[i].Tanggal)
		fmt.Println("Jam          :", dataReservasi[i].Jam)
		fmt.Println("---------------------------")
		i++
	}
}

func updateStatusMeja(no int) {
	var idx int

	idx = sequentialSearchMeja(no)

	if idx != -1 {
		dataMeja[idx].Status = "Dipesan"
		dataMeja[idx].Frekuensi++
	}
}

// ===================== SEARCHING =====================

// Sequential Search
func sequentialSearchMeja(no int) int {
	var i int
	var idx int = -1

	for i < jumlahMeja {
		if dataMeja[i].NoMeja == no {
			idx = i
		}
		i++
	}

	return idx
}

// Binary Search
func binarySearchMeja(no int) int {
	var left int = 0
	var right int = jumlahMeja - 1
	var mid int
	var idx int = -1

	sortMejaNoAsc()

	for left <= right {
		mid = (left + right) / 2

		if dataMeja[mid].NoMeja == no {
			idx = mid
			left = right + 1
		} else if dataMeja[mid].NoMeja < no {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return idx
}

func menuCariMeja() {
	var pilih int
	var no int
	var idx int

	fmt.Println("\n===== CARI MEJA =====")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	fmt.Print("Masukkan nomor meja: ")
	fmt.Scan(&no)

	if pilih == 1 {
		idx = sequentialSearchMeja(no)
	} else if pilih == 2 {
		idx = binarySearchMeja(no)
	}

	if idx != -1 {
		fmt.Println("Meja ditemukan")
		fmt.Println("Nomor     :", dataMeja[idx].NoMeja)
		fmt.Println("Kapasitas :", dataMeja[idx].Kapasitas)
		fmt.Println("Status    :", dataMeja[idx].Status)
	} else {
		fmt.Println("Meja tidak ditemukan")
	}
}

// ===================== SORTING =====================

// Selection Sort Ascending
func selectionSortKapasitasAsc() {
	var i, j, min int
	var temp Meja

	i = 0

	for i < jumlahMeja-1 {
		min = i
		j = i + 1

		for j < jumlahMeja {
			if dataMeja[j].Kapasitas < dataMeja[min].Kapasitas {
				min = j
			}
			j++
		}

		temp = dataMeja[i]
		dataMeja[i] = dataMeja[min]
		dataMeja[min] = temp

		i++
	}
}

// Selection Sort Descending
func selectionSortKapasitasDesc() {
	var i, j, max int
	var temp Meja

	i = 0

	for i < jumlahMeja-1 {
		max = i
		j = i + 1

		for j < jumlahMeja {
			if dataMeja[j].Kapasitas > dataMeja[max].Kapasitas {
				max = j
			}
			j++
		}

		temp = dataMeja[i]
		dataMeja[i] = dataMeja[max]
		dataMeja[max] = temp

		i++
	}
}

// Insertion Sort Ascending
func insertionSortKapasitasAsc() {
	var i, j int
	var temp Meja

	i = 1

	for i < jumlahMeja {
		temp = dataMeja[i]
		j = i - 1

		for j >= 0 && dataMeja[j].Kapasitas > temp.Kapasitas {
			dataMeja[j+1] = dataMeja[j]
			j--
		}

		dataMeja[j+1] = temp
		i++
	}
}

// Insertion Sort Descending
func insertionSortKapasitasDesc() {
	var i, j int
	var temp Meja

	i = 1

	for i < jumlahMeja {
		temp = dataMeja[i]
		j = i - 1

		for j >= 0 && dataMeja[j].Kapasitas < temp.Kapasitas {
			dataMeja[j+1] = dataMeja[j]
			j--
		}

		dataMeja[j+1] = temp
		i++
	}
}

func sortMejaNoAsc() {
	var i, j int
	var temp Meja

	i = 1

	for i < jumlahMeja {
		temp = dataMeja[i]
		j = i - 1

		for j >= 0 && dataMeja[j].NoMeja > temp.NoMeja {
			dataMeja[j+1] = dataMeja[j]
			j--
		}

		dataMeja[j+1] = temp
		i++
	}
}

func menuTampilMeja() {
	var pilih int

	fmt.Println("\n===== SORTING DATA MEJA =====")
	fmt.Println("1. Selection Sort Ascending")
	fmt.Println("2. Selection Sort Descending")
	fmt.Println("3. Insertion Sort Ascending")
	fmt.Println("4. Insertion Sort Descending")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortKapasitasAsc()
	} else if pilih == 2 {
		selectionSortKapasitasDesc()
	} else if pilih == 3 {
		insertionSortKapasitasAsc()
	} else if pilih == 4 {
		insertionSortKapasitasDesc()
	}

	tampilMeja()
}

// ===================== STATISTIK =====================

func statistik() {
	var i, j int
	var tanggal [MAX]string
	var jumlah [MAX]int
	var banyakTanggal int
	var ditemukan bool

	fmt.Println("\n===== STATISTIK RESERVASI =====")

	for i < jumlahReservasi {
		ditemukan = false
		j = 0

		for j < banyakTanggal {
			if strings.ToLower(tanggal[j]) == strings.ToLower(dataReservasi[i].Tanggal) {
				jumlah[j]++
				ditemukan = true
			}
			j++
		}

		if ditemukan == false {
			tanggal[banyakTanggal] = dataReservasi[i].Tanggal
			jumlah[banyakTanggal] = 1
			banyakTanggal++
		}

		i++
	}

	i = 0
	for i < banyakTanggal {
		fmt.Println("Tanggal :", tanggal[i])
		fmt.Println("Jumlah Reservasi :", jumlah[i])
		fmt.Println("----------------------")
		i++
	}

	mejaTerlaris()
}

func mejaTerlaris() {
	var i int
	var max int = 0

	i = 1

	for i < jumlahMeja {
		if dataMeja[i].Frekuensi > dataMeja[max].Frekuensi {
			max = i
		}
		i++
	}

	if jumlahMeja > 0 {
		fmt.Println("Meja paling sering dipesan:")
		fmt.Println("Nomor Meja :", dataMeja[max].NoMeja)
		fmt.Println("Jumlah Pesanan :", dataMeja[max].Frekuensi)
	}
}