// Package analyse calcule les statistiques du journal.
// Ses fonctions sont pures : elles ne lisent ni n'écrivent aucun fichier et n'affichent rien.
package analyse

import (
	"sort"
	"vigie/parser"
)

// Stats regroupe les comptages du journal.
type Stats struct {
	Evenements  int            // lignes valides
	Ignorees    int            // lignes invalides
	Reussites   int            // résultat OK
	Echecs      int            // résultat FAIL
	EchecsParIP map[string]int // adresse IP -> nombre d'échecs
}

// Compte associe un nom (adresse IP ou utilisateur) à un nombre d'échecs.
type Compte struct {
	Nom    string
	Nombre int
}

// Calculer compte les réussites, les échecs et les échecs par adresse IP.
func Calculer(evenements []parser.Evenement, ignorees int) Stats {
	Eve := Stats{0, ignorees, 0, 0, map[string]int{}}
	for _, res := range evenements {
		Eve.Evenements++
		if res.Resultat == "OK" {
			Eve.Reussites++
		}
		if res.Resultat == "FAIL" {
			Eve.Echecs++
			Eve.EchecsParIP[res.IP]++
		}
	}
	return Eve
}

// Classer transforme la map en tranche triée par nombre décroissant,
// puis par nom (ordre alphabétique) en cas d'égalité.
func Classer(nombres map[string]int) []Compte {
	class := make([]Compte, 0)
	perso := Compte{}
	for ip, nombre := range nombres {
		perso.Nom = ip
		perso.Nombre = nombre
		class = append(class, perso)
	}
	sort.Slice(class, func(i, j int) bool {
		return class[i].Nom < class[j].Nom
	})
	sort.Slice(class, func(i, j int) bool {
		return class[i].Nombre > class[j].Nombre
	})

	return class
}
