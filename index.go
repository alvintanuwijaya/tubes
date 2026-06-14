package main

import "fmt"

const NMAX int = 10000
const NMAXSERVIS int = 999

type dataKendaraan struct {
	platNomor     string
	tahunProduksi int
}

type dataRiwayatServis struct {
	tanggalPerbaikan int
	jenisKerusakan   string
}

type dataPemilik struct {
	nama         string
	alamat       string
	nomorTelepon string
}

type data struct {
	kendaraan     dataKendaraan
	pemilik       dataPemilik
	riwayatServis [NMAXSERVIS]dataRiwayatServis
	nServis       int
}

type tabData [NMAX]data

func tampilanAdmin(k *tabData, nData *int) {
	var pilihan int

	for pilihan != 6 {
		fmt.Println("====================================================")
		fmt.Println("||               === MENU ADMIN ===               ||")
		fmt.Println("====================================================")
		fmt.Println("1. Tambah Data                                     |")
		fmt.Println("2. Hapus Data                                      |")
		fmt.Println("3. Ubah Data Kendaraan                             |")
		fmt.Println("4. Ubah Data Pemilik                               |")
		fmt.Println("5. Tambah Riwayat Servis                           |")
		fmt.Println("6. Keluar                                          |")
		fmt.Println("====================================================")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahData(k, nData)
			tampilan(*k, *nData)

		} else if pilihan == 2 {
			hapusDataKendaraan(k, nData)
			tampilan(*k, *nData)

		} else if pilihan == 3 {
			ubahDataKendaraan(k, *nData)
			tampilan(*k, *nData)

		} else if pilihan == 4 {
			ubahDataPemilik(k, *nData)
			tampilan(*k, *nData)

		} else if pilihan == 5 {
			tambahDataRiwayat(k, *nData)
			tampilan(*k, *nData)

		} else if pilihan == 6 {
			fmt.Println("====================================================")
			fmt.Println("||                 Terima kasih                   ||")
			fmt.Println("====================================================")

		} else {
			fmt.Println("Fitur tidak tersedia, silakan pilih fitur kembali!")
		}
	}
}

func tampilanManager(k *tabData, nData *int) {
	var pilihan int

	for pilihan != 6 {
		fmt.Println("====================================================")
		fmt.Println("||        === MENU MANAGER OPERASIONAL ===        ||")
		fmt.Println("====================================================")
		fmt.Println("1. Cari Data Pemilik Berdasarkan Nama              |")
		fmt.Println("2. Cari Riwayat Servis Berdasarkan Tanggal         |")
		fmt.Println("3. Selection Sort Tahun Produksi                   |")
		fmt.Println("4. Insertion Sort Tanggal Servis                   |")
		fmt.Println("5. Statistik                                       |")
		fmt.Println("6. Keluar                                          |")
		fmt.Println("====================================================")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			sequence(*k, *nData)

		} else if pilihan == 2 {
			tanggalInsertion(k, *nData)
			binary(*k, *nData)

		} else if pilihan == 3 {
			tahunSelection(k, *nData)
			tampilan(*k, *nData)

		} else if pilihan == 4 {
			tanggalInsertion(k, *nData)
			tampilan(*k, *nData)

		} else if pilihan == 5 {
			tampilkanStatistik(*k, *nData)

		} else if pilihan == 6 {
			fmt.Println("====================================================")
			fmt.Println("||                 Terima kasih                   ||")
			fmt.Println("====================================================")

		} else {
			fmt.Println("Fitur tidak tersedia, silakan pilih fitur kembali!")
		}
	}
}

func tambahData(k *tabData, nData *int) {
	var n, i int

	fmt.Print("Masukan jumlah data kendaraan: ")
	fmt.Scan(&n)

	if *nData+n <= NMAX {
		for i = 0; i < n; i++ {
			fmt.Println("Data ke-", i+1)

			fmt.Print("Masukan nama pemilik: ")
			fmt.Scan(&(*k)[*nData].pemilik.nama)

			fmt.Print("Masukan alamat pemilik: ")
			fmt.Scan(&(*k)[*nData].pemilik.alamat)

			fmt.Print("Masukan nomor telepon: ")
			fmt.Scan(&(*k)[*nData].pemilik.nomorTelepon)

			fmt.Print("Masukan plat nomor kendaraan: ")
			fmt.Scan(&(*k)[*nData].kendaraan.platNomor)

			fmt.Print("Masukan tahun produksi kendaraan: ")
			fmt.Scan(&(*k)[*nData].kendaraan.tahunProduksi)

			*nData = *nData + 1
		}
	} else {
		fmt.Println("Data melebihi kapasitas!")
	}
}

func hapusDataKendaraan(k *tabData, nData *int) {
	var index int

	fmt.Print("Masukan index yang ingin dihapus: ")
	fmt.Scan(&index)
	index = index - 1

	if index < 0 || index >= *nData {
		fmt.Println("Data tidak ditemukan!")
	} else {
		for index <= *nData-2 {
			(*k)[index] = (*k)[index+1]
			index = index + 1
		}

		*nData = *nData - 1
		fmt.Println("Data berhasil dihapus.")
	}
}

