package main
import "fmt"

const NMAX int = 10000
var nData int = 1
type dataKendaraan struct {
	platNomor string
	tahunProduksi int
}

type dataRiwayatServis struct {
	tanggalPerbaikan int
	jenisKerusakan string
}

type dataPemilik struct {
	nama string
	alamat string
	nomorTelepon int
}

type data struct {
	kendaraan dataKendaraan
	pemilik dataPemilik
	riwayatServis dataRiwayatServis
}
type tabData [NMAX]data


func tampilanAdmin(pilihan) {

}

func tampilanManager()


func tambahData(k *tabData, nData *int) {
	for *nData < 4 {
		fmt.Scan(&k[*nData].pemilik.nama, &k[*nData].pemilik.alamat, &k[*nData].pemilik.nomorTelepon)
		fmt.Scan(&k[*nData].kendaraan.platNomor, &k[*nData].kendaraan.tahunProduksi)
		*nData++
	}
}

func hapusDataKendaraan(k *tabData, nData *int) {
	var index int
	fmt.Scan(&index)
	if index <= 0 && index >= *nData {
		fmt.Println("Data tidak ditemukan!")
	} else {
		for index <= *nData-2 {
			k[index] = k[index+1]
			index++
		}
		*nData--
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
		fmt.Println("Data lama :", (*k)[index].kendaraan.platNomor, (*k)[index].kendaraan.tahunProduksi)
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
		fmt.Println("Data lama :", (*k)[index].pemilik.nama, (*k)[index].pemilik.alamat, (*k)[index].pemilik.nomorTelepon)
		fmt.Scan(&namaBaru, &alamatBaru, &nomorTeleponBaru)
		(*k)[index].pemilik.nama = namaBaru
		(*k)[index].pemilik.alamat = alamatBaru
		(*k)[index].pemilik.nomorTelepon = nomorTeleponBaru
	}
}

func tambahDataRiwayat(k tabData, nData int) {
	var index int
	fmt.Scan(&k[index].riwayatServis.jenisKerusakan, &k[index].riwayatServis.tanggalPerbaikan)
}

func sequence(k tabData, nData int) {
	var index int
	var platNomor string
	fmt.Scan(&platNomor)
	for index = 1; index <= nData; index++ {
		if platNomor == k[index].kendaraan.platNomor {
			fmt.Println(k[index].kendaraan.platNomor, k[index].kendaraan.tahunProduksi)
		} else {
			fmt.Println("Data tidak ditemukan!")
		}
	}
}

func urutanData(k tabData, nData int) {
	var panjangPlat, i, j int 
	var subString string 
	for i = 1; i <=nData; i++ {
		subString = ""
		panjangPlat = len(k[i].kendaraan.platNomor)
		for k[i].kendaraan.platNomor[j] != '_' {
			subString += string(k[i].kendaraan.platNomor[j])

		}

		
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
			if k[i].kendaraan.platNomor < k[idx].kendaraan.platNomor{
				idx = i
			}
			i++
		}
		temp = k[pass-1]
		k[pass-1] = k[idx]
		k[idx] = temp
		pass++	
	}
}

func tampilan(k tabData, nData int) {
	var i int
	for i = 1; i < nData; i++ {
		fmt.Println(k[i].pemilik.nama, k[i].pemilik.alamat, k[i].pemilik.nomorTelepon)
		fmt.Println(k[i].kendaraan.platNomor, k[i].kendaraan.tahunProduksi)
	}
}

func binary(k tabData, nData int) {
	var karakter string

}

func tahunSelection()

func tanggalInsertion()

func tampilkanStatistik()



func main(){
	
	var k tabData
	fmt.Println("====================================================")
	fmt.Println("||                 SELAMAT DATANG                 ||")
	fmt.Println("====================================================")
	fmt.Println("")

	if pilihan == 1 {
		tampilanAdmin()
	} else if pilihan == 2 {
		tampilanManager()
	}
	*/
	tambahData(&k, &nData)
	tampilan(k, nData)
	urutanData(&k, nData)
	tampilan(k, nData)
}