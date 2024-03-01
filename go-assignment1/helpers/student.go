package helpers

import "fmt"

type Student struct {
	noAbsen   string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

// Generate data, global variable, map of Student struct
var students = map[string]Student{
	"1": {
		noAbsen:   "1",
		nama:      "Komang Adyanata",
		alamat:    "Denpasar",
		pekerjaan: "Programmer",
		alasan:    "Tertarik untuk belajar dan mengimplementasikan bahasa pemrograman Go",
	},
	"2": {
		noAbsen:   "2",
		nama:      "Nama2",
		alamat:    "Alamat2",
		pekerjaan: "Pekerjaan2",
		alasan:    "Alasan2",
	},
	"3": {
		noAbsen:   "3",
		nama:      "Nama3",
		alamat:    "Alamat3",
		pekerjaan: "Pekerjaan3",
		alasan:    "Alasan3",
	},
	"4": {
		noAbsen:   "4",
		nama:      "Nama4",
		alamat:    "Alamat4",
		pekerjaan: "Pekerjaan4",
		alasan:    "Alasan4",
	},
	"5": {
		noAbsen:   "5",
		nama:      "Nama5",
		alamat:    "Alamat5",
		pekerjaan: "Pekerjaan5",
		alasan:    "Alasan5",
	},
}

func AddStudent(p_NoAbsen string, p_Nama string, p_Alamat string, p_Pekerjaan string, p_Alasan string) {
	students[p_NoAbsen] = Student{
		noAbsen:   p_NoAbsen,
		nama:      p_Nama,
		alamat:    p_Alamat,
		pekerjaan: p_Pekerjaan,
		alasan:    p_Alasan,
	}
}

func DeleteStudent(p_Key string) {
	delete(students, p_Key)
}

func printStudent(p_Student Student) {
	fmt.Println("No. Absen\t:", p_Student.noAbsen)
	fmt.Println("Nama\t\t:", p_Student.nama)
	fmt.Println("Alamat\t\t:", p_Student.alamat)
	fmt.Println("Pekerjaan\t:", p_Student.pekerjaan)
	fmt.Println("Alasan\t\t:", p_Student.alasan)
	fmt.Printf("\n\n")
}

func GetDetail(inputCli string, isPrintAll bool) {
	if !isPrintAll {
		value, exist := students[inputCli]
		if exist {
			printStudent(value)
		} else {
			fmt.Println("Absen", inputCli, "not Exist")
		}
	} else {
		//Loop all students
		for key, value := range students {
			fmt.Println("Key -------", key)
			printStudent(value)
		}
	}

}
