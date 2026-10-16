// Package parser découpe une ligne du journal en champs et vérifie qu'elle est valide.
package parser

import (
	"errors"
	"fmt"
	"strings"
)

// ErrLigneInvalide signale une ligne qui ne respecte pas le format attendu.
var ErrLigneInvalide = errors.New("ligne invalide")

// NbChamps est le nombre de champs d'une ligne du journal.
const NbChamps = 5

// Evenement représente une tentative de connexion.
type Evenement struct {
	Date, Heure, Resultat, Utilisateur, IP string
}

// Analyser découpe une ligne et renvoie un Evenement,
// ou une erreur qui enveloppe ErrLigneInvalide.
func Analyser(ligne string) (Evenement, error) {
	champs := strings.Fields(ligne)
	if len(champs) != NbChamps {
		err := fmt.Errorf("%w : %d champs au lieu de %d", ErrLigneInvalide, len(champs), NbChamps)
		return Evenement{}, err
	}

	champs[2] = strings.ToUpper(champs[2])
	if champs[2] != "OK" && champs[2] != "FAIL" {
		err := fmt.Errorf("%w : résultat inconnu %q", ErrLigneInvalide, champs[2])
		return Evenement{}, err
	}

	result := Evenement{champs[0], champs[1], champs[2], champs[3], champs[4]}
	return result, nil
}

// Journal analyse toutes les lignes du texte.
// Les lignes vides sont sautées sans avertissement ;
// les lignes invalides sont écartées et signalées dans avertissements.
func Journal(texte string) (evenements []Evenement, avertissements []string) {
	for i, ligne := range strings.Split(texte, "\n") {
		verif := strings.TrimSpace(ligne)
		if verif != "" {
			continue
		}
		eve, err := Analyser(ligne)
		if err != nil {
			erreur := fmt.Sprintf("ligne %d ignorée : %v", i, err)
			avertissements = append(avertissements, erreur)
			continue
		}
		evenements = append(evenements, eve)
	}
	return evenements, avertissements
}
