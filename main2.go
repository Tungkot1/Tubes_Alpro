package main
import (
	"fmt"
	"sort"
	"strings"
)
type Mahasiswa struct {
	NIM   string
	Nama  string
	Prodi string
}
type MataKuliah struct {
	Kode string
	Nama string
	SKS  int
}

type KRS struct {
	NIM     string
	KodeMKs []string
}

type Nilai struct {
	NIM        string
	KodeMK     string
	NilaiHuruf string
}

var daftarMahasiswa []Mahasiswa
var daftarMatkul []MataKuliah
var daftarKRS []KRS
var daftarNilai []Nilai

func main() {
	isiDataMataKuliah()

	for {
		fmt.Println("\n=== Aplikasi Dosen Wali ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tambah Nilai")
		fmt.Println("3. Lihat Daftar Mahasiswa")
		fmt.Println("4. Cari Mahasiswa (NIM/Nama)")
		fmt.Println("5. Urutkan Mahasiswa (IPK/Nama)")
		fmt.Println("6. Tambah Mata Kuliah ke Mahasiswa")
		fmt.Println("7. Hapus Mata Kuliah dari Mahasiswa")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			tambahMahasiswa()
		case 2:
			tambahMatkul()
		case 3:
			tambahNilai()
		case 4:
		case 3:
			tampilkanMahasiswa()
		case 5:
		case 4:
			cariMahasiswa()
		case 6:
		case 5:
			urutkanMahasiswa()
		case 6:
			tambahMatkulMahasiswa()
		case 7:
			hapusMatkulMahasiswa()
		case 8:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func isiDataMataKuliah() {
	daftarMatkul = []MataKuliah{
		{"IF101", "Algoritma dan Pemrograman", 3},
		{"IF102", "Struktur Data", 2},
		{"IF103", "Basis Data", 2},
		{"IF104", "Etika Ai", 2},
	}
}

func tambahMahasiswa() {
	var nim, nama, prodi string

	fmt.Print("NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Prodi: ")
	fmt.Scanln(&prodi)
	daftarMahasiswa = append(daftarMahasiswa, Mahasiswa{NIM: nim, Nama: nama, Prodi: prodi})
	fmt.Println("Mahasiswa ditambahkan.")
}

func tambahMatkul() {
	var kode, nama string
	var sks int

	fmt.Print("Kode MK: ")
	fmt.Scanln(&kode)

	fmt.Print("Nama MK: ")
	fmt.Scanln(&nama)

	fmt.Print("SKS: ")
	fmt.Scanln(&sks)

	daftarMatkul = append(daftarMatkul, MataKuliah{Kode: kode, Nama: nama, SKS: sks})
	fmt.Println("Mata kuliah ditambahkan.")
}

func tambahNilai() {
	var nim, kode, nilai string

	fmt.Print("NIM Mahasiswa: ")
	fmt.Scanln(&nim)
	fmt.Print("Kode MK: ")
	fmt.Scanln(&kode)
	fmt.Print("Nilai Huruf (A/B/C/D/E): ")
	fmt.Scanln(&nilai)
	nilai = strings.ToUpper(nilai)
	daftarNilai = append(daftarNilai, Nilai{NIM: nim, KodeMK: kode, NilaiHuruf: nilai})
	fmt.Println("Nilai ditambahkan.")
}
func konversiNilai(huruf string) float64 {
	switch huruf {
	case "A":
		return 4.0
	case "B":
		return 3.0
	case "C":
		return 2.0
	case "D":
		return 1.0
	default:
		return 0.0
	}
}
func hitungIPK(nim string) float64 {
	totalSKS := 0
	totalNilai := 0.0
	for _, n := range daftarNilai {
		if n.NIM == nim {
			for _, mk := range daftarMatkul {
				if mk.Kode == n.KodeMK {
					bobot := konversiNilai(n.NilaiHuruf)
					totalNilai += bobot * float64(mk.SKS)
					totalSKS += mk.SKS
				}
			}
		}
	}
	if totalSKS == 0 {
		return 0.0
	}
	return totalNilai / float64(totalSKS)
}
func tampilkanMahasiswa() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}
	fmt.Println("\nDaftar Mahasiswa:")
	for _, m := range daftarMahasiswa {
		fmt.Printf("%s - %s - %s - IPK: %.2f\n", m.NIM, m.Nama, m.Prodi, hitungIPK(m.NIM))
		tampilkanMatkulMahasiswa(m.NIM)
	}
}

