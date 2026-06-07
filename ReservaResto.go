package main
import "fmt"

const MAX = 100

type Meja struct {
	Nomor, Kapasitas, JumlahPesan int
}

type Pelanggan struct {
	ID int
	Nama, NoHP string
}

type Reservasi struct {
	IDReservasi, IDPelanggan, NomorMeja, Durasi int
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
		fmt.Println("5. Pelanggan")
		fmt.Println("6. Cari Meja")
		fmt.Println("7. Urutkan Meja")
		fmt.Println("8. Reservasi")
		fmt.Println("9. Statistik")
		fmt.Println("10. Keluar")
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
			menuPelanggan()
		} else if pilih == 6 {
			menuCari()
		} else if pilih == 7 {
			menuSort()
		} else if pilih == 8 {
			menuReservasi()
		} else if pilih == 9 {
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
		var nomor int
		var sudahAda bool
		fmt.Print("Nomor meja : ")
		fmt.Scan(&nomor)
		sudahAda = false
		for i := 0; i < n && !sudahAda; i++ {
			if dataMeja[i].Nomor == nomor {
				sudahAda = true
			}
		}
		if sudahAda {
			fmt.Println("Nomor meja sudah terdaftar")
			return
		}
		dataMeja[n].Nomor = nomor
		fmt.Print("Kapasitas : ")
		fmt.Scan(&dataMeja[n].Kapasitas)
		dataMeja[n].JumlahPesan = 0
		fmt.Println("Data meja berhasil ditambahkan")
	}
}

