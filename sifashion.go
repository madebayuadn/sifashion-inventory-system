package main

import "fmt"

type baju struct {
	kode, nama, ukuran, warna string
	stok                      int
	harga                     float64
}
type tabBaju [100]baju

func readData(A *tabBaju, N *int) {
	var kode string
	fmt.Println("Kode Nama Ukuran Warna Stok Harga (ketik 'Selesai' untuk berhenti)")
	fmt.Scan(&kode)
	for kode != "Selesai" {
		A[*N].kode = kode
		fmt.Scan(&A[*N].nama, &A[*N].ukuran, &A[*N].warna, &A[*N].stok, &A[*N].harga)
		*N++
		fmt.Scan(&kode)
	}
}

func printData(A tabBaju, N int) {
	var i int
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-8s %-16s %-8s %-10s %-6s %-10s\n", "Kode", "Nama", "Ukuran", "Warna", "Stok", "Harga")
	fmt.Println("------------------------------------------------------------------")
	for i = 0; i < N; i++ {
		fmt.Printf("%-8s %-16s %-8s %-10s %-6d %-10.2f\n",
			A[i].kode, A[i].nama, A[i].ukuran, A[i].warna, A[i].stok, A[i].harga)
	}
	fmt.Println("------------------------------------------------------------------")
}

func ubahData(A *tabBaju, N int, x string) {
	var i int
	var found bool = false
	var konfirm string
	for i = 0; i < N; i++ {
		if A[i].nama == x {
			fmt.Printf("Ditemukan: %s %s %s - ubah? (y/n): ", A[i].kode, A[i].ukuran, A[i].warna)
			fmt.Scan(&konfirm)
			if konfirm == "y" {
				fmt.Print("Kode Nama Ukuran Warna Stok Harga (Baru): ")
				fmt.Scan(&A[i].kode, &A[i].nama, &A[i].ukuran, &A[i].warna, &A[i].stok, &A[i].harga)
				fmt.Println("Data berhasil diubah.")
			}
			found = true
		}
	}
	if !found {
		fmt.Println("Baju dengan nama", x, "tidak ditemukan.")
	}
}

func hapusData(A *tabBaju, N *int, x string) {
	var i, j int
	var found bool = false
	for i = 0; i < *N; i++ {
		if A[i].nama == x {
			for j = i; j < *N-1; j++ {
				A[j] = A[j+1]
			}
			*N--
			found = true
			i--
		}
	}
	if found {
		fmt.Println("Data dengan nama", x, "berhasil dihapus.")
	} else {
		fmt.Println("Baju dengan nama", x, "tidak ditemukan.")
	}
}