func cariMahasiswa() {
	var keyword string
	fmt.Print("Masukkan NIM atau Nama: ")
	fmt.Scanln(&keyword)
	keyword = strings.ToLower(keyword)
	found := false
	for _, m := range daftarMahasiswa {
		if strings.Contains(strings.ToLower(m.NIM), keyword) || strings.Contains(strings.ToLower(m.Nama), keyword) {
			fmt.Printf("%s - %s - %s - IPK: %.2f\n", m.NIM, m.Nama, m.Prodi, hitungIPK(m.NIM))
			tampilkanMatkulMahasiswa(m.NIM)
			found = true
		}
	}
	if !found {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}
}
func urutkanMahasiswa() {
	var pilihan int
	fmt.Println("Urut berdasarkan:")
	fmt.Println("1. IPK (Tinggi ke Rendah)")
	fmt.Println("2. Nama (A-Z)")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		sort.Slice(daftarMahasiswa, func(i, j int) bool {
			return hitungIPK(daftarMahasiswa[i].NIM) > hitungIPK(daftarMahasiswa[j].NIM)
		})
		fmt.Println("Data diurutkan berdasarkan IPK.")
	case 2:
		sort.Slice(daftarMahasiswa, func(i, j int) bool {
			return daftarMahasiswa[i].Nama < daftarMahasiswa[j].Nama
		})
		fmt.Println("Data diurutkan berdasarkan Nama.")
	default:
		fmt.Println("Pilihan tidak valid.")
	}

	tampilkanMahasiswa()
}

func tampilkanMatkulMahasiswa(nim string) {
	totalSKS := 0
	for _, krs := range daftarKRS {
		if krs.NIM == nim {
			fmt.Println("  Mata kuliah yang diambil:")
			for _, kode := range krs.KodeMKs {
				for _, mk := range daftarMatkul {
					if mk.Kode == kode {
						fmt.Printf("    - %s (%s) - %d SKS\n", mk.Nama, mk.Kode, mk.SKS)
						totalSKS += mk.SKS
					}
				}
			}
			fmt.Printf("  Total SKS: %d\n", totalSKS)
			break
		}
	}
}

func tambahMatkulMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scanln(&nim)

	found := false
	for _, m := range daftarMahasiswa {
		if m.NIM == nim {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	tampilkanMatkul()
	fmt.Println("Pilih mata kuliah yang ingin ditambahkan (pisahkan dengan koma): ")
	var input string
	fmt.Scanln(&input)
	kodeList := strings.Split(input, ",")
	for i := range kodeList {
		kodeList[i] = strings.ToUpper(strings.TrimSpace(kodeList[i]))
	}

	for i, krs := range daftarKRS {
		if krs.NIM == nim {
			for _, kode := range kodeList {
				if !contains(krs.KodeMKs, kode) && mataKuliahAda(kode) {
					daftarKRS[i].KodeMKs = append(daftarKRS[i].KodeMKs, kode)
				}
			}
			fmt.Println("Mata kuliah berhasil ditambahkan.")
			return
		}
	}

	var mkBaru []string
	for _, kode := range kodeList {
		if mataKuliahAda(kode) {
			mkBaru = append(mkBaru, kode)
		}
	}
	daftarKRS = append(daftarKRS, KRS{NIM: nim, KodeMKs: mkBaru})
	fmt.Println("Mata kuliah berhasil ditambahkan.")
}

func hapusMatkulMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scanln(&nim)

	index := -1
	for i, krs := range daftarKRS {
		if krs.NIM == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Mahasiswa tidak ditemukan atau belum memiliki mata kuliah.")
		return
	}

	fmt.Println("Mata kuliah yang diambil:")
	for _, kode := range daftarKRS[index].KodeMKs {
		for _, mk := range daftarMatkul {
			if mk.Kode == kode {
				fmt.Printf("  - %s (%s)\n", mk.Nama, mk.Kode)
			}
		}
	}

	fmt.Print("Masukkan kode matkul yang ingin dihapus: ")
	var kodeHapus string
	fmt.Scanln(&kodeHapus)
	kodeHapus = strings.ToUpper(strings.TrimSpace(kodeHapus))

	baru := []string{}
	terhapus := false
	for _, kode := range daftarKRS[index].KodeMKs {
		if kode != kodeHapus {
			baru = append(baru, kode)
		} else {
			terhapus = true
		}
	}
	daftarKRS[index].KodeMKs = baru

	if terhapus {
		fmt.Println("Mata kuliah berhasil dihapus.")
	} else {
		fmt.Println("Kode mata kuliah tidak ditemukan.")
	}
}

func tampilkanMatkul() {
	fmt.Println("\nDaftar Mata Kuliah:")
	for _, mk := range daftarMatkul {
		fmt.Printf("- %s: %s (%d SKS)\n", mk.Kode, mk.Nama, mk.SKS)
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func mataKuliahAda(kode string) bool {
	for _, mk := range daftarMatkul {
		if mk.Kode == kode {
			return true
		}
	}
	return false
}
