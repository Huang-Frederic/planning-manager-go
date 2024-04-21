### Service de Réservation en Ligne

Ce service de réservation en ligne est une application permettant de gérer des réservations de salles. L'application offre des fonctionnalités telles que la liste des salles disponibles, la création, l'annulation et la visualisation des réservations. Le projet est écrit en langage Go et utilise une base de données MySQL pour stocker les données.

## Table des matières

1. Fonctionnalités
2. Installation
3. Configuration
4. Utilisation
5. Structure du projet
6. Dépendances

# Fonctionnalités du menu

- Lister les salles disponibles : Créez une fonction qui affiche les salles qui ne sont pas réservées pour une plage horaire spécifique.
- Créer une réservation : Demandez à l'utilisateur de fournir les détails nécessaires pour une réservation (nom de la salle, date, heure) et vérifiez si la salle est disponible avant de créer la réservation.
- Annuler une réservation : Demandez à l'utilisateur de fournir les détails nécessaires pour annuler une réservation (nom de la salle, date, heure) et vérifiez si la réservation existe avant de l'annuler.
- Visualiser les réservations : Affichez les réservations existantes pour une salle spécifique, permettez de filtrer pour une date spécifique.
- Exporter les réservations : Exportez les réservations existantes dans un fichier json qui sera localisé dans le dossier "export_datas".

# Installation

- Go installé sur votre système. Si ce n'est pas le cas, téléchargez et installez-le à partir du site officiel de Go.
- Une base de données MySQL.

Etapes d'installation

- git clone <URL_DU_REPO>
- cd <NOM_DU_REPOT>
- go mod tidy

# Configuration

- Créez un fichier .env à la racine du projet et spécifiez les paramètres de connexion à votre base de données. Voici un exemple de contenu pour le fichier .env :

DBHost="localhost"
DBPort="3306"
DBUser="go_user"
DBPassword="azerty"
DBName="go_sql"

# Utilisation

- go run main.go

# Structure du projet

.
├── README.md
├── main.go
├── go.mod
├── go.sum
├── .env
├── export_datas
└── util
. ├── run.go
. ├── room.go
. ├── reservation.go
. ├── menu.go
. ├── db.go
. ├── classes.go

- main.go : Fichier principal contenant la fonction main.
- go.mod et go.sum : Fichiers de configuration Go.
- .env : Fichier de configuration pour les paramètres de la base de données.
- util/ : Répertoire contenant les fichiers de code source de l'application.
  -> run.go, room.go, reservation.go, menu.go, db.go, classes.go, export.go : Fichiers contenant la logique métier de l'application.

# Dépendances

Ce projet utilise les dépendances suivantes :

- github.com/go-sql-driver/mysql
- github.com/joho/godotenv

# Répartition

Frédéric Huang : .%
Franck Zhuang : .%
