# Script pour préparer le projet pour GitHub
Write-Host "📦 Préparation du projet pour GitHub" -ForegroundColor Cyan

# Créer le repo git s'il n'existe pas
if (!(Test-Path ".git")) {
    Write-Host "🔧 Initialisation du repository Git..." -ForegroundColor Yellow
    git init
    git add .
    git commit -m "Initial commit - Site Le Vélo Vert"
} else {
    Write-Host "✅ Repository Git déjà initialisé" -ForegroundColor Green
}

# Vérifier le statut
Write-Host "📊 Statut du repository :" -ForegroundColor Yellow
git status --short

Write-Host ""
Write-Host "🎯 Prochaines étapes :" -ForegroundColor Green
Write-Host "1. Créez un repository sur GitHub : https://github.com/new" -ForegroundColor White
Write-Host "2. Copiez l'URL du repository" -ForegroundColor White
Write-Host "3. Exécutez : git remote add origin <URL>" -ForegroundColor White
Write-Host "4. Exécutez : git push -u origin main" -ForegroundColor White
Write-Host ""
Write-Host "🚀 Puis déployez sur Railway !" -ForegroundColor Green
