package main
import "fmt"

const MAX = 100

type Meja struct {
	Nomor, Kapasitas, JumlahPesan int
	Tersedia bool
}

type Pelanggan struct {
	ID int
	Nama, NoHP string
}

type Reservasi struct {
	IDReservasi, IDPelanggan, NomorMeja int
	Tanggal, Jam string
}

var dataMeja [MAX]Meja
var dataPelanggan [MAX]Pelanggan
var dataReservasi [MAX]Reservasi

func main() {
	var pilih int
	for pilih != 9 {
		fmt.Println("\n===== RESERVARESTO =====")
		fmt.Println("1. Tambah Meja")
		fmt.Println("2. Tampil Meja")
		fmt.Println("3. Edit Meja")
		fmt.Println("4. Hapus Meja")
		fmt.Println("5. Cari Meja")
		fmt.Println("6. Urutkan Meja")
		fmt.Println("7. Reservasi")
		fmt.Println("8. Statistik")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahMeja()
		} else if pilih == 2 {
			tampilMeja()
		} else if pilih == 3 {
			editMeja()
		} else if pilih == 4 {
			hapusMeja()
		} else if pilih == 5 {
			menuCari()
		} else if pilih == 6 {
			menuSort()
		} else if pilih == 7 {
			menuReservasi()
		} else if pilih == 8 {
			statistik()
		}
	}
}

func jumlahMeja() int {
	i := 0
	for i < MAX && dataMeja[i].Nomor != 0 {
		i++
	}
	return i
}

func jumlahReservasi() int {
	i := 0
	for i < MAX && dataReservasi[i].IDReservasi != 0 {
		i++
	}
	return i
}

func jumlahPelanggan() int {
	i := 0
	for i < MAX && dataPelanggan[i].ID != 0 {
		i++
	}
	return i
}

func tambahMeja() {
	n := jumlahMeja()
	if n < MAX {
		fmt.Print("Nomor meja : ")
		fmt.Scan(&dataMeja[n].Nomor)
		fmt.Print("Kapasitas : ")
		fmt.Scan(&dataMeja[n].Kapasitas)
		dataMeja[n].Tersedia = true
		dataMeja[n].JumlahPesan = 0
		fmt.Println("Data meja berhasil ditambahkan")
	}
}

func tampilMeja() {
	n := jumlahMeja()
	fmt.Println("\n===== DATA MEJA =====")
	for i := 0; i < n; i++ {
		fmt.Println("Nomor meja :", dataMeja[i].Nomor)
		fmt.Println("Kapasitas  :", dataMeja[i].Kapasitas)
		if dataMeja[i].Tersedia {
			fmt.Println("Status     : Tersedia")
		} else {
			fmt.Println("Status     : Dipesan")
		}
		fmt.Println("Jumlah Pesan :", dataMeja[i].JumlahPesan)
		fmt.Println("---------------------------")
	}
}

