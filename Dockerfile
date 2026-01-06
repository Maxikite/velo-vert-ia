# Utiliser une image Go officielle pour builder
FROM golang:1.21-alpine AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers Go
COPY go.mod ./
RUN go mod download

# Copier le code source
COPY . .

# Builder l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o velo-vert .

# Image finale légère
FROM alpine:latest

# Installer ca-certificates pour les requêtes HTTPS
RUN apk --no-cache add ca-certificates

# Créer un utilisateur non-root
RUN adduser -D -s /bin/sh appuser

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable depuis l'étape de build
COPY --from=builder /app/velo-vert .

# Créer les dossiers nécessaires
RUN mkdir -p data static/css static/images templates

# Copier les fichiers statiques et templates
COPY static/ ./static/
COPY templates/ ./templates/

# Changer les permissions
RUN chown -R appuser:appuser /app
USER appuser

# Exposer le port
EXPOSE 8080

# Commande pour démarrer l'application
CMD ["./velo-vert"]
