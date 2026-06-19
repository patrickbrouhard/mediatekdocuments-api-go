# syntax=docker/dockerfile:1
# Indique la version de la syntaxe Dockerfile utilisée par Docker BuildKit.

##################################
### Construction du binaire Go ###
# Docker travaille ici dans un environnement isolé basé sur Alpine Linux, avec Go déjà installé.
# On peut le voir comme un mini-système Linux temporaire dans lequel Docker :
# 1. crée un répertoire de travail /app ;
# 2. copie les fichiers de dépendances, télécharge les modules Go, puis copie le code ;
# 3. compile le binaire de l'application.
##################################

# Utilise l'image golang:1.24-alpine, qui contient Go et les outils nécessaires pour compiler l'application.
# Nomme cette étape "builder" afin de pouvoir y faire référence plus tard.
FROM golang:1.26-alpine AS builder

# Définit /app comme répertoire de travail ; Docker le crée s'il n'existe pas.
WORKDIR /app

# Copie les fichiers qui décrivent les dépendances Go dans /app.
COPY go.mod go.sum ./

# Télécharge les dépendances Go nécessaires à la compilation.
RUN go mod download

# Copie le reste du code source du projet dans /app.
COPY . .

# Compile l'API pour Linux, sans dépendance C externe, dans /app/bin/api.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

##################################
### Construction de l'image finale ###
# Cette étape :
# 1. part d'une image Alpine neuve et minimale ;
# 2. copie le binaire depuis l'étape builder ;
# 3. crée un utilisateur non-root ;
# 4. configure le lancement du binaire avec cet utilisateur.
##################################

# Crée l'image finale minimale qui servira uniquement à exécuter l'API.
FROM alpine:3.21

# Définit /app comme dossier de travail de l'application exécutée.
WORKDIR /app

# Crée un utilisateur non-root dédié à l'exécution de l'API.
RUN adduser -D -H -u 10001 appuser

# Copie uniquement le binaire compilé depuis l'étape builder.
COPY --from=builder /app/bin/api /app/api

# Les instructions suivantes et le processus de l'application utiliseront appuser.
USER appuser

# Documente que l'application est prévue pour écouter sur le port 8080.
EXPOSE 8080

# Définit le binaire lancé automatiquement au démarrage du conteneur.
ENTRYPOINT ["/app/api"]

###################
### Définitions ###
# Dockerfile : recette / plan ;
# image : résultat construit, réutilisable ;
# conteneur : instance concrète et exécutée de cette image.
###################
