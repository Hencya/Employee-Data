package main

import (
	"fmt"
	"os"
)

const N = 1000000

type dip struct {
	kode     string
	nama     string
	golongan string
	umur     int
	alamat   string
	reward   string
	absen    historyAbsen
}
type historyAbsen struct {
	jamMasuk  int
	jamKeluar int
	lembur    int //pulang lembur
	cepat     int //pulang cepat
}
type ArrType struct { //tipe array data pegawainya
	T [N]dip
	n int //banyak array
}

var (
	data ArrType
)

func main() {
	if fileExists("dataPegawai") {
		loadData()
	}
	menu()
	saveData()
}

func masukDataPegawai() { // Memasukan data pegawai
	var (
		t dip //temporary
		m int
	)
	m = data.n
	fmt.Println("Masukkan data pegawai dan diakhiri dengan 0 0 0 0 0")
	fmt.Print("Masukan Kode Pegawai\t: ")
	fmt.Scan(&t.kode)
	fmt.Print("Masukan Nama Pegawai\t: ")
	fmt.Scan(&t.nama)
	fmt.Print("Masukan Golongan Pegawai: ")
	fmt.Scan(&t.golongan)
	fmt.Print("Masukan Umur Pegawai\t: ")
	fmt.Scan(&t.umur)
	fmt.Print("Masukan Alamat Pegawai\t: ")
	fmt.Scan(&t.alamat)
	for m < N && (t.kode != "0" || t.nama != "0" || t.golongan != "0" || t.umur != 0 || t.alamat != "0") {
		data.T[m] = t
		m++
		fmt.Println()
		fmt.Print("Masukan Kode Pegawai\t: ")
		fmt.Scan(&t.kode)
		fmt.Print("Masukan Nama Pegawai\t: ")
		fmt.Scan(&t.nama)
		fmt.Print("Masukan Golongan Pegawai\t: ")
		fmt.Scan(&t.golongan)
		fmt.Print("Masukan Umur Pegawai\t: ")
		fmt.Scan(&t.umur)
		fmt.Print("Masukan Alamat Pegawai\t: ")
		fmt.Scan(&t.alamat)
	}
	data.n = m
	sortAscendKode()
}

func binarySearch(key string) (bool, int) {
	var indFound int
	kr := 0
	kn := data.n - 1
	found := false
	for kr <= kn && !found {
		med := (kr + kn) / 2
		if data.T[med].kode > key {
			kn = med - 1
		} else if data.T[med].kode < key {
			kr = med + 1
		} else if data.T[med].kode == key {
			found = true
			indFound = med
		} else {
			fmt.Println("Error 404: Not Found")
		}
	}
	return found, indFound
}

func cariKodePegawai(key string) { //key adalah kode yang dicari
	found, indFound := binarySearch(key)
	if found == true {
		fmt.Println("Nama pegawai\t:", data.T[indFound].nama)
		fmt.Println("Golongan pegawai\t:", data.T[indFound].golongan)
		fmt.Println("Umur pegawai\t:", data.T[indFound].umur)
		fmt.Println("Alamat pegawai\t:", data.T[indFound].alamat)
		fmt.Println("Reward pegawai\t:", data.T[indFound].reward)
	} else {
		fmt.Println("Error 404: Not Found")
	}
}

func menu() {
	var input int
	var k string
	fmt.Print(`
===============================================================================================
				Aplikasi Kelola Absensi Pegawai 
===============================================================================================
Program menu :
1. Input data pegawai 
2. Cari data pegawai 
3. Tambahkan histori pegawai
4. Tampilkan histori pegawai
5. Proses reward
6. Menampilkann Reward
7. Menampilkan Data Pegawai(ascending)
8. Exit
===============================================================================================	
`)

	for input != 8 {
		fmt.Print("Masukan nomer: ")
		fmt.Scan(&input)
		if input == 1 {
			masukDataPegawai()
			menu()
		} else if input == 2 {
			fmt.Print("Kode: ")
			fmt.Scan(&k)
			cariKodePegawai(k)
			menu()
		} else if input == 3 {
			historiAbsensi()
			menu()
		} else if input == 4 {
			tampilHistori()
			menu()
		} else if input == 5 {
			reward()
		} else if input == 6 {
			tampilReward()
		} else if input == 7 {
			sortAscendNama()
		} else if input == 8 {

		} else {
			fmt.Printf("Input tidak sesuai. Maukah anda mengulang kembali ke menu?\n1. Ya\n2. Kadit\n")
			fmt.Scan(&input)
			if input == 1 {
				menu()
			} else {
				saveData()
			}
		}
	}
}

