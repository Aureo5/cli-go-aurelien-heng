# Vigie

Outil d'écriture de rapport en golang dans le terminal.
Créé un fichier avec le rapport d'écriture des tentative de connexions au serveurs. 

# Installation

Prérequis : Go 1.22 ou plus récent (`go version`).
```
git clone https://github.com/Aureo5/cli-go-aurelien-heng.git
```
puis
```
cd cli-go-aurelien-heng/projet
```

# Utilisation

```
go run . (fichier d'entrée) (nom de fichier corriger voulu)
```
ou si vous voulez personnaliser le seuil d'érreure autorisée:
```
go run . -seuil (nombre autorisé) (fichier d'entrée) (nom de fichier corriger voulu)
```
Exemple: 
```
go run . exemple.txt sortie.txt
```
exemple.txt :
```

2026-10-14 12:00:00 FAIL admin 10.0.0.5

2026-10-14 12:00:30 OK alice 192.168.1.20
```
____________________________________________________________________________________________________
sortie.txt :
```
RAPPORT VIGIE
Événements analysés : 2
Lignes ignorées : 0
Connexions réussies : 1
Connexions échouées : 1

Échecs par adresse IP :
- 10.0.0.5 : 1

Alertes (au moins 3 échecs) :
Aucune alerte.
```
Notes :
- Toute ligne vide ne seront pas pris en compte.

## Gestion des erreurs
Les messages d'erreur sont affichés sur la sortie d'erreur et le programme renvoie le code 1 dans les cas suivants :

    arguments manquants
    fichier introuvable
    fichier d'entrée identique au fichier de sortie

# Architecture des fichiers
```
.
├── vigie/
│   ├── analyse/
│   │   ├── analyse.go
│   │   └── analyse_test.go
│   ├── fileio/
│   │   ├── fileio.go
│   │   └── fileio_test.go
│   ├── parser/
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── rapport/
│   │   ├── rapport.go
│   │   └── rapport_test.go
│   ├── tests/
│   ├── go.mod
│   ├── journal.txt
│   └── LISEZMOI.txt
├── main.go
├── ouvrir.bat
├── ouvrir.command
├── ouvrir.sh
├── rapport.txt
└── README.md
```
# Test

Le fichier test contient des fichier pour montré certaines fonctionalitées.
2 types de fichiers par test: entree.txt et attendu.txt.
Lancé l'outil avec l'entrée et comparée (fc.exe) le fichier de sortie avec attendu.txt comme ci-dessous.
```
go run . entree.txt sortie.txt
fc.exe attendu.txt sortie.txt
```
# Credit

Aurélien Heng
