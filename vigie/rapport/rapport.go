// Package rapport met en forme le rapport texte.
// Ses fonctions sont pures : elles construisent du texte sans l'afficher ni l'écrire.
package rapport

import (
	"fmt"
	"strings"
	"vigie/analyse"
)

// Formater construit le texte complet du rapport avec le seuil d'alerte donné.
func Formater(stats analyse.Stats, seuil int) string {
	result := []string(nil)
	result = append(result, "RAPPORT VIGIE")
	result = append(result, fmt.Sprintf("Événements analysés : %d", stats.Evenements))
	result = append(result, fmt.Sprintf("Lignes ignorées : %d", stats.Ignorees))
	result = append(result, fmt.Sprintf("Connexions réussies : %d", stats.Reussites))
	result = append(result, fmt.Sprintf("Connexions échouées : %d", stats.Echecs))
	result = append(result, "")
	result = append(result, "Échecs par adresse IP :")
	if len(stats.EchecsParIP) == 0 {
		result = append(result, "Aucun échec.")
	} else {
		comptes := analyse.Classer(stats.EchecsParIP)
		for _, c := range comptes {
			result = append(result, fmt.Sprintf("- %s : %d", c.Nom, c.Nombre))
		}
	}
	result = append(result, "")
	result = append(result, fmt.Sprintf("Alertes (au moins %d échecs) :", seuil))
	comptes := analyse.Classer(stats.EchecsParIP)
	nbAlertes := 0
	for _, c := range comptes {
		if c.Nombre >= seuil {
			result = append(result, fmt.Sprintf("- %s : %d", c.Nom, c.Nombre))
			nbAlertes++
		}
	}
	if nbAlertes == 0 {
		result = append(result, "Aucune alerte.")
	}
	return strings.Join(result, "\n") + "\n"
}