func loadData() {
	var t dip
	f, err := os.OpenFile("dataPegawai", os.O_CREATE|os.O_RDONLY, 0777)

	if err != nil {
		fmt.Println(err)
	}
	for a := 0; a < N; a++ {
		fmt.Fscanf(f, "%s {%s} %s %d {%s} {%s}", t.kode, t.nama, t.golongan, t.umur, t.alamat, t.reward)
		if t.kode != "0" || t.nama != "0" || t.golongan != "0" || t.umur != 0 || t.alamat != "0" {
			data.T[a] = t
			data.n++
		} else {
			break
		}
	}
	f.Close()
}
func saveData() {
	f, err := os.OpenFile("dataPegawai", os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		fmt.Println(err)
	}

	for a := 0; a < data.n; a++ {
		fmt.Fprintf(f, "%s {%s} %s %d {%s} {%s}\n", data.T[a].kode, data.T[a].nama, data.T[a].golongan, data.T[a].umur, data.T[a].alamat, data.T[a].reward)
	}
	fmt.Fprintln(f, "0 {0} 0 0 {0} {0}")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func historiAbsensi() {
	var (
		i   int
		key string
	)
	fmt.Print("Masukkan kode pegawai: ")
	fmt.Scan(&key)
	found, indFound := binarySearch(key)
	if found == false {
		fmt.Println("Kode pegawai yang ingin dimasukan histori absensinya tidak ada")
	} else {
		fmt.Print("masukan jam masuk dari pegawai")
		fmt.Scan(&data.T[indFound].absen.jamMasuk)
		fmt.Print("masukan jam keluar dari pegawai")
		fmt.Scan(&data.T[indFound].absen.jamKeluar)

		if data.T[indFound].absen.jamMasuk > 8 {
			data.T[indFound].absen.lembur = data.T[indFound].absen.jamMasuk - 8
		} else if data.T[indFound].absen.jamMasuk < 8 {
			data.T[indFound].absen.cepat = 8 - data.T[indFound].absen.jamMasuk
		}
	}

	fmt.Println(`
		Apakah anda ingin memasukkan data lagi?
		1. Ya
		2. Tidak
		`)
	fmt.Scan(&i)
	if i == 1 {
		historiAbsensi()
	} else {
		menu()
	}

}

func tampilHistori() {
	var kode string
	fmt.Println("Masukkan kode pegawai: ")
	fmt.Scan(&kode)
	found, indFound := binarySearch(kode)
	if found == false {
		fmt.Println("Kode pegawai yang ingin ditampilkan histori absensinya tidak ada")
	} else {
		fmt.Println("histori absensi dari pegawai yang dicari adalah :")
		fmt.Println(data.T[indFound].absen)
	}
}

func sortAscendKode() {
	a := 1
	for a < data.n {
		b := a - 1
		temp := data.T[a]
		for b >= 0 && (data.T[b].kode > temp.kode) {
			data.T[b+1] = data.T[b]
			b--
		}
		data.T[b+1] = temp
		a++
	}
}

func sortAscendNama() {
	a := 1
	for a < data.n {
		b := a - 1
		temp := data.T[a]
		for b >= 0 && (data.T[b].nama > temp.nama) {
			data.T[b+1] = data.T[b]
			b--
		}
		data.T[b+1] = temp
		a++
	}
	i := 0
	for i < data.n {
		fmt.Println(data.T[i])
		i++
	}
}

func reward() {
	var i int
	i = 0
	for i < data.n {

		if data.T[i].absen.jamMasuk == 8 && data.T[i].absen.jamKeluar == 16 {
			data.T[i].reward = "Menghargai waktu"
		}
		if data.T[i].absen.lembur >= 10 {
			data.T[i].reward = "Pekerja keras"

		}
		i++
	}
}

func tampilReward() {
	var i int
	for i < data.n {
		if data.T[i].reward == "pekerja keras" {
			fmt.Println(data.T[i].reward)
		}
		i++
	}
}
