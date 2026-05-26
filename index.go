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

func tampilanAdmin(pilihan int) {
	fmt.Println("=== MENU ADMIN ===")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Ubah Data Kendaraan")
	fmt.Println("4. Ubah Data Pemilik")
	fmt.Println("5. Tambah Riwayat Servis")
	fmt.Println("6. Tampilkan Data")
}

func tampilanManager() {
	fmt.Println("=== MENU MANAGER ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("3. Selection Sort Tahun")
	fmt.Println("4. Insertion Sort Tanggal")
	fmt.Println("5. Statistik")
}

func tambahData(k *tabData, nData *int) {
	for *nData < 4 {

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
			index++
		}

		*nData = *nData - 1
	}
}

func ubahDataKendaraan(k *tabData, nData int) {
	var index int
	var platNomorBaru string
	var tahunProduksiBaru int

	fmt.Scan(&index)

	if index > nData || index <= 0 {
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

	if index > nData || index <= 0 {
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

	fmt.Scan(
		&(*k)[index].riwayatServis.jenisKerusakan,
		&(*k)[index].riwayatServis.tanggalPerbaikan,
	)
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

	if !found {
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

			i++
		}

		temp = (*k)[pass-1]
		(*k)[pass-1] = (*k)[idx]
		(*k)[idx] = temp

		pass++
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
	}
}

func binary(k tabData, nData int) {
	var left, right, mid int
	var platNomor string
	var found bool

	fmt.Scan(&platNomor)

	left = 1
	right = nData - 1
	found = false

	for left <= right {

		mid = (left + right) / 2

		if platNomor == k[mid].kendaraan.platNomor {

			fmt.Println(
				k[mid].kendaraan.platNomor,
				k[mid].kendaraan.tahunProduksi,
			)

			found = true

			left = right + 1

		} else if platNomor <
			k[mid].kendaraan.platNomor {

			right = mid - 1

		} else {

			left = mid + 1
		}
	}

	if !found {
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

			i++
		}

		temp = (*k)[pass]
		(*k)[pass] = (*k)[idx]
		(*k)[idx] = temp

		pass++
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
			i--
		}

		(*k)[i+1] = temp

		pass++
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

func main() {

	var k tabData

	fmt.Println("====================================================")
	fmt.Println("||                 SELAMAT DATANG                 ||")
	fmt.Println("====================================================")
	fmt.Println("")

	tambahData(&k, &nData)

	tampilan(k, nData)

	urutanData(&k, nData)

	tampilan(k, nData)
}
