### Service de Réservation en Ligne

## Structure du Code

Le projet est organisé en utilisant une strucutre simple et logique, les fonctions du package util sont organisés par fonctionnalités.

- main.go : Le point d'entrée de l'application. Il contient la fonction main qui initialise l'application et appelle la fonction principale pour démarrer le service de réservation en ligne.
- util/ : Ce répertoire contient tous les fichiers source de l'application, organisés par fonctionnalité.
  - run.go : Ce fichier contient la fonction Run() qui est responsable de l'exécution du menu principal de l'application.
  - room.go : Contient les fonctions relatives à la gestion des salles, y compris la fonction listRooms() pour afficher les salles disponibles.
  - reservation.go : Contient les fonctions liées à la création, à l'annulation et à la visualisation des réservations.
  - menu.go : Définit les menus principaux et secondaires de l'application ainsi que les fonctions associées pour interagir avec l'utilisateur.
  - db.go : Gère la connexion à la base de données MySQL, initialise les tables si elles n'existent pas, et contient des fonctions pour récupérer, créer et annuler des réservations.
  - classes.go : Définit les structures de données utilisées dans l'application, telles que Room et Reservation.

## Choix de conception

- Utilisation de Go : Imposé au projet.
- Organisation Modulaire : Le code est organisé de manière modulaire, avec chaque fonctionnalité de l'application placée dans un fichier séparé dans le répertoire util/.
- Utilisation de Packages Externes : Le projet utilise plusieurs packages externes. Le package github.com/go-sql-driver/mysql est utilisé pour interagir avec la base de données MySQL, et github.com/joho/godotenv pour charger les variables d'environnement à partir d'un fichier .env.
- Gestion des Erreurs : Une gestion d'erreur est effectuée dans tout le code, pour qu'en aucuns cas, il n'y ait d'erreurs forçant l'utilisateur à quitter l'application.
- Documentation Intégrée : Le code est commenté de manière à fournir une documentation claire sur son fonctionnement.

## Logique

- Liste des Salles Disponibles : La fonction listRooms() permet obtenir la liste des salles disponibles pour une plage horaire donnée et affiche les résultats à l'utilisateur.
- Création de Réservation : La fonction createReservationFromUserInput() demande à l'utilisateur de fournir les détails nécessaires pour créer une réservation, puis vérifie si la salle est disponible avant de créer la réservation.
- Annulation de Réservation : La fonction cancelReservationFromUserInput() permet à l'utilisateur d'annuler une réservation en fournissant l'identifiant de la réservation à annuler.
- Visualisation des Réservations : La fonction listReservations() affiche les réservations existantes pour une salle spécifique ou pour toutes les salles, en fonction de la demande de l'utilisateur.
- Exportation des Réservations en JSON : La fonction exportReservationsToJSON() exporte toutes les réservations existantes dans un fichier JSON pour une sauvegarde ou une utilisation ultérieure.