func insertionSortHargaAscending(A *tabBaju, N int) {
	var i, pass int
	var temp baju
	for pass = 1; pass < N; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.harga < A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func insertionSortHargaDescending(A *tabBaju, N int) {
	var i, pass int
	var temp baju
	for pass = 1; pass < N; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.harga > A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func insertionSortWarna(A *tabBaju, N int) {
	var i, pass int
	var temp baju
	for pass = 1; pass < N; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.warna < A[i-1].warna {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

func selectionSortStokAscending(A *tabBaju, N int) {
	var i, j, minIdx int
	var temp baju
	for i = 0; i < N-1; i++ {
		minIdx = i
		for j = i + 1; j < N; j++ {
			if A[j].stok < A[minIdx].stok {
				minIdx = j
			}
		}
		temp = A[i]
		A[i] = A[minIdx]
		A[minIdx] = temp
	}
}

func selectionSortStokDescending(A *tabBaju, N int) {
	var i, j, minIdx int
	var temp baju
	for i = 0; i < N-1; i++ {
		minIdx = i
		for j = i + 1; j < N; j++ {
			if A[j].stok > A[minIdx].stok {
				minIdx = j
			}
		}
		temp = A[i]
		A[i] = A[minIdx]
		A[minIdx] = temp
	}
}

func sequentialSearchUkuran(A tabBaju, N int, x string) {
	var i int
	var found bool = false
	fmt.Println("-- Hasil Sequential Search Ukuran:", x, "--")
	for i = 0; i < N; i++ {
		if A[i].ukuran == x {
			fmt.Printf("%-8s %-16s %-8s %-10s %-6d %-10.2f\n",
				A[i].kode, A[i].nama, A[i].ukuran, A[i].warna, A[i].stok, A[i].harga)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan baju dengan ukuran", x)
	}
}

func binarySearchWarna(A tabBaju, N int, x string) {
	var left, right, mid, start, end, i int
	var found bool = false
	left = 0
	right = N - 1
	fmt.Println("-- Hasil Binary Search Warna:", x, "--")
	for left <= right && !found {
		mid = (left + right) / 2
		if x < A[mid].warna {
			right = mid - 1
		} else if x > A[mid].warna {
			left = mid + 1
		} else {
			start = mid
			for start > 0 && A[start-1].warna == x {
				start--
			}
			end = mid
			for end < N-1 && A[end+1].warna == x {
				end++
			}
			for i = start; i <= end; i++ {
				fmt.Printf("%-8s %-16s %-8s %-10s %-6d %-10.2f\n",
					A[i].kode, A[i].nama, A[i].ukuran, A[i].warna, A[i].stok, A[i].harga)
			}
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan baju dengan warna", x)
	}
}

func statistik(A tabBaju, N int) {
	var i, totalStok, iMaks, iMin, j, total int
	var duplikat bool

	for i = 0; i < N; i++ {
		totalStok += A[i].stok
		if A[i].stok > A[iMaks].stok {
			iMaks = i
		}
		if A[i].stok < A[iMin].stok {
			iMin = i
		}
	}
	fmt.Println("-- Statistik --")
	fmt.Println("Total produk   :", N)
	fmt.Println("Total stok     :", totalStok)
	fmt.Printf("Stok terbanyak : %s (%s) - %d pcs\n", A[iMaks].kode, A[iMaks].nama, A[iMaks].stok)
	fmt.Printf("Stok tersedikit: %s (%s) - %d pcs\n", A[iMin].kode, A[iMin].nama, A[iMin].stok)

	fmt.Println("\n-- Total Stok per Ukuran --")
	var ukuranSudah [100]string
	var ukuranCount int = 0
	for i = 0; i < N; i++ {
		duplikat = false
		for j = 0; j < ukuranCount; j++ {
			if ukuranSudah[j] == A[i].ukuran {
				duplikat = true
			}
		}
		if !duplikat {
			ukuranSudah[ukuranCount] = A[i].ukuran
			ukuranCount++
		}
	}
	for i = 0; i < ukuranCount; i++ {
		total = 0
		for j = 0; j < N; j++ {
			if A[j].ukuran == ukuranSudah[i] {
				total += A[j].stok
			}
		}
		fmt.Printf("Ukuran %-4s : %d pcs\n", ukuranSudah[i], total)
	}

	fmt.Println("\n-- Total Stok per Warna --")
	var warnaSudah [100]string
	var warnaCount int = 0
	for i = 0; i < N; i++ {
		duplikat = false
		for j = 0; j < warnaCount; j++ {
			if warnaSudah[j] == A[i].warna {
				duplikat = true
			}
		}
		if !duplikat {
			warnaSudah[warnaCount] = A[i].warna
			warnaCount++
		}
	}
	for i = 0; i < warnaCount; i++ {
		total = 0
		for j = 0; j < N; j++ {
			if A[j].warna == warnaSudah[i] {
				total += A[j].stok
			}
		}
		fmt.Printf("Warna %-10s : %d pcs\n", warnaSudah[i], total)
	}
}

func main() {
	var A tabBaju
	var N int = 0
	var pilihan, psort string
	var sortedA tabBaju
	var selesai bool = false

	for !selesai {
		fmt.Println("\n===== SIFASHION =====")
		fmt.Println("1. Input Data Baju")
		fmt.Println("2. Tampilkan Semua Data")
		fmt.Println("3. Ubah Data by Nama")
		fmt.Println("4. Hapus Data by Nama")
		fmt.Println("5. Insertion Sort Harga")
		fmt.Println("6. Selection Sort Stok")
		fmt.Println("7. Sequential Search by Ukuran")
		fmt.Println("8. Binary Search by Warna")
		fmt.Println("9. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case "0":
			fmt.Println("Terima kasih!")
			selesai = true
		case "1":
			readData(&A, &N)
		case "2", "3", "4", "5", "6", "7", "8", "9":
			if N == 0 {
				fmt.Println("Belum ada data. Silakan input data terlebih dahulu.")
			} else {
				switch pilihan {
				case "2":
					printData(A, N)
				case "3":
					var nama string
					fmt.Print("Nama baju yang ingin diubah: ")
					fmt.Scan(&nama)
					ubahData(&A, N, nama)
				case "4":
					var nama string
					fmt.Print("Nama baju yang ingin dihapus: ")
					fmt.Scan(&nama)
					hapusData(&A, &N, nama)
				case "5":
					sortedA = A
					fmt.Println("1. Sort Harga Ascending")
					fmt.Println("2. Sort Harga Descending")
					fmt.Print("Pilihan: ")
					fmt.Scan(&psort)

					switch psort {
					case "1":
						insertionSortHargaAscending(&sortedA, N)
						fmt.Println("-- Hasil Sort Harga Secara Ascending --")
						printData(sortedA, N)
					case "2":
						insertionSortHargaDescending(&sortedA, N)
						fmt.Println("-- Hasil Sort Harga Secara Descending --")
						printData(sortedA, N)
					default:
						fmt.Println("Pilihan Tidak Valid")
					}

				case "6":
					sortedA = A
					fmt.Println("1. Sort Stok Ascending")
					fmt.Println("2. Sort Stok Descending")
					fmt.Print("Pilihan: ")
					fmt.Scan(&psort)

					switch psort {
					case "1":
						selectionSortStokAscending(&sortedA, N)
						fmt.Println("-- Hasil Sort Stok Secara Ascending --")
						printData(sortedA, N)
					case "2":
						selectionSortStokDescending(&sortedA, N)
						fmt.Println("-- Hasil Sort Stok Secara Descending --")
						printData(sortedA, N)
					default:
						fmt.Println("Pilihan Tidak Valid")
					}
				case "7":
					var ukuran string
					fmt.Print("Cari ukuran: ")
					fmt.Scan(&ukuran)
					sequentialSearchUkuran(A, N, ukuran)
				case "8":
					var warna string
					sortedA = A
					insertionSortWarna(&sortedA, N)
					fmt.Print("Cari warna: ")
					fmt.Scan(&warna)
					binarySearchWarna(sortedA, N, warna)
				case "9":
					statistik(A, N)
				}
			}
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
