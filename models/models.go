package models

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Velo représente un vélo disponible à la location
type Velo struct {
	ID          int     `json:"id"`
	Nom         string  `json:"nom"`
	Type        string  `json:"type"` // VTT, Ville, Électrique, Enfant
	Description string  `json:"description"`
	PrixJour    float64 `json:"prix_jour"`
	Disponible  bool    `json:"disponible"`
	Image       string  `json:"image"`
}

// Reservation représente une réservation de vélo
type Reservation struct {
	ID        int       `json:"id"`
	VeloID    int       `json:"velo_id"`
	NomVelo   string    `json:"nom_velo"`
	NomClient string    `json:"nom_client"`
	Email     string    `json:"email"`
	Telephone string    `json:"telephone"`
	DateDebut time.Time `json:"date_debut"`
	DateFin   time.Time `json:"date_fin"`
	PrixTotal float64   `json:"prix_total"`
	Status    string    `json:"status"` // En attente, Confirmée, Annulée
}

// Itineraire représente un itinéraire touristique
type Itineraire struct {
	ID          int      `json:"id"`
	Nom         string   `json:"nom"`
	Description string   `json:"description"`
	Distance    float64  `json:"distance"`   // en km
	Duree       int      `json:"duree"`      // en minutes
	Difficulte  string   `json:"difficulte"` // Facile, Moyen, Difficile
	Points      []string `json:"points"`     // Points d'intérêt
	Image       string   `json:"image"`
}

// Actualite représente une actualité sur les événements locaux
type Actualite struct {
	ID      int       `json:"id"`
	Titre   string    `json:"titre"`
	Contenu string    `json:"contenu"`
	Date    time.Time `json:"date"`
	Image   string    `json:"image"`
}

// Conseil représente un conseil de sécurité
type Conseil struct {
	ID        int    `json:"id"`
	Titre     string `json:"titre"`
	Contenu   string `json:"contenu"`
	Categorie string `json:"categorie"` // Sécurité, Entretien, Équipement
}

// Données globales (pour simplifier, on utilise des slices en mémoire)
var (
	Velos        []Velo
	Reservations []Reservation
	Itineraires  []Itineraire
	Actualites   []Actualite
	Conseils     []Conseil
)

