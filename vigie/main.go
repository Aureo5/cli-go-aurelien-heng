// Vigie - analyseur de journal de connexions du campus
package main

import (
	"flag"
	"fmt"
	"os"
	"vigie/analyse"
	"vigie/fileio"
	"vigie/parser"
	"vigie/rapport"
)

// Exercice 7 : c'est ici que vous assemblerez le programme.
func main() {
	seuil := flag.Int("seuil", 3, "seuil d'erreur")
	flag.Parse()
	err := lancer(flag.Args(), *seuil)
	if err != nil {
		fmt.Println("vigie : ", err)
		os.Exit(1)
	}
}

func lancer(args []string, seuil int) error {
	if len(args) != 2 {
		return fmt.Errorf("usage : go run . [-seuil N] <journal> <rapport>)")
	}
	if args[0] == args[1] {
		return fmt.Errorf("Les fichiers doivent être différent.")
	}
	if seuil < 1 {
		return fmt.Errorf("Seuil doit être supérieur a 1")
	}
	entree := args[0]
	sortie := args[1]
	contenue, err := fileio.Lire(entree)
	if err != nil {
		return err
	}
	Eve, aver := parser.Journal(contenue)
	if aver != nil {
		return fmt.Errorf("attention : %q", aver)
	}
	stat := analyse.Calculer(Eve, len(aver))
	fileio.Ecrire(sortie, rapport.Formater(stat, seuil))
	fmt.Printf("Rapport écrit : %q", sortie)
	return nil
}
