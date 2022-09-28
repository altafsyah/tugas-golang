package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	fmt.Println("Halo")

	var dataTeman = map[string]Biodata{
		"1": Biodata{
			nama:      "Adam",
			pekerjaan: "Front End Developer",
			alamat:    "Bandung",
			alasan:    "Menambah ilmu",
		},
		"2": Biodata{
			nama:      "Andi",
			pekerjaan: "Back End Developer",
			alamat:    "Jakarta",
			alasan:    "Memperdalam skill",
		}, "3": Biodata{
			nama:      "Budi",
			pekerjaan: "Front End Developer",
			alamat:    "Bogor",
			alasan:    "Menambah ilmu",
		}, "4": Biodata{
			nama:      "John",
			pekerjaan: "QA Engineer",
			alamat:    "Pontianak",
			alasan:    "Menambah wawasan",
		}, "5": Biodata{
			nama:      "Tono",
			pekerjaan: "UI Designer",
			alamat:    "Yogyakarta",
			alasan:    "Mencari relasi",
		},
	}
	showBiodata(dataTeman[os.Args[1]])
}

func showBiodata(dataTeman Biodata) {
	fmt.Println("Nama : ", dataTeman.nama, "\nPekerjaan : ", dataTeman.pekerjaan, "\nAlamat : ", dataTeman.alamat, "\nAlasan : ", dataTeman.alasan)
}
