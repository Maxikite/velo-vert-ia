package main

import (
	"fmt"
	"log"
	"net/http"
	"velo-vert/handlers"
	"velo-vert/models"
)

func main() {
	// Initialiser les données
	models.InitData()

	// Charger les données sauvegardées
	if err := models.LoadData(); err != nil {
		fmt.Printf("Erreur lors du chargement des données: %v\n", err)
	}

	// Configurer les routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/velos", handlers.VelosHandler)
	http.HandleFunc("/itineraires", handlers.ItinerairesHandler)
	http.HandleFunc("/reservation", handlers.ReservationHandler)
	http.HandleFunc("/reservations", handlers.ReservationsHandler)
	http.HandleFunc("/actualites", handlers.ActualitesHandler)
	http.HandleFunc("/conseils", handlers.ConseilsHandler)

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","service":"velo-vert"}`))
	})

	// Servir les fichiers statiques (CSS, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	fmt.Println("🚴 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