func tampilMeja() {
	n := jumlahMeja()
	fmt.Println("\n===== DATA MEJA =====")
	for i := 0; i < n; i++ {
		fmt.Println("Nomor meja   :", dataMeja[i].Nomor)
		fmt.Println("Kapasitas    :", dataMeja[i].Kapasitas)
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

func menuPelanggan() {
	var pilih int
	for pilih != 5 {
		fmt.Println("\n===== DATA PELANGGAN =====")
		fmt.Println("1. Tambah Pelanggan")
		fmt.Println("2. Tampil Pelanggan")
		fmt.Println("3. Edit Pelanggan")
		fmt.Println("4. Hapus Pelanggan")
		fmt.Println("5. Kembali")
		fmt.Print("Pilih : ")
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
	n := jumlahPelanggan()
	if n < MAX {
		var id int
		var sudahAda bool
		fmt.Print("ID Pelanggan : ")
		fmt.Scan(&id)
		sudahAda = false
		for i := 0; i < n && !sudahAda; i++ {
			if dataPelanggan[i].ID == id {
				sudahAda = true
			}
		}
		if sudahAda {
			fmt.Println("ID pelanggan sudah terdaftar")
			return
		}
		dataPelanggan[n].ID = id
		fmt.Print("Nama : ")
		fmt.Scan(&dataPelanggan[n].Nama)
		fmt.Print("No HP : ")
		fmt.Scan(&dataPelanggan[n].NoHP)
		fmt.Println("Data pelanggan berhasil ditambahkan")
	}
}

func tampilPelanggan() {
	n := jumlahPelanggan()
	fmt.Println("\n===== DATA PELANGGAN =====")
	for i := 0; i < n; i++ {
		fmt.Println("ID    :", dataPelanggan[i].ID)
		fmt.Println("Nama  :", dataPelanggan[i].Nama)
		fmt.Println("No HP :", dataPelanggan[i].NoHP)
		fmt.Println("-------------------------")
	}
}

func editPelanggan() {
	var id int
	var ketemu bool
	n := jumlahPelanggan()
	fmt.Print("Masukkan ID Pelanggan : ")
	fmt.Scan(&id)
	i := 0
	ketemu = false
	for i < n && !ketemu {
		if dataPelanggan[i].ID == id {
			fmt.Print("Nama baru : ")
			fmt.Scan(&dataPelanggan[i].Nama)
			fmt.Print("No HP baru : ")
			fmt.Scan(&dataPelanggan[i].NoHP)
			fmt.Println("Data berhasil diubah")
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusPelanggan() {
	var id int
	var ketemu bool
	n := jumlahPelanggan()
	fmt.Print("Masukkan ID Pelanggan : ")
	fmt.Scan(&id)
	i := 0
	ketemu = false
	for i < n && !ketemu {
		if dataPelanggan[i].ID == id {
			for j := i; j < n-1; j++ {
				dataPelanggan[j] = dataPelanggan[j+1]
			}
			dataPelanggan[n-1] = Pelanggan{}
			fmt.Println("Data berhasil dihapus")
			ketemu = true
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

func selectionSortKapasitas() {
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

func insertionSortKapasitas() {
	var temp Meja
	n := jumlahMeja()
	for i := 1; i < n; i++ {
		temp = dataMeja[i]
		j := i - 1
		for j >= 0 && dataMeja[j].Kapasitas > temp.Kapasitas {
			dataMeja[j+1] = dataMeja[j]
			j--
		}
		dataMeja[j+1] = temp
	}
}

func menuSort() {
	var pilih int
	fmt.Println("\n===== MENU SORT =====")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		selectionSortKapasitas()
		tampilMeja()
	} else if pilih == 2 {
		insertionSortKapasitas()
		tampilMeja()
	}
}

func menuReservasi() {
	var pilih int
	for pilih != 4 {
		fmt.Println("\n===== MENU RESERVASI =====")
		fmt.Println("1. Tambah Reservasi")
		fmt.Println("2. Tampil Reservasi")
		fmt.Println("3. Cek Ketersediaan Meja")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahReservasi()
		} else if pilih == 2 {
			tampilReservasi()
		} else if pilih == 3 {
			cekKetersediaanMeja()
		}
	}
}

func tambahReservasi() {
	var ditemukan, bentrok, pelangganAda bool
	var idRes, idPel, nomorMeja, durasi int
	var tanggal, jam string
	fmt.Print("ID Reservasi : ")
	fmt.Scan(&idRes)
	for i := 0; i < jumlahReservasi(); i++ {
		if dataReservasi[i].IDReservasi == idRes {
			fmt.Println("ID reservasi sudah terdaftar")
			return
		}
	}
	fmt.Print("ID Pelanggan : ")
	fmt.Scan(&idPel)
	pelangganAda = false
	for i := 0; i < jumlahPelanggan() && !pelangganAda; i++ {
		if dataPelanggan[i].ID == idPel {
			pelangganAda = true
		}
	}
	if !pelangganAda {
		fmt.Println("ID pelanggan tidak terdaftar")
		return
	}
	fmt.Print("Nomor Meja : ")
	fmt.Scan(&nomorMeja)
	fmt.Print("Tanggal : ")
	fmt.Scan(&tanggal)
	fmt.Print("Jam : ")
	fmt.Scan(&jam)
	fmt.Print("Durasi (menit) : ")
	fmt.Scan(&durasi)
	jamBaru, menitBaru := 0, 0
	fmt.Sscanf(jam, "%d:%d", &jamBaru, &menitBaru)
	waktuBaruMulai := jamBaru*60 + menitBaru
	waktuBaruSelesai := waktuBaruMulai + durasi
	jumlah := jumlahMeja()
	ditemukan = false
	for i := 0; i < jumlah && !ditemukan; i++ {
		if dataMeja[i].Nomor == nomorMeja {
			ditemukan = true
			n := jumlahReservasi()
			bentrok = false
			for j := 0; j < n && !bentrok; j++ {
				if dataReservasi[j].NomorMeja == nomorMeja && dataReservasi[j].Tanggal == tanggal {
					jamLama, menitLama := 0, 0
					fmt.Sscanf(dataReservasi[j].Jam, "%d:%d", &jamLama, &menitLama)
					waktuLamaMulai := jamLama*60 + menitLama
					waktuLamaSelesai := waktuLamaMulai + dataReservasi[j].Durasi + 1
					if waktuBaruMulai < waktuLamaSelesai && waktuBaruSelesai > waktuLamaMulai {
						bentrok = true
					}
				}
			}
			if bentrok {
				fmt.Println("Meja sudah dipesan pada rentang waktu tersebut")
			} else {
				if n < MAX {
					dataReservasi[n].IDReservasi = idRes
					dataReservasi[n].IDPelanggan = idPel
					dataReservasi[n].NomorMeja = nomorMeja
					dataReservasi[n].Tanggal = tanggal
					dataReservasi[n].Jam = jam
					dataReservasi[n].Durasi = durasi
					dataMeja[i].JumlahPesan++
					fmt.Println("Reservasi berhasil ditambahkan")
				}
			}
		}
	}
	if !ditemukan {
		fmt.Println("Nomor meja tidak ditemukan")
	}
}

func tampilReservasi() {
	var nama, noHP string
	n := jumlahReservasi()
	fmt.Println("\n===== DATA RESERVASI =====")
	for i := 0; i < n; i++ {
		for j := 0; j < jumlahPelanggan(); j++ {
			if dataPelanggan[j].ID == dataReservasi[i].IDPelanggan {
				nama = dataPelanggan[j].Nama
				noHP = dataPelanggan[j].NoHP
			}
		}
		fmt.Println("ID Reservasi   :", dataReservasi[i].IDReservasi)
		fmt.Println("ID Pelanggan   :", dataReservasi[i].IDPelanggan)
		fmt.Println("Nama Pelanggan :", nama)
		fmt.Println("No HP          :", noHP)
		fmt.Println("Nomor Meja     :", dataReservasi[i].NomorMeja)
		fmt.Println("Tanggal        :", dataReservasi[i].Tanggal)
		fmt.Println("Jam            :", dataReservasi[i].Jam)
		fmt.Println("Durasi         :", dataReservasi[i].Durasi, "menit")
		fmt.Println("-----------------------------")
	}
}

func cekKetersediaanMeja() {
	var nomorMeja int
	var tanggal, jam string
	var bentrok bool
	fmt.Print("Nomor Meja : ")
	fmt.Scan(&nomorMeja)
	fmt.Print("Tanggal : ")
	fmt.Scan(&tanggal)
	fmt.Print("Jam : ")
	fmt.Scan(&jam)
	jamCek, menitCek := 0, 0
	fmt.Sscanf(jam, "%d:%d", &jamCek, &menitCek)
	waktuCek := jamCek*60 + menitCek
	bentrok = false
	n := jumlahReservasi()
	for i := 0; i < n && !bentrok; i++ {
		if dataReservasi[i].NomorMeja == nomorMeja && dataReservasi[i].Tanggal == tanggal {
			jamLama, menitLama := 0, 0
			fmt.Sscanf(dataReservasi[i].Jam, "%d:%d", &jamLama, &menitLama)
			waktuMulai := jamLama*60 + menitLama
			waktuSelesai := waktuMulai + dataReservasi[i].Durasi + 1
			if waktuCek >= waktuMulai && waktuCek < waktuSelesai {
				bentrok = true
			}
		}
	}
	if bentrok {
		fmt.Println("Meja tidak tersedia")
	} else {
		fmt.Println("Meja tersedia")
	}
}

func statistik() {
	n := jumlahReservasi()
	fmt.Println("\n===== STATISTIK =====")
	if n == 0 {
		fmt.Println("Belum ada data reservasi")
		return
	}
	for i := 0; i < n; i++ {
		sudahDitampilkan := false
		for j := 0; j < i; j++ {
			if dataReservasi[j].Tanggal == dataReservasi[i].Tanggal {
				sudahDitampilkan = true
			}
		}
		if !sudahDitampilkan {
			tanggal := dataReservasi[i].Tanggal
			totalReservasi := 0
			jumlahTerbanyak := 0
			for j := 0; j < n; j++ {
				if dataReservasi[j].Tanggal == tanggal {
					totalReservasi++
					hitung := 0
					for k := 0; k < n; k++ {
						if dataReservasi[k].Tanggal == tanggal &&
							dataReservasi[k].NomorMeja == dataReservasi[j].NomorMeja {
							hitung++
						}
					}
					if hitung > jumlahTerbanyak {
						jumlahTerbanyak = hitung
					}
				}
			}
			fmt.Println("Tanggal :", tanggal)
			fmt.Println("Total reservasi :", totalReservasi)
			fmt.Print("No meja paling sering dipesan : ")
			jumlahSeri := 0
			for j := 0; j < n; j++ {
				if dataReservasi[j].Tanggal == tanggal {
					sudahDicetak := false
					for k := 0; k < j; k++ {
						if dataReservasi[k].Tanggal == tanggal && dataReservasi[k].NomorMeja == dataReservasi[j].NomorMeja {
							sudahDicetak = true
						}
					}
					if !sudahDicetak {
						hitung := 0
						for k := 0; k < n; k++ {
							if dataReservasi[k].Tanggal == tanggal && dataReservasi[k].NomorMeja == dataReservasi[j].NomorMeja {
								hitung++
							}
						}
						if hitung == jumlahTerbanyak {
							if jumlahSeri > 0 {
								fmt.Print(" & ")
							}
							fmt.Print(dataReservasi[j].NomorMeja)
							jumlahSeri++
						}
					}
				}
			}
			fmt.Println()
			fmt.Println("Jumlah pesanan meja tersebut :", jumlahTerbanyak)
			fmt.Println("---------------------------")
		}
	}
}