// InitData initialise les données de démonstration
func InitData() {
	// Initialiser les vélos
	Velos = []Velo{
		{
			ID:          1,
			Nom:         "VTT Montagne Expert",
			Type:        "VTT",
			Description: "Vélo tout-terrain idéal pour les randonnées en montagne avec suspension complète.",
			PrixJour:    25.0,
			Disponible:  true,
			Image:       "https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=300&fit=crop",
		},
		{
			ID:          2,
			Nom:         "Vélo Ville Confort",
			Type:        "Ville",
			Description: "Vélo urbain confortable avec panier et éclairage intégré pour vos déplacements quotidiens.",
			PrixJour:    15.0,
			Disponible:  true,
			Image:       "https://images.unsplash.com/photo-1485965120184-e220f721d03e?w=400&h=300&fit=crop",
		},
		{
			ID:          3,
			Nom:         "Vélo Électrique Boost",
			Type:        "Électrique",
			Description: "Vélo électrique assisté pour faciliter vos randonnées même sur les terrains difficiles.",
			PrixJour:    35.0,
			Disponible:  true,
			Image:       "https://images.unsplash.com/photo-1571068316344-75bc76f77890?w=400&h=300&fit=crop",
		},
		{
			ID:          4,
			Nom:         "Vélo Enfant Découverte",
			Type:        "Enfant",
			Description: "Vélo adapté aux enfants avec roulettes de sécurité et design coloré.",
			PrixJour:    12.0,
			Disponible:  true,
			Image:       "https://images.unsplash.com/photo-1554224155-6726b3ff858f?w=400&h=300&fit=crop",
		},
	}

	// Initialiser les itinéraires
	Itineraires = []Itineraire{
		{
			ID:          1,
			Nom:         "Balade le long de la rivière",
			Description: "Itinéraire paisible le long de la rivière avec vue sur la nature environnante.",
			Distance:    12.5,
			Duree:       90,
			Difficulte:  "Facile",
			Points:      []string{"Pont historique", "Aire de pique-nique", "Cascade naturelle"},
			Image:       "https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=400&h=300&fit=crop",
		},
		{
			ID:          2,
			Nom:         "Ascension du Col Vert",
			Description: "Défi sportif menant au sommet du Col Vert avec panorama exceptionnel.",
			Distance:    25.0,
			Duree:       180,
			Difficulte:  "Difficile",
			Points:      []string{"Village pittoresque", "Refuge de montagne", "Sommet panoramique"},
			Image:       "https://images.unsplash.com/photo-1506905925346-21bda4d32df4?w=400&h=300&fit=crop",
		},
		{
			ID:          3,
			Nom:         "Circuit urbain découverte",
			Description: "Parcours dans le centre-ville pour découvrir l'architecture et l'histoire locale.",
			Distance:    8.0,
			Duree:       60,
			Difficulte:  "Facile",
			Points:      []string{"Cathédrale", "Musée municipal", "Parc central"},
			Image:       "https://images.unsplash.com/photo-1449824913935-59a10b8d2000?w=400&h=300&fit=crop",
		},
	}

	// Initialiser les actualités
	Actualites = []Actualite{
		{
			ID:      1,
			Titre:   "Festival du Vélo Vert - Édition 2024",
			Contenu: "Rejoignez-nous pour la 5ème édition du Festival du Vélo Vert ! Animations, démonstrations et parcours découverte vous attendent tout le week-end.",
			Date:    time.Now().AddDate(0, 0, 7),
			Image:   "https://images.unsplash.com/photo-1544191694-4514c75a9c2e?w=400&h=300&fit=crop",
		},
		{
			ID:      2,
			Titre:   "Nouveau parcours cyclable inauguré",
			Contenu: "La municipalité a inauguré un nouveau parcours cyclable de 15km reliant le centre-ville aux quartiers périphériques.",
			Date:    time.Now().AddDate(0, 0, -3),
			Image:   "https://images.unsplash.com/photo-1571019613454-1cb2f99b2d8b?w=400&h=300&fit=crop",
		},
	}

	// Initialiser les conseils
	Conseils = []Conseil{
		{
			ID:        1,
			Titre:     "Port du casque obligatoire",
			Contenu:   "Le port du casque est obligatoire pour tous les utilisateurs de vélo. Choisissez un casque à votre taille et ajustez-le correctement.",
			Categorie: "Sécurité",
		},
		{
			ID:        2,
			Titre:     "Vérification avant départ",
			Contenu:   "Avant chaque location, vérifiez l'état des freins, des pneus et de l'éclairage. Signalez tout problème au personnel.",
			Categorie: "Sécurité",
		},
		{
			ID:        3,
			Titre:     "Entretien des freins",
			Contenu:   "Les freins doivent être vérifiés régulièrement. Si vous sentez une anomalie, faites contrôler le vélo immédiatement.",
			Categorie: "Entretien",
		},
		{
			ID:        4,
			Titre:     "Équipement recommandé",
			Contenu:   "Pour une sécurité optimale, équipez-vous de gants, de lunettes de soleil et d'un gilet réfléchissant.",
			Categorie: "Équipement",
		},
	}
}

// SaveData sauvegarde les données dans des fichiers JSON
func SaveData() error {
	// Sauvegarder les réservations
	data, err := json.MarshalIndent(Reservations, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage des réservations: %v", err)
	}
	if err := os.WriteFile("data/reservations.json", data, 0644); err != nil {
		return fmt.Errorf("erreur lors de la sauvegarde des réservations: %v", err)
	}

	return nil
}

// LoadData charge les données depuis les fichiers JSON
func LoadData() error {
	// Charger les réservations
	data, err := os.ReadFile("data/reservations.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Fichier n'existe pas, c'est normal pour la première exécution
			return nil
		}
		return fmt.Errorf("erreur lors de la lecture des réservations: %v", err)
	}

	if err := json.Unmarshal(data, &Reservations); err != nil {
		return fmt.Errorf("erreur lors du décodage des réservations: %v", err)
	}

	// Remplir automatiquement les noms de vélos pour les réservations existantes
	for i := range Reservations {
		for _, velo := range Velos {
			if Reservations[i].VeloID == velo.ID {
				Reservations[i].NomVelo = velo.Nom
				break
			}
		}
	}

	return nil
}
