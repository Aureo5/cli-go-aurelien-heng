// Package fileio lit et écrit les fichiers texte.
package fileio

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Lire renvoie le contenu du fichier sous forme de texte.
// Fichier absent : erreur "fichier introuvable : <chemin>".
// Autre erreur : "lecture de <chemin> : <erreur d'origine>" (avec %w).
func Lire(chemin string) (string, error) {
	texte, err := os.ReadFile(chemin)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = fmt.Errorf("fichier introuvable : %s", chemin)
			return "", err
		}
		err = fmt.Errorf("Lecture de %s : %w", chemin, err)
		return "", err
	}
	return string(texte), nil
}

// Ecrire écrit le texte dans le fichier (créé ou remplacé).
// En cas d'erreur : "écriture de <chemin> : <erreur d'origine>" (avec %w).
func Ecrire(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("écriture de %s : %w", chemin, err)
	}
	return nil
}
