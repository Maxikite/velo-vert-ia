# 🚀 Guide de Déploiement - Le Vélo Vert

## 📋 Options de déploiement

### 1. 🐳 Déploiement avec Docker (Recommandé)

#### Prérequis
- Docker installé
- Docker Compose installé

#### Déploiement local
```bash
# Construire et démarrer
docker-compose up -d

# Vérifier que ça fonctionne
curl http://localhost:8080
```

#### Commandes utiles
```bash
# Arrêter
docker-compose down

# Redémarrer
docker-compose restart

# Voir les logs
docker-compose logs -f

# Mettre à jour
docker-compose build --no-cache && docker-compose up -d
```

### 2. 🌐 Déploiement sur serveur distant

#### Avec le script automatique
```bash
chmod +x deploy.sh
./deploy.sh
# Choisir option 2 et suivre les instructions
```

#### Manuellement
```bash
# Sur votre serveur
git clone <votre-repo>
cd velo-vert

# Installer Docker
curl -fsSL https://get.docker.com | sh

# Déployer
docker-compose up -d

# Vérifier
curl http://localhost:8080
```

### 3. 📦 Binaire autonome

#### Créer le binaire
```bash
# Pour Linux
GOOS=linux GOARCH=amd64 go build -o velo-vert .

# Pour Windows
GOOS=windows GOARCH=amd64 go build -o velo-vert.exe .

# Pour macOS
GOOS=darwin GOARCH=amd64 go build -o velo-vert .
```

#### Déployer le binaire
```bash
# Copier le binaire et les dossiers
scp velo-vert user@server:/path/to/app/
scp -r static/ templates/ data/ user@server:/path/to/app/

# Sur le serveur
cd /path/to/app
chmod +x velo-vert
./velo-vert
```

## 🌍 Plateformes de déploiement cloud

### 🚀 Railway (Facile)
1. Se connecter sur [Railway.app](https://railway.app)
2. Créer un nouveau projet
3. Connecter votre repo GitHub
4. Railway détecte automatiquement le Dockerfile
5. L'application est déployée !

### 🟣 Render (Gratuit)
1. Se connecter sur [Render.com](https://render.com)
2. Créer un "Web Service"
3. Connecter votre repo GitHub
4. Configuration :
   - **Runtime** : Docker
   - **Port** : 8080

### 🔄 Vercel (Gratuit)
1. Se connecter sur [Vercel.com](https://vercel.com)
2. Créer un nouveau projet
3. Configuration :
   - **Framework** : Other
   - **Build Command** : `go build -o api .`
   - **Output Directory** : `.`
   - **Install Command** : (laisser vide)

### ☁️ Heroku (Payant)
```bash
# Créer l'app
heroku create votre-app-velo-vert

# Déployer
git push heroku main
```

## 🔧 Configuration

### Variables d'environnement
```bash
# Port (défaut: 8080)
PORT=8080

# Base de données (futur)
DATABASE_URL=postgresql://...

# Mode debug
DEBUG=true
```

### Domaines personnalisés
Pour utiliser un domaine personnalisé, configurez votre DNS pour pointer vers l'IP de votre serveur.

## 📊 Monitoring

### Logs
```bash
# Avec Docker
docker-compose logs -f

# Avec systemd (Linux)
journalctl -u velo-vert -f
```

### Health Check
L'application expose un endpoint de santé :
```bash
curl http://votre-domaine/health
```

## 🔒 Sécurité

### HTTPS (Obligatoire pour la production)
- Utilisez un reverse proxy (Nginx, Caddy)
- Obtenez un certificat SSL (Let's Encrypt)
- Forcez HTTPS

### Configuration Nginx (exemple)
```nginx
server {
    listen 80;
    server_name votre-domaine.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### Firewall
```bash
# Ouvrir seulement le port 80/443 et SSH
ufw allow 80
ufw allow 443
ufw allow 22
ufw enable
```

## 🚀 Performance

### Optimisations
- Utilisez un CDN pour les images statiques
- Configurez la compression gzip
- Utilisez un cache (Redis) pour les sessions
- Optimisez les images

### Scaling
Pour gérer plus de trafic :
- Utilisez un load balancer
- Déployez plusieurs instances
- Utilisez une base de données externe

## 🆘 Dépannage

### L'application ne démarre pas
```bash
# Vérifier les logs
docker-compose logs

# Vérifier le port
netstat -tlnp | grep 8080
```

### Erreur 502 Bad Gateway
- Vérifiez que l'application écoute sur le bon port
- Vérifiez la configuration du reverse proxy

### Images qui ne s'affichent pas
- Vérifiez les permissions des dossiers `static/`
- Vérifiez les URLs des images dans `models.go`

---

🎉 **Votre application "Le Vélo Vert" est maintenant déployée !**
