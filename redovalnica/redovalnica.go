/*
	Package redovalnica vsebuje funkcije za upravljane z ocenami študentov

	Izvaža naslednje funkcije:
		- DodajOceno: doda študentu novo oceno
		- IzpisVsehOcen: izpiše vse študente in vse njihove ocene
		- IzpisiKoncniUspeh: izračuna in izpiše končni uspeh vseh študentov

	Pomožne spremenljivke:
		- MinOcena: minimanla možna ocena
		- MaxOcena: maksimalna možna ocena
		- StOcen: najmanjše število ocen potrebnih za pozitivno oceno

	Skrita funkcija:
		- povprecje: izračuna povprečje ocen posameznega študenta
		- najdiVpisno: najde vpisno številko študenta
*/
package redovalnica

import (
	"fmt"
)

// Student predstavlja študenta z imenom, priimkom in ocenami
type Student struct {
    Ime     string
    Priimek string
    Ocene   []int
}

var (
	MinOcena = 1
	MaxOcena = 10
	StOcen = 6
)

// DodajOceno doda oceno izbranemu študentu
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		fmt.Println("Študent s vprisno številko", vpisnaStevilka, "ne obstaja!")
		return
	}

	if ocena < MinOcena || ocena > MaxOcena {
		fmt.Println("Ocena", ocena, "ni veljavna!")
		return
	}

	student.Ocene = append(student.Ocene, ocena)

	studenti[vpisnaStevilka] = student
	fmt.Println("Študentu", vpisnaStevilka, "je bila dodana ocena", ocena)
}

// IzpisVsehOcen izpiše vse študente in njihove ocene
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s %s: %v\n", vpisna, student.Ime, student.Priimek, student.Ocene)
	}
}

// IzpisiKoncniUspeh izpiše končni uspeh za vsakega študenta
func IzpisiKoncniUspeh(studenti map[string]Student) {
	fmt.Println("\nKONČNI USPEH:")
	for _, student := range studenti {
		p := povprecje(studenti, najdiVpisno(studenti, student))
		if p == -1.0 {
			fmt.Printf("%s %s: ni podatkov o študentu\n", student.Ime, student.Priimek)
			continue
		}
		fmt.Printf("%s %s: povprečna ocena %.1f -> ", student.Ime, student.Priimek, p)
		if p >= 9 {
			fmt.Println("Odličen študent!")
		} else if p >= 6 {
			fmt.Println("Povprečen študent")
		} else {
			fmt.Println("Neuspešen študent")
		}
	}
}

// povprecje izračuna povprečno oceno
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		return -1.0
	}

	if len(student.Ocene) < StOcen {
		return 0.0
	}

	vsota := 0
	for _, ocena := range student.Ocene {
		vsota += ocena
	}

	return float64(vsota) / float64(len(student.Ocene))
}

// najde vpisno številko od študenta
func najdiVpisno(studenti map[string]Student, cilj Student) string {
	for vpisna, s := range studenti {
		if s.Ime == cilj.Ime && s.Priimek == cilj.Priimek {
			return vpisna
		}
	}
	return ""
}