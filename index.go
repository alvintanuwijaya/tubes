package main
import("fmt")

const NMAX int = 10000


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
	riwayatServis dataRiwayatServis
}

type tabData [NMAX]data

func tampilanAdmin(k *tabData, nData *int) {
	var pilihan int

	for pilihan != 6{
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
		switch pilihan {
			case 1:
				tambahData(k, nData)
				tampilan(*k, *nData)
			case 2:
				hapusDataKendaraan(k, nData)
				tampilan(*k, *nData)
			case 3:
				ubahDataKendaraan(k, *nData)
				tampilan(*k, *nData)
			case 4:
				ubahDataPemilik(k, *nData)
				tampilan(*k, *nData)
			case 5:
				tambahDataRiwayat(k, *nData)
				tampilan(*k, *nData)
			case 6:
				fmt.Println("====================================================")
				fmt.Println("||                 Terima kasih                   ||")
				fmt.Println("||         sudah menggunakan aplikasi ini         ||")
				fmt.Println("====================================================")
		}
		if pilihan <= 0 || pilihan > 6 {
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
		fmt.Println("1. Cari Plat Nomor                                 |")
		fmt.Println("2. Cari Data Kendaraan                             |")
		fmt.Println("3. Selection Sort Tahun                            |")
		fmt.Println("4. Insertion Sort Tanggal                          |")
		fmt.Println("5. Statistik                                       |")
		fmt.Println("6. Keluar                                          |")
		fmt.Println("====================================================")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)
		switch pilihan {
			case 1:
				sequence(*k, *nData)
			case 2:
				urutanData(k, *nData)
				binary(*k, *nData)
			case 3:
				tahunSelection(k, *nData)
				tampilan(*k, *nData)
			case 4:
				tanggalInsertion(k, *nData)
				tampilan(*k, *nData)
			case 5:
				tampilkanStatistik(*k, *nData)
			case 6:
				fmt.Println("====================================================")
				fmt.Println("||                 Terima kasih                   ||")
				fmt.Println("||         sudah menggunakan aplikasi ini         ||")
				fmt.Println("====================================================")
		}
		if pilihan <= 0 || pilihan > 6 {
			fmt.Println("Fitur tidak tersedia, silakan pilih fitur kembali!")
		}
	}
}

func tambahData(k *tabData, nData *int) {
	var n, i int
	fmt.Scan(&n)
	if *nData+n <= NMAX {
        for i = 0; i < n; i++ {
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
}

func hapusDataKendaraan(k *tabData, nData *int) {
	var index int

	fmt.Scan(&index)
	index = index - 1

	if index < 0 || index >= *nData {
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
	index = index - 1

	if index >= nData || index < 0 {
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
	var index int
	var namaBaru, alamatBaru, nomorTeleponBaru string

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

		fmt.Scan(&namaBaru, &alamatBaru, &nomorTeleponBaru)

		(*k)[index].pemilik.nama = namaBaru
		(*k)[index].pemilik.alamat = alamatBaru
		(*k)[index].pemilik.nomorTelepon = nomorTeleponBaru
	}
}

func tambahDataRiwayat(k *tabData, nData int) {
	var index int

	fmt.Scan(&index)
	index = index - 1

	if index >= nData || index < 0 {
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

	for index = 0; index < nData && !found; index++ {

		if platNomor == k[index].kendaraan.platNomor {
			fmt.Print(index+1, " ")
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

	pass = 1

	for pass < nData {

		idx = pass - 1
		i = pass

		for i < nData {

			if (*k)[i].kendaraan.platNomor < (*k)[idx].kendaraan.platNomor {
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

	fmt.Println(" ")
    for i = 0; i < nData; i++ {
		fmt.Println(" ")
        fmt.Print(i+1, " ")
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

	left = 0
	right = nData - 1
	found = false

	for left <= right && !found {

		mid = (left + right) / 2

		if platNomor == k[mid].kendaraan.platNomor {
			fmt.Print(mid+1, " ")
			fmt.Println(
				k[mid].kendaraan.platNomor,
				k[mid].kendaraan.tahunProduksi,
			)
			found = true
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

	pass = 0

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

	pass = 1

	for pass < nData {

		temp = (*k)[pass]
		i = pass - 1

		for i >= 0 && temp.riwayatServis.tanggalPerbaikan < (*k)[i].riwayatServis.tanggalPerbaikan {

			(*k)[i+1] = (*k)[i]
			i--
		}

		(*k)[i+1] = temp

		pass++
	}
}

func tampilkanStatistik(k tabData, nData int) {
	var i, j, maxCount, bulan int
    var modusKerusakan string
    var arrKerusakan [NMAX]string
    var arrCount [NMAX]int
	var countBulan [13]int
	var n int

	for i = 0; i < nData; i++ {
        bulan = (k[i].riwayatServis.tanggalPerbaikan / 100) % 100
        if bulan >= 1 && bulan <= 12 {
            countBulan[bulan]++
        }

        var found bool = false
        for j = 0; j < n && !found; j++ {
            if arrKerusakan[j] == k[i].riwayatServis.jenisKerusakan {
                arrCount[j]++
                found = true
            }
        }
        if !found {
            arrKerusakan[n] = k[i].riwayatServis.jenisKerusakan
            arrCount[n] = 1
            n++
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

    if nData > 0 {
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
            fmt.Println("")
        } else {
            fmt.Println("ilihan menu tidak tersedia.")
        }
    }
}