func editMeja() {
	var nomor int
	var ketemu bool
	n := jumlahMeja()
	fmt.Print("Masukkan nomor meja : ")
	fmt.Scan(&nomor)
	i := 0
	ketemu = false
	for i < n && !ketemu {
		if dataMeja[i].Nomor == nomor {
			fmt.Print("Kapasitas baru : ")
			fmt.Scan(&dataMeja[i].Kapasitas)
			ketemu = true
			fmt.Println("Data berhasil diubah")
		}
		i++
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusMeja() {
	var nomor int
	var ketemu bool
	n := jumlahMeja()
	fmt.Print("Masukkan nomor meja : ")
	fmt.Scan(&nomor)
	i := 0
	ketemu = false
	for i < n && !ketemu {
		if dataMeja[i].Nomor == nomor {
			ketemu = true
			for j := i; j < n-1; j++ {
				dataMeja[j] = dataMeja[j+1]
			}
			dataMeja[n-1] = Meja{}
			fmt.Println("Data berhasil dihapus")
		}
		i++
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearchKapasitas(kapasitas int) int {
	n := jumlahMeja()
	i := 0
	for i < n {
		if dataMeja[i].Kapasitas == kapasitas {
			return i
		}
		i++
	}
	return -1
}

func binarySearchNomor(nomor int) int {
	var kiri, kanan, tengah int
	n := jumlahMeja()
	kiri = 0
	kanan = n - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if dataMeja[tengah].Nomor == nomor {
			return tengah
		} else if nomor < dataMeja[tengah].Nomor {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func menuCari() {
	var pilih, key, hasil int
	fmt.Println("\n===== MENU CARI =====")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Print("Masukkan kapasitas : ")
		fmt.Scan(&key)
		hasil = sequentialSearchKapasitas(key)
		if hasil != -1 {
			fmt.Println("Data ditemukan")
			fmt.Println("Nomor meja :", dataMeja[hasil].Nomor)
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	} else if pilih == 2 {
		selectionSortNomorAsc()
		fmt.Print("Masukkan nomor meja : ")
		fmt.Scan(&key)
		hasil = binarySearchNomor(key)
		if hasil != -1 {
			fmt.Println("Data ditemukan")
			fmt.Println("Kapasitas :", dataMeja[hasil].Kapasitas)
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
}

func selectionSortNomorAsc() {
	var min int
	var temp Meja
	n := jumlahMeja()
	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if dataMeja[j].Nomor < dataMeja[min].Nomor {
				min = j
			}
		}
		temp = dataMeja[i]
		dataMeja[i] = dataMeja[min]
		dataMeja[min] = temp
	}
}

func selectionSortKapasitasAsc() {
	var min int
	var temp Meja
	n := jumlahMeja()
	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if dataMeja[j].Kapasitas < dataMeja[min].Kapasitas {
				min = j
			}
		}
		temp = dataMeja[i]
		dataMeja[i] = dataMeja[min]
		dataMeja[min] = temp
	}
}

func insertionSortKapasitasDesc() {
	var temp Meja
	n := jumlahMeja()
	for i := 1; i < n; i++ {
		temp = dataMeja[i]
		j := i - 1
		for j >= 0 && dataMeja[j].Kapasitas < temp.Kapasitas {
			dataMeja[j+1] = dataMeja[j]
			j--
		}
		dataMeja[j+1] = temp
	}
}

func menuSort() {
	var pilih int
	fmt.Println("\n===== MENU SORT =====")
	fmt.Println("1. Selection Sort Ascending")
	fmt.Println("2. Insertion Sort Descending")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSortKapasitasAsc()
		tampilMeja()
	} else if pilih == 2 {
		insertionSortKapasitasDesc()
		tampilMeja()
	}
}

func menuReservasi() {
	var pilih int
	for pilih != 4 {
		fmt.Println("\n===== MENU RESERVASI =====")
		fmt.Println("1. Tambah Reservasi")
		fmt.Println("2. Tampil Reservasi")
		fmt.Println("3. Kosongkan Meja")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahReservasi()
		} else if pilih == 2 {
			tampilReservasi()
		} else if pilih == 3 {
			kosongkanMeja()
		}
	}
}

func tambahReservasi() {
	var ditemukan bool
	n := jumlahReservasi()
	jumlah := jumlahMeja()
	if n < MAX {
		fmt.Print("ID Reservasi : ")
		fmt.Scan(&dataReservasi[n].IDReservasi)
		fmt.Print("ID Pelanggan : ")
		fmt.Scan(&dataReservasi[n].IDPelanggan)
		fmt.Print("Nomor Meja : ")
		fmt.Scan(&dataReservasi[n].NomorMeja)
		fmt.Print("Tanggal : ")
		fmt.Scan(&dataReservasi[n].Tanggal)
		fmt.Print("Jam : ")
		fmt.Scan(&dataReservasi[n].Jam)
		i := 0
		ditemukan = false
		for i < jumlah && !ditemukan {
			if dataMeja[i].Nomor == dataReservasi[n].NomorMeja {
				if dataMeja[i].Tersedia {
					dataMeja[i].Tersedia = false
					dataMeja[i].JumlahPesan++
					fmt.Println("Reservasi berhasil ditambahkan")
				} else {
					fmt.Println("Meja sedang digunakan")
				}
				ditemukan = true
			}
			i++
		}
		if !ditemukan {
			fmt.Println("Nomor meja tidak ditemukan")
		}
	}
}

func tampilReservasi() {
	n := jumlahReservasi()
	fmt.Println("\n===== DATA RESERVASI =====")
	for i := 0; i < n; i++ {
		fmt.Println("ID Reservasi :", dataReservasi[i].IDReservasi)
		fmt.Println("ID Pelanggan :", dataReservasi[i].IDPelanggan)
		fmt.Println("Nomor Meja   :", dataReservasi[i].NomorMeja)
		fmt.Println("Tanggal      :", dataReservasi[i].Tanggal)
		fmt.Println("Jam          :", dataReservasi[i].Jam)
		fmt.Println("---------------------------")
	}
}

func kosongkanMeja() {
	var nomor int
	var ketemu bool
	n := jumlahMeja()
	fmt.Print("Masukkan nomor meja : ")
	fmt.Scan(&nomor)
	i := 0
	ketemu = false
	for i < n && !ketemu {
		if dataMeja[i].Nomor == nomor {
			dataMeja[i].Tersedia = true
			fmt.Println("Meja berhasil dikosongkan")
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("Meja tidak ditemukan")
	}
}

func statistik() {
	var hitung, mejaTerlaris, jumlahTerbanyak int
	n := jumlahReservasi()
	jumlahTerbanyak = 0
	mejaTerlaris = 0
	for i := 0; i < n; i++ {
		hitung = 0
		for j := 0; j < n; j++ {
			if dataReservasi[i].NomorMeja == dataReservasi[j].NomorMeja {
				hitung++
			}
		}
		if hitung > jumlahTerbanyak {
			jumlahTerbanyak = hitung
			mejaTerlaris = dataReservasi[i].NomorMeja
		}
	}
	fmt.Println("\n===== STATISTIK =====")
	fmt.Println("Meja paling sering dipesan :", mejaTerlaris)
	fmt.Println("Jumlah reservasi :", jumlahTerbanyak)
}