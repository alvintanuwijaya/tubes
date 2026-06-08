package main

import "fmt"

const NMAX int = 10000

var nData int = 1

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
	nomorTelepon int
}

type data struct {
	kendaraan     dataKendaraan
	pemilik       dataPemilik
	riwayatServis dataRiwayatServis
}

type tabData [NMAX]data

func tampilanAdmin() {
	fmt.Println("=== MENU ADMIN ===")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Ubah Data Kendaraan")
	fmt.Println("4. Ubah Data Pemilik")
	fmt.Println("5. Tambah Riwayat Servis")
	fmt.Println("6. Tampilkan Data")
	fmt.Println("0. Kembali")
}

func tampilanManager() {
	fmt.Println("=== MENU MANAGER ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("3. Selection Sort Tahun")
	fmt.Println("4. Insertion Sort Tanggal")
	fmt.Println("5. Statistik")
	fmt.Println("0. Kembali")
}

func tambahData(k *tabData, nData *int) {
	var tambah string

	tambah = "y"

	for *nData < NMAX && tambah == "y" {
		fmt.Scan(
			&(*k)[*nData].pemilik.nama,
			&(*k)[*nData].pemilik.alamat,
			&(*k)[*nData].pemilik.nomorTelepon,
		)

		fmt.Scan(
			&(*k)[*nData].kendaraan.platNomor,
			&(*k)[*nData].kendaraan.tahunProduksi,
		)

		*nData = *nData + 1
		fmt.Scan(&tambah)
	}
}

func hapusDataKendaraan(k *tabData, nData *int) {
	var index int

	fmt.Scan(&index)

	if index <= 0 || index >= *nData {
		fmt.Println("Data tidak ditemukan!")
	} else {
		for index <= *nData-2 {
			(*k)[index] = (*k)[index+1]
			index = index + 1
		}

		*nData = *nData - 1
	}
}

func ubahDataKendaraan(k *tabData, nData int) {
	var index int
	var platNomorBaru string
	var tahunProduksiBaru int

	fmt.Scan(&index)

	if index <= 0 || index >= nData {
		fmt.Println("Index tidak ditemukan!")
	} else {
		fmt.Println(
			"Data lama :",
			(*k)[index].kendaraan.platNomor,
			(*k)[index].kendaraan.tahunProduksi,
		)

		fmt.Scan(&platNomorBaru, &tahunProduksiBaru)

		(*k)[index].kendaraan.platNomor = platNomorBaru
		(*k)[index].kendaraan.tahunProduksi = tahunProduksiBaru
	}
}

func ubahDataPemilik(k *tabData, nData int) {
	var index, nomorTeleponBaru int
	var namaBaru, alamatBaru string

	fmt.Scan(&index)

	if index <= 0 || index >= nData {
		fmt.Println("Index tidak ditemukan!")
	} else {
		fmt.Println(
			"Data lama :",
			(*k)[index].pemilik.nama,
			(*k)[index].pemilik.alamat,
			(*k)[index].pemilik.nomorTelepon,
		)

		fmt.Scan(&namaBaru, &alamatBaru, &nomorTeleponBaru)

		(*k)[index].pemilik.nama = namaBaru
		(*k)[index].pemilik.alamat = alamatBaru
		(*k)[index].pemilik.nomorTelepon = nomorTeleponBaru
	}
}

func tambahDataRiwayat(k *tabData, nData int) {
	var index int

	fmt.Scan(&index)

	if index <= 0 || index >= nData {
		fmt.Println("Index tidak ditemukan!")
	} else {
		fmt.Scan(
			&(*k)[index].riwayatServis.jenisKerusakan,
			&(*k)[index].riwayatServis.tanggalPerbaikan,
		)
	}
}

func sequence(k tabData, nData int) {
	var index int
	var platNomor string
	var found bool

	fmt.Scan(&platNomor)

	found = false

	for index = 1; index < nData; index++ {
		if platNomor == k[index].kendaraan.platNomor {
			fmt.Println(
				k[index].kendaraan.platNomor,
				k[index].kendaraan.tahunProduksi,
			)

			found = true
		}
	}

	if found == false {
		fmt.Println("Data tidak ditemukan!")
	}
}

func urutanData(k *tabData, nData int) {
	var i, idx, pass int
	var temp data

	pass = 2

	for pass < nData {
		idx = pass - 1
		i = pass

		for i < nData {
			if (*k)[i].kendaraan.platNomor <
				(*k)[idx].kendaraan.platNomor {

				idx = i
			}

			i = i + 1
		}

		temp = (*k)[pass-1]
		(*k)[pass-1] = (*k)[idx]
		(*k)[idx] = temp

		pass = pass + 1
	}
}

func tampilan(k tabData, nData int) {
	var i int

	for i = 1; i < nData; i++ {
		fmt.Println(
			k[i].pemilik.nama,
			k[i].pemilik.alamat,
			k[i].pemilik.nomorTelepon,
		)

		fmt.Println(
			k[i].kendaraan.platNomor,
			k[i].kendaraan.tahunProduksi,
		)

		fmt.Println(
			k[i].riwayatServis.jenisKerusakan,
			k[i].riwayatServis.tanggalPerbaikan,
		)
	}
}

func binary(k tabData, nData int) {
	var left, right, mid int
	var platNomor string
	var found bool

	urutanData(&k, nData)

	fmt.Scan(&platNomor)

	left = 1
	right = nData - 1
	found = false

	for left <= right && found == false {
		mid = (left + right) / 2

		if platNomor == k[mid].kendaraan.platNomor {
			fmt.Println(
				k[mid].kendaraan.platNomor,
				k[mid].kendaraan.tahunProduksi,
			)

			found = true

		} else if platNomor < k[mid].kendaraan.platNomor {
			right = mid - 1

		} else {
			left = mid + 1
		}
	}

	if found == false {
		fmt.Println("Data tidak ditemukan!")
	}
}

func tahunSelection(k *tabData, nData int) {
	var pass, i, idx int
	var temp data

	pass = 1

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
}

func tanggalInsertion(k *tabData, nData int) {
	var pass, i int
	var temp data

	pass = 2

	for pass < nData {
		temp = (*k)[pass]
		i = pass - 1

		for i >= 1 &&
			temp.riwayatServis.tanggalPerbaikan <
				(*k)[i].riwayatServis.tanggalPerbaikan {

			(*k)[i+1] = (*k)[i]
			i = i - 1
		}

		(*k)[i+1] = temp

		pass = pass + 1
	}
}

func tampilkanStatistik(k tabData, nData int) {
	var i int

	fmt.Println("Jumlah data :", nData-1)

	for i = 1; i < nData; i++ {
		fmt.Println(
			k[i].kendaraan.platNomor,
			k[i].riwayatServis.jenisKerusakan,
		)
	}
}

func menuAdmin(k *tabData, nData *int) {
	var pilihan int

	pilihan = -1

	for pilihan != 0 {
		tampilanAdmin()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahData(k, nData)

		} else if pilihan == 2 {
			hapusDataKendaraan(k, nData)

		} else if pilihan == 3 {
			ubahDataKendaraan(k, *nData)

		} else if pilihan == 4 {
			ubahDataPemilik(k, *nData)

		} else if pilihan == 5 {
			tambahDataRiwayat(k, *nData)

		} else if pilihan == 6 {
			tampilan(*k, *nData)
		}
	}
}

func menuManager(k *tabData, nData int) {
	var pilihan int

	pilihan = -1

	for pilihan != 0 {
		tampilanManager()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			sequence(*k, nData)

		} else if pilihan == 2 {
			binary(*k, nData)

		} else if pilihan == 3 {
			tahunSelection(k, nData)
			tampilan(*k, nData)

		} else if pilihan == 4 {
			tanggalInsertion(k, nData)
			tampilan(*k, nData)

		} else if pilihan == 5 {
			tampilkanStatistik(*k, nData)
		}
	}
}

func main() {
	var k tabData
	var pilihan int

	pilihan = -1

	fmt.Println("====================================================")
	fmt.Println("||                 SELAMAT DATANG                 ||")
	fmt.Println("====================================================")
	fmt.Println("")

	for pilihan != 0 {
		fmt.Println("=== MENU UTAMA ===")
		fmt.Println("1. Admin")
		fmt.Println("2. Manager")
		fmt.Println("0. Keluar")

		fmt.Scan(&pilihan)

		if pilihan == 1 {
			menuAdmin(&k, &nData)

		} else if pilihan == 2 {
			menuManager(&k, nData)
		}
	}
}
