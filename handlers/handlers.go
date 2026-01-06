package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"velo-vert/models"
)

// TemplateData représente les données passées aux templates
type TemplateData struct {
	Title        string
	Velos        []models.Velo
	Itineraires  []models.Itineraire
	Actualites   []models.Actualite
	Conseils     []models.Conseil
	Reservations []models.Reservation
	Velo         *models.Velo
	Itineraire   *models.Itineraire
	Actualite    *models.Actualite
	Reservation  *models.Reservation
	Message      string
	Error        string
}

// HomeHandler gère la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := TemplateData{
		Title:       "Le Vélo Vert - Location de vélos écologiques",
		Velos:       models.Velos[:3], // Afficher seulement 3 vélos sur la page d'accueil
		Itineraires: models.Itineraires[:2], // Afficher 2 itinéraires
		Actualites:  models.Actualites,
	}

	renderTemplate(w, "home", data)
}

// VelosHandler gère la page des vélos
func VelosHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Nos vélos - Le Vélo Vert",
		Velos: models.Velos,
	}

	renderTemplate(w, "velos", data)
}

// ItinerairesHandler gère la page des itinéraires
func ItinerairesHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title:       "Itinéraires touristiques - Le Vélo Vert",
		Itineraires: models.Itineraires,
	}

	renderTemplate(w, "itineraires", data)
}

// ActualitesHandler gère la page des actualités
func ActualitesHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title:      "Actualités - Le Vélo Vert",
		Actualites: models.Actualites,
	}

	renderTemplate(w, "actualites", data)
}

// ConseilsHandler gère la page des conseils de sécurité
func ConseilsHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title:    "Conseils de sécurité - Le Vélo Vert",
		Conseils: models.Conseils,
	}

	renderTemplate(w, "conseils", data)
}

// ReservationsHandler gère la page des réservations
func ReservationsHandler(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title:       "Toutes les réservations - Le Vélo Vert",
		Reservations: models.Reservations,
	}

	renderTemplate(w, "reservations", data)
}

// ReservationHandler gère les réservations
func ReservationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Afficher le formulaire de réservation
		veloIDStr := r.URL.Query().Get("velo_id")
		if veloIDStr == "" {
			// Rediriger vers la liste des vélos si pas d'ID
			http.Redirect(w, r, "/velos", http.StatusSeeOther)
			return
		}

		veloID, err := strconv.Atoi(veloIDStr)
		if err != nil {
			http.Error(w, "ID de vélo invalide", http.StatusBadRequest)
			return
		}

		// Trouver le vélo
		var selectedVelo *models.Velo
		for i := range models.Velos {
			if models.Velos[i].ID == veloID {
				selectedVelo = &models.Velos[i]
				break
			}
		}

		if selectedVelo == nil {
			http.Error(w, "Vélo non trouvé", http.StatusNotFound)
			return
		}

		data := TemplateData{
			Title: "Réserver un vélo - Le Vélo Vert",
			Velo:  selectedVelo,
		}

		renderTemplate(w, "reservation", data)
		return
	}

	if r.Method == "POST" {
		// Traiter la soumission du formulaire
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Erreur lors du traitement du formulaire", http.StatusBadRequest)
			return
		}

		veloID, err := strconv.Atoi(r.FormValue("velo_id"))
		if err != nil {
			http.Error(w, "ID de vélo invalide", http.StatusBadRequest)
			return
		}

		dateDebut, err := time.Parse("2006-01-02", r.FormValue("date_debut"))
		if err != nil {
			data := TemplateData{
				Title: "Réserver un vélo - Le Vélo Vert",
				Error: "Format de date invalide",
			}
			renderTemplate(w, "reservation", data)
			return
		}

		dateFin, err := time.Parse("2006-01-02", r.FormValue("date_fin"))
		if err != nil {
			data := TemplateData{
				Title: "Réserver un vélo - Le Vélo Vert",
				Error: "Format de date invalide",
			}
			renderTemplate(w, "reservation", data)
			return
		}

		// Calculer le nombre de jours
		duration := dateFin.Sub(dateDebut)
		jours := int(duration.Hours() / 24)
		if jours <= 0 {
			data := TemplateData{
				Title: "Réserver un vélo - Le Vélo Vert",
				Error: "La date de fin doit être après la date de début",
			}
			renderTemplate(w, "reservation", data)
			return
		}

		// Trouver le vélo pour calculer le prix
		var selectedVelo *models.Velo
		for i := range models.Velos {
			if models.Velos[i].ID == veloID {
				selectedVelo = &models.Velos[i]
				break
			}
		}

		if selectedVelo == nil {
			http.Error(w, "Vélo non trouvé", http.StatusNotFound)
			return
		}

		prixTotal := float64(jours) * selectedVelo.PrixJour

		// Créer la réservation
		reservation := models.Reservation{
			ID:        len(models.Reservations) + 1,
			VeloID:    veloID,
			NomVelo:   selectedVelo.Nom,
			NomClient: r.FormValue("nom"),
			Email:     r.FormValue("email"),
			Telephone: r.FormValue("telephone"),
			DateDebut: dateDebut,
			DateFin:   dateFin,
			PrixTotal: prixTotal,
			Status:    "En attente",
		}

		// Ajouter à la liste
		models.Reservations = append(models.Reservations, reservation)

		// Sauvegarder les données
		if err := models.SaveData(); err != nil {
			fmt.Printf("Erreur lors de la sauvegarde: %v\n", err)
		}

		// Rediriger vers une page de confirmation
		data := TemplateData{
			Title:       "Réservation confirmée - Le Vélo Vert",
			Reservation: &reservation,
			Message:     "Votre réservation a été enregistrée avec succès !",
		}

		renderTemplate(w, "confirmation", data)
		return
	}

	http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
}

// renderTemplate rend un template avec les données fournies
func renderTemplate(w http.ResponseWriter, templateName string, data TemplateData) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+templateName+".html",
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur de template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Erreur d'exécution du template: %v", err), http.StatusInternalServerError)
		return
	}
}
