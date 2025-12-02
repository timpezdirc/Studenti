package main

import (
	"log"
	"os"
	"strconv"

	r "github.com/timpezdirc/Studenti/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	studenti := make(map[string]Student)

	studenti["63220000"] = Student{ime: "Janez", priimek: "Novak", ocene: []int{5, 6, 7, 6, 8}}
	studenti["63220001"] = Student{ime: "Tine", priimek: "Kranjc", ocene: []int{10, 10, 8, 9, 9}}
	studenti["63220002"] = Student{ime: "Ana", priimek: "Kovač", ocene: []int{10, 6, 7, 6, 9}}

	app := &cli.App{
		Name:  "redovalnica",
		Usage: "Upravljanje ocen študentov",

		Flags: []cli.Flag{
			&cli.IntFlag{Name: "stOcen", Value: 6, Destination: &r.StOcen},
			&cli.IntFlag{Name: "minOcena", Value: 1, Destination: &r.MinOcena},
			&cli.IntFlag{Name: "maxOcena", Value: 10, Destination: &r.MaxOcena},
		},

		Commands: []*cli.Command{
			{
				Name:  "dodaj",
				Usage: "Dodaj oceno študentu",
				Action: func(ctx *cli.Context) error {
					vpisna := ctx.Args().Get(0)
					ocenaStr := ctx.Args().Get(1)

					ocena, err := strconv.Atoi(ocenaStr)
					if err != nil {
						return err
					}

					return r.DodajOceno(studenti, vpisna, ocena)
				},
			},
			{
				Name:  "izpisi",
				Usage: "Izpiše vse ocene",
				Action: func(ctx *cli.Context) error {
					r.IzpisVsehOcen(studenti)
					return nil
				},
			},
			{
				Name:  "uspeh",
				Usage: "Izpiše končni uspeh študentov",
				Action: func(ctx *cli.Context) error {
					r.IzpisiKoncniUspeh(studenti)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}