func ubahDataKendaraan(k *tabData, nData int) {
	var index int
	var platNomorBaru string
	var tahunProduksiBaru int

	fmt.Print("Masukan index data kendaraan yang ingin diubah: ")
	fmt.Scan(&index)
	index = index - 1

	if index >= nData || index < 0 {
		fmt.Println("Index tidak ditemukan!")
	} else {
		fmt.Println(
			"Data lama :",
			(*k)[index].kendaraan.platNomor,
			(*k)[index].kendaraan.tahunProduksi,
		)

		fmt.Print("Masukan plat nomor baru: ")
		fmt.Scan(&platNomorBaru)

		fmt.Print("Masukan tahun produksi baru: ") //ini buat ubah tahun produksi kalau misalny ada  salah input sebelumnya misal w205 harusnya 2015 tapi input sbelumnya 2014 -alvin
		fmt.Scan(&tahunProduksiBaru)

		(*k)[index].kendaraan.platNomor = platNomorBaru
		(*k)[index].kendaraan.tahunProduksi = tahunProduksiBaru

		fmt.Println("Data kendaraan berhasil diubah.")
	}
}

func ubahDataPemilik(k *tabData, nData int) {
	var index int
	var namaBaru, alamatBaru, nomorTeleponBaru string

	fmt.Print("Masukan index data pemilik yang ingin diubah: ")
	fmt.Scan(&index)
	index = index - 1

	if index >= nData || index < 0 {
		fmt.Println("Index tidak ditemukan!")
	} else {
		fmt.Println(
			"Data lama :",
			(*k)[index].pemilik.nama,
			(*k)[index].pemilik.alamat,
			(*k)[index].pemilik.nomorTelepon,
		)

		fmt.Print("Masukan nama baru: ")
		fmt.Scan(&namaBaru)

		fmt.Print("Masukan alamat baru: ")
		fmt.Scan(&alamatBaru)

		fmt.Print("Masukan nomor telepon baru: ")
		fmt.Scan(&nomorTeleponBaru)

		(*k)[index].pemilik.nama = namaBaru
		(*k)[index].pemilik.alamat = alamatBaru
		(*k)[index].pemilik.nomorTelepon = nomorTeleponBaru

		fmt.Println("Data pemilik berhasil diubah.")
	}
}

func tambahDataRiwayat(k *tabData, nData int) {
	var index, n int

	fmt.Print("Masukan index kendaraan: ")
	fmt.Scan(&index)
	index = index - 1

	if index >= nData || index < 0 {
		fmt.Println("Index diluar batas!")
	} else {
		n = (*k)[index].nServis

		if n >= NMAXSERVIS {
			fmt.Println("Riwayat servis sudah penuh!")
		} else {
			fmt.Print("Masukan jenis kerusakan: ")
			fmt.Scan(&(*k)[index].riwayatServis[n].jenisKerusakan)

			fmt.Print("Masukan tanggal perbaikan, contoh 20260526: ")
			fmt.Scan(&(*k)[index].riwayatServis[n].tanggalPerbaikan)

			(*k)[index].nServis = (*k)[index].nServis + 1

			fmt.Println("Riwayat servis berhasil ditambahkan.")
		}
	}
}

func sequence(k tabData, nData int) {
	var index int
	var nama string
	var found bool

	fmt.Print("Masukan nama pemilik yang dicari: ")
	fmt.Scan(&nama)

	found = false

	fmt.Println("=== DATA PEMILIK ===")

	for index = 0; index < nData; index++ {
		if nama == k[index].pemilik.nama {
			fmt.Println("Nama Pemilik:", k[index].pemilik.nama)
			fmt.Println("Alamat:", k[index].pemilik.alamat)
			fmt.Println("Nomor Telepon:", k[index].pemilik.nomorTelepon)

			found = true
		}
	}

	if found == false {
		fmt.Println("Data pemilik tidak ditemukan!")
	}
}

func tampilan(k tabData, nData int) {
	var i, j int

	fmt.Println(" ")
	fmt.Println("=== DATA KENDARAAN ===")

	if nData == 0 {
		fmt.Println("Belum ada data.")
	} else {
		for i = 0; i < nData; i++ {
			fmt.Println(" ")
			fmt.Println("Index:", i+1)
			fmt.Println("Nama:", k[i].pemilik.nama)
			fmt.Println("Alamat:", k[i].pemilik.alamat)
			fmt.Println("Nomor Telepon:", k[i].pemilik.nomorTelepon)
			fmt.Println("Plat Nomor:", k[i].kendaraan.platNomor)
			fmt.Println("Tahun Produksi:", k[i].kendaraan.tahunProduksi)

			if k[i].nServis == 0 {
				fmt.Println("Riwayat Servis: Belum ada")
			} else {
				fmt.Println("Riwayat Servis:")
				for j = 0; j < k[i].nServis; j++ {
					fmt.Println(
						j+1,
						k[i].riwayatServis[j].jenisKerusakan,
						k[i].riwayatServis[j].tanggalPerbaikan,
					)
				}
			}
		}
	}
}

func tahunSelection(k *tabData, nData int) {
	var pass, i, idx int
	var temp data

	pass = 0

	for pass < nData-1 {
		idx = pass
		i = pass + 1

		for i < nData {
			if (*k)[i].kendaraan.tahunProduksi <
				(*k)[idx].kendaraan.tahunProduksi {

				idx = i
			}

			i = i + 1
		}

		temp = (*k)[pass]
		(*k)[pass] = (*k)[idx]
		(*k)[idx] = temp

		pass = pass + 1
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tahun produksi.")
}

func tanggalTerakhir(d data) int {
	var hasil int

	hasil = 0

	if d.nServis > 0 {
		hasil = d.riwayatServis[d.nServis-1].tanggalPerbaikan
	}

	return hasil
}

func tanggalInsertion(k *tabData, nData int) {
	var pass, i int
	var temp data

	pass = 1

	for pass < nData {
		temp = (*k)[pass]
		i = pass - 1

		for i >= 0 && tanggalTerakhir(temp) > tanggalTerakhir((*k)[i]) {
			(*k)[i+1] = (*k)[i]
			i = i - 1
		}

		(*k)[i+1] = temp

		pass = pass + 1
	}

	fmt.Println("Data berhasil diurutkan berdasarkan tanggal servis terakhir.")
}

func binary(k tabData, nData int) {
	var left, right, mid int
	var tanggal int
	var found bool
	var i int

	fmt.Print("Masukan tanggal servis yang dicari, contoh 20260526: ")
	fmt.Scan(&tanggal)

	left = 0
	right = nData - 1
	found = false

	fmt.Println("=== RIWAYAT SERVIS KENDARAAN ===")

	for left <= right && found == false {
		mid = (left + right) / 2

		if tanggal == tanggalTerakhir(k[mid]) {
			fmt.Println("Tanggal Servis:", tanggalTerakhir(k[mid]))
			fmt.Println("Plat Nomor:", k[mid].kendaraan.platNomor)
			fmt.Println("Tahun Produksi:", k[mid].kendaraan.tahunProduksi)

			for i = 0; i < k[mid].nServis; i++ {
				if k[mid].riwayatServis[i].tanggalPerbaikan == tanggal {
					fmt.Println("Jenis Kerusakan:", k[mid].riwayatServis[i].jenisKerusakan)
				}
			}

			found = true

		} else if tanggal < tanggalTerakhir(k[mid]) {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}

	if found == false {
		fmt.Println("Riwayat servis tidak ditemukan!")
	}
}

func tampilkanStatistik(k tabData, nData int) {
	var i, j, s, maxCount, bulan int
	var modusKerusakan string
	var arrKerusakan [NMAX]string
	var arrCount [NMAX]int
	var countBulan [13]int
	var n int
	var found bool

	for i = 0; i < nData; i++ {
		for s = 0; s < k[i].nServis; s++ {
			bulan = (k[i].riwayatServis[s].tanggalPerbaikan / 100) % 100

			if bulan >= 1 && bulan <= 12 {
				countBulan[bulan] = countBulan[bulan] + 1
			}

			found = false
			for j = 0; j < n && found == false; j++ {
				if arrKerusakan[j] == k[i].riwayatServis[s].jenisKerusakan {
					arrCount[j] = arrCount[j] + 1
					found = true
				}
			}

			if found == false {
				arrKerusakan[n] = k[i].riwayatServis[s].jenisKerusakan
				arrCount[n] = 1
				n = n + 1
			}
		}
	}

	fmt.Println("--- Statistik Kendaraan per Bulan ---")
	for i = 1; i <= 12; i++ {
		if countBulan[i] > 0 {
			fmt.Println("Bulan", i, ":", countBulan[i], "kendaraan")
		}
	}

	for i = 0; i < n; i++ {
		if arrCount[i] > maxCount {
			maxCount = arrCount[i]
			modusKerusakan = arrKerusakan[i]
		}
	}

	if n > 0 {
		fmt.Println("Kategori kerusakan paling sering muncul:", modusKerusakan)
	} else {
		fmt.Println("Belum ada data servis.")
	}
}

func main() {
	var pilihan int
	var k tabData
	var nData int

	pilihan = 0

	for pilihan != 3 {
		fmt.Println("====================================================")
		fmt.Println("||                 SELAMAT DATANG                 ||")
		fmt.Println("====================================================")
		fmt.Println("| Silahkan pilih                                   |")
		fmt.Println("| 1. Admin                                         |")
		fmt.Println("| 2. Manajer operasional                           |")
		fmt.Println("| 3. Keluar                                        |")
		fmt.Println("====================================================")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tampilanAdmin(&k, &nData)

		} else if pilihan == 2 {
			tampilanManager(&k, &nData)

		} else if pilihan == 3 {
			fmt.Println("Program selesai.")

		} else {
			fmt.Println("Pilihan menu tidak tersedia.")
		}
	}
}
