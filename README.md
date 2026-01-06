# Le Vélo Vert - Site Web de Location de Vélos

## Description

Le Vélo Vert est une application web développée en Go pour une entreprise de location de vélos écologiques. Le site permet aux utilisateurs de :

- Découvrir les différents types de vélos disponibles (VTT, Ville, Électrique, Enfant)
- Explorer les itinéraires touristiques de la région
- Consulter les actualités sur la mobilité verte
- Accéder à des conseils de sécurité
- Réserver en ligne des vélos

## Fonctionnalités

### 🏠 Page d'accueil
- Présentation des vélos populaires
- Aperçu des itinéraires touristiques
- Actualités récentes
- Navigation intuitive

### 🚴‍♂️ Catalogue de vélos
- 4 types de vélos : VTT, Ville, Électrique, Enfant
- Prix par jour
- Statut de disponibilité
- Réservation en ligne

### 🗺️ Itinéraires touristiques
- 3 parcours thématiques
- Informations détaillées (distance, durée, difficulté)
- Points d'intérêt
- Images illustratives

### 📰 Actualités
- Événements locaux liés à la mobilité verte
- Festival du Vélo Vert
- Nouveaux parcours cyclables

### 📚 Conseils de sécurité
- Conseils par catégories (Sécurité, Entretien, Équipement)
- Informations pratiques pour une utilisation sécurisée

### 📝 Système de réservation
- Formulaire de réservation complet
- Calcul automatique du prix
- Confirmation par email simulée
- Gestion des données de réservation

## Architecture Technique

### Structure du projet
```
velo-vert/
├── main.go                 # Point d'entrée de l'application
├── handlers/
│   └── handlers.go         # Gestionnaires HTTP
├── models/
│   └── models.go           # Modèles de données
├── templates/              # Templates HTML
│   ├── base.html          # Template de base
│   ├── home.html          # Page d'accueil
│   ├── velos.html         # Catalogue vélos
│   ├── itineraires.html   # Itinéraires
│   ├── actualites.html    # Actualités
│   ├── conseils.html      # Conseils
│   ├── reservation.html   # Formulaire réservation
│   └── confirmation.html  # Confirmation réservation
├── static/                 # Fichiers statiques
│   ├── css/
│   │   └── style.css      # Styles CSS
│   └── images/            # Images (placeholders)
└── data/                   # Stockage des données
    └── reservations.json  # Sauvegarde des réservations
```

### Technologies utilisées
- **Go** : Langage de programmation principal
- **HTML5** : Structure des pages
- **CSS3** : Stylisation responsive
- **JavaScript** : Interactions côté client
- **Unsplash API** : Images de démonstration

### Modèles de données
- **Velo** : Informations sur les vélos
- **Reservation** : Données de réservation
- **Itineraire** : Parcours touristiques
- **Actualite** : Événements et nouvelles
- **Conseil** : Conseils de sécurité

## Installation et lancement

### Prérequis
- Go 1.19 ou supérieur installé
- Navigateur web moderne

### Installation
1. Cloner le projet :
```bash
git clone <url-du-projet>
cd velo-vert
```

2. Installer les dépendances :
```bash
go mod tidy
```

### Lancement
```bash
go run main.go
```

L'application sera accessible sur : http://localhost:8080

## Fonctionnalités développées

### ✅ Initialisation du projet
- Module Go configuré
- Structure de dossiers organisée

### ✅ Modèles de données
- Structures Go pour tous les entités
- Données de démonstration
- Système de persistance JSON

### ✅ Gestionnaires HTTP
- Routes pour toutes les pages
- Traitement des formulaires
- Gestion des erreurs

### ✅ Templates HTML
- Template de base réutilisable
- Pages responsive design
- Intégration des données dynamiques

### ✅ Stylisation CSS
- Design moderne et écologique
- Interface responsive
- Animations et transitions

### ✅ Logique métier
- Calcul des prix de location
- Validation des réservations
- Gestion des disponibilités

### ✅ Stockage des données
- Sauvegarde automatique des réservations
- Chargement au démarrage
- Format JSON pour la persistance

## Utilisation

1. **Navigation** : Utilisez le menu principal pour accéder aux différentes sections
2. **Réservation** : Cliquez sur "Réserver" depuis la page vélos ou d'accueil
3. **Formulaire** : Remplissez le formulaire de réservation avec vos informations
4. **Confirmation** : Recevez une confirmation avec les détails de votre réservation

## Améliorations possibles

- [ ] Ajout d'une base de données (PostgreSQL, MySQL)
- [ ] Système d'authentification utilisateur
- [ ] Interface d'administration
- [ ] Intégration de paiements en ligne
- [ ] API REST pour applications mobiles
- [ ] Géolocalisation des vélos
- [ ] Système de notation et commentaires
- [ ] Notifications par email réelles
- [ ] Optimisation des images
- [ ] Tests unitaires et d'intégration

## Auteur

Développé dans le cadre d'un projet IPI B2 - IA dans le code

## Licence

Ce projet est destiné à des fins éducatives.
