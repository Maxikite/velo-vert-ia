#!/bin/bash

echo "🚀 Déploiement de Le Vélo Vert"
echo "================================"

# Fonction pour vérifier si Docker est installé
check_docker() {
    if ! command -v docker &> /dev/null; then
        echo "❌ Docker n'est pas installé. Veuillez l'installer : https://docs.docker.com/get-docker/"
        exit 1
    fi
}

# Fonction pour vérifier si docker-compose est installé
check_docker_compose() {
    if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
        echo "❌ Docker Compose n'est pas installé. Veuillez l'installer."
        exit 1
    fi
}

# Déploiement local avec Docker
deploy_local() {
    echo "🏠 Déploiement local avec Docker"
    echo "-------------------------------"

    check_docker
    check_docker_compose

    echo "🔨 Construction de l'image Docker..."
    docker-compose build

    echo "🚀 Démarrage du conteneur..."
    docker-compose up -d

    echo "✅ Application déployée localement !"
    echo "🌐 Accessible sur : http://localhost:8080"
    echo ""
    echo "📋 Commandes utiles :"
    echo "  - Arrêter : docker-compose down"
    echo "  - Logs : docker-compose logs -f"
    echo "  - Redémarrer : docker-compose restart"
}

# Déploiement sur un serveur distant
deploy_remote() {
    echo "🌐 Déploiement sur serveur distant"
    echo "----------------------------------"

    read -p "Adresse du serveur (ex: user@192.168.1.100) : " server
    read -p "Port SSH (défaut: 22) : " port
    port=${port:-22}

    echo "🔄 Déploiement vers $server:$port"

    # Créer le dossier data s'il n'existe pas
    mkdir -p data

    # Transférer les fichiers
    echo "📤 Transfert des fichiers..."
    scp -P $port -r . $server:~/velo-vert/

    # Se connecter et déployer
    ssh -p $port $server << EOF
        cd ~/velo-vert

        # Vérifier si Docker est installé
        if ! command -v docker &> /dev/null; then
            echo "❌ Docker n'est pas installé sur le serveur distant."
            echo "Installez Docker avec : curl -fsSL https://get.docker.com | sh"
            exit 1
        fi

        # Construire et démarrer
        echo "🔨 Construction de l'image..."
        docker-compose build

        echo "🚀 Démarrage du service..."
        docker-compose up -d

        echo "✅ Application déployée !"
        echo "🌐 Vérifiez l'IP publique du serveur sur le port 8080"
EOF
}

# Menu principal
echo "Choisissez votre méthode de déploiement :"
echo "1) 🚀 Déploiement local avec Docker"
echo "2) 🌐 Déploiement sur serveur distant"
echo "3) 📦 Création d'un binaire autonome"
echo ""

read -p "Votre choix (1-3) : " choice

case $choice in
    1)
        deploy_local
        ;;
    2)
        deploy_remote
        ;;
    3)
        echo "📦 Création du binaire..."
        echo "Binaire créé : velo-vert.exe"
        echo "Lancez-le avec : ./velo-vert.exe"
        echo ""
        echo "⚠️  Note : Le binaire doit être lancé dans le même dossier"
        echo "   que les dossiers 'static', 'templates' et 'data'"
        ;;
    *)
        echo "❌ Choix invalide"
        exit 1
        ;;
esac

echo ""
echo "🎉 Déploiement terminé !"
