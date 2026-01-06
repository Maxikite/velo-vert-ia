# 📤 Comment uploader votre code sur GitHub

## 🚀 Méthode ultra-simple (Recommandée)

### Étape 1 : Créer le repository
1. Allez sur [github.com](https://github.com)
2. Cliquez "New repository"
3. Nom : `velo-vert-site`
4. **IMPORTANT** : Rendez-le **PUBLIC**
5. Ne cochez rien d'autre
6. Cliquez "Create repository"

### Étape 2 : Uploader tous les fichiers
1. Sur la page de votre repository, cliquez sur "uploading an existing file"
2. **Sélectionnez TOUS les fichiers** du dossier `velo-vert` :
   - Tous les `.go` files
   - Tous les dossiers (`handlers/`, `models/`, `templates/`, `static/`, `data/`)
   - `Dockerfile`, `docker-compose.yml`
   - `go.mod`, `go.sum`
   - Tous les fichiers `.md`
3. Glissez-déposez ou cliquez "choose your files"
4. Cliquez "Commit changes"

### Étape 3 : Vérifier
Votre repository devrait contenir tous ces fichiers :
```
velo-vert-site/
├── Dockerfile ✅
├── docker-compose.yml ✅
├── go.mod ✅
├── go.sum ✅
├── main.go ✅
├── handlers/
├── models/
├── templates/
├── static/
├── data/
└── *.md files ✅
```

## 🎯 Ensuite : Déployer sur Render

Une fois uploadé sur GitHub :
1. Allez sur [render.com](https://render.com)
2. "New +" → "Web Service"
3. Connectez votre repo `velo-vert-site`
4. Runtime : Docker
5. Port : 8080

**C'est tout ! Votre site sera en ligne en 5 minutes ! 🎉**
