# Rapport de Projet

## Qu'est-ce que nous avons fait

Nous avons réalisé les tâches suivantes pour ce projet :

1. **Création des Dockerfiles :**
   - **docs** : Nous avons créé un Dockerfile pour le projet `docs` en utilisant une build multi-étapes pour optimiser la taille de l'image. La première étape construit l'application avec Node.js, et la deuxième étape utilise Nginx pour servir les fichiers statiques.
   - **vote-api** : Nous avons créé un Dockerfile pour le projet `vote-api` en utilisant une build multi-étapes pour optimiser la taille de l'image. La première étape compile l'application Go, et la deuxième étape utilise une image Ubuntu pour exécuter l'application.
   - **web-client** : Nous avons créé un Dockerfile pour le projet `web-client` en utilisant une build multi-étapes pour optimiser la taille de l'image. La première étape construit l'application avec Node.js, et la deuxième étape utilise une image Node.js Alpine pour exécuter l'application.

2. **Configuration de Docker Compose :**
   - Nous avons créé un fichier `docker-compose.yml` pour exécuter les applications `docs`, `web-client` et `vote-api` avec une base de données PostgreSQL. Ce fichier configure et orchestre les services, définit les dépendances entre eux et expose les ports nécessaires pour accéder aux applications.

3. **Tests en local :**
   - Nous avons testé les applications en local en utilisant Docker Compose avec les liens suivants :
     - `docs` : [http://localhost:8081](http://localhost:8081)
     - `web-client` : [http://localhost:3000](http://localhost:3000)
     - `vote-api` : [http://localhost:8080/votes](http://localhost:8080/votes)

4. **Déploiement des applications :**
   - **docs** : Nous avons déployé l'application `docs` sur Netlify.
   - **vote-api** : Nous avons déployé l'application `vote-api` sur Render en utilisant une base de données PostgreSQL hébergée sur Neon.
   - **web-client** : Nous avons déployé l'application `web-client` sur Render.

5. **Mise en place du pipeline CI/CD :**
   - Nous avons configuré un pipeline CI/CD avec GitHub Actions pour automatiser le build, les tests et le déploiement des applications.

## Comment se déploie-t-il

### Déploiement de `docs`

L'application `docs` est déployée sur Netlify. Voici les étapes pour le déploiement :
1. Connectez-vous à Netlify et créez un nouveau site à partir de Git.
2. Choisissez votre fournisseur Git et autorisez Netlify à accéder à votre dépôt.
3. Sélectionnez votre dépôt et configurez les paramètres de build :
   - **Branch to deploy** : `main`
   - **Base directory** : `docs`
   - **Build command** : `npm run build`
   - **Publish directory** : `build`
4. Cliquez sur "Deploy site". Netlify va commencer le processus de build et déployer votre site.

### Déploiement de `vote-api`

L'application `vote-api` est déployée sur Render. Voici les étapes pour le déploiement :
1. **Inscription et connexion à Neon :**
   - Allez sur Neon et inscrivez-vous ou connectez-vous à votre compte.
   - Créez un nouveau projet et une nouvelle base de données.
   - Obtenez l'URL de connexion PostgreSQL.

2. **Créer un compte Render :**
   - Si vous n'avez pas encore de compte Render, inscrivez-vous sur Render.

3. **Créer un nouveau service :**
   - Connectez-vous à votre compte Render.
   - Cliquez sur "New" et sélectionnez "Web Service".
   - Connectez votre dépôt GitHub à Render.
   - Sélectionnez le dépôt contenant votre projet `vote-api` et la branche `main`.

4. **Configurer le service :**
   - **Name** : `vote-api`
   - **Environment** : `Docker`  
   - **Root Directory** : `vote-api`  
   - **Build Command** : Laissez vide si vous utilisez un Dockerfile.
   - **Start Command** : Laissez vide si vous utilisez un Dockerfile.
   - Ajoutez la variable d'environnement `PG_URL` avec la valeur de l'URL de connexion PostgreSQL.

5. **Configurer le Dockerfile :**
   - Render détectera automatiquement votre Dockerfile si vous en avez un à la racine de votre projet ou dans le répertoire `vote-api`.

### Déploiement de `web-client`
Avant de déployer `web-client`, nous devons modifier les fichiers `viteApi.test.ts` et `voteApi.ts` pour utiliser l'URL de l'API déployée sur Render.

#### Modification des fichiers

1. **Modifier `viteApi.test.ts` :**
   - Remplacez l'URL de l'API locale par l'URL de l'API déployée sur Render :
     ```typescript
     // const MOCK_API_BASE_URL = "http://localhost:8080";
     const MOCK_API_BASE_URL = "https://tp-projet-finale.onrender.com";
     ```

2. **Modifier `voteApi.ts` :**
   - Remplacez l'URL de l'API locale par l'URL de l'API déployée sur Render :
     ```typescript
     // const voteApiBaseUrl = process.env.VOTE_API_BASE_URL ?? "http://localhost:8080";
     const voteApiBaseUrl = process.env.VOTE_API_BASE_URL ?? "https://tp-projet-finale.onrender.com";
     ```
### Déploiement de [web-client] (suite)

L'application `web-client` est déployée sur Render. Voici les étapes pour le déploiement :
1. **Créer un compte Render :**
   - Si vous n'avez pas encore de compte Render, inscrivez-vous sur Render.

2. **Créer un nouveau service :**
   - Connectez-vous à votre compte Render.
   - Cliquez sur "New" et sélectionnez "Web Service".
   - Connectez votre dépôt GitHub à Render.
   - Sélectionnez le dépôt contenant votre projet `web-client` et la branche `main`.

3. **Configurer le service :**
   - **Name** : `web-client`
   - **Environment** : `Docker`
   - **Root Directory** : `web-client`  
   - **Build Command** : Laissez vide si vous utilisez un Dockerfile.
   - **Start Command** : Laissez vide si vous utilisez un Dockerfile.
   - Ajoutez les variables d'environnement suivantes :
     - `NODE_ENV` : `production`
     - `VOTE_API_BASE_URL` : `https://tp-projet-finale.onrender.com`
     - `PORT` : `3000`

4. **Configurer le Dockerfile :**
   - Render détectera automatiquement votre Dockerfile si vous en avez un à la racine de votre projet ou dans le répertoire `web-client`.

## Mise en place du pipeline CI/CD

Nous avons configuré un pipeline CI/CD avec GitHub Actions pour automatiser le build, les tests et le déploiement des applications. Voici comment nous avons configuré le pipeline :

### Configuration du pipeline CI/CD

1. **Créer un fichier de workflow GitHub Actions :**
   - Créez un fichier nommé `deploy.yml` dans le répertoire `.github/workflows` de votre dépôt.

2. **Définir les déclencheurs :**
   - Le pipeline se déclenche sur chaque push vers la branche `main`.

3. **Définir les jobs :**
   - Le pipeline est divisé en trois jobs : `build-and-deploy-vote-api`, `build-and-deploy-web-client`, et `build-and-deploy-docs`.

4. **Configurer les jobs :**
   - Chaque job utilise une machine virtuelle Ubuntu pour exécuter les étapes de build et de déploiement.

### Explication des jobs

#### Job `build-and-deploy-vote-api`

- **Checkout code** : Récupère le code source du dépôt.
- **Set up Go** : Configure l'environnement Go.
- **Build Docker image for vote-api** : Construit l'image Docker pour `vote-api`.
- **Push Docker image to Docker Hub** : Pousse l'image Docker vers Docker Hub.
- **Deploy vote-api to Render** : Déploie `vote-api` sur Render en utilisant l'API de Render.

#### Job `build-and-deploy-web-client`

- **Modifier les fichiers `viteApi.test.ts` et `voteApi.ts`** : Modifie les fichiers pour utiliser l'URL de l'API déployée sur Render.
- **Checkout code** : Récupère le code source du dépôt.
- **Set up Node.js** : Configure l'environnement Node.js.
- **Build Docker image for web-client** : Construit l'image Docker pour `web-client`.
- **Push Docker image to Docker Hub** : Pousse l'image Docker vers Docker Hub.
- **Deploy web-client to Render** : Déploie `web-client` sur Render en utilisant l'API de Render.

#### Job `build-and-deploy-docs`

- **Checkout code** : Récupère le code source du dépôt.
- **Set up Node.js** : Configure l'environnement Node.js.
- **Install dependencies for docs** : Installe les dépendances pour `docs`.
- **Build the docs site** : Construit le site `docs`.
- **Build Docker image for docs** : Construit l'image Docker pour `docs`.
- **Push Docker image to Docker Hub** : Pousse l'image Docker vers Docker Hub.
- **Deploy docs to Netlify** : Déploie `docs` sur Netlify en utilisant une action GitHub dédiée.

### Obtention des secrets nécessaires

Pour configurer les secrets nécessaires au pipeline CI/CD, suivez les étapes ci-dessous :

#### Docker Hub

1. **DOCKER_USERNAME** et **DOCKER_PASSWORD** :
   - Créez un compte sur [Docker Hub](https://hub.docker.com/).
   - Allez dans les paramètres de votre compte et générez un mot de passe d'application si nécessaire.
   - Ajoutez `DOCKER_USERNAME` et `DOCKER_PASSWORD` dans les secrets de votre dépôt GitHub :
     - Allez dans les paramètres de votre dépôt GitHub.
     - Cliquez sur "Secrets and variables" > "Actions".
     - Cliquez sur "New repository secret" et ajoutez `DOCKER_USERNAME` et `DOCKER_PASSWORD`.

#### Render

1. **RENDER_API_KEY** :
   - Créez un compte sur [Render](https://render.com/).
   - Allez dans les paramètres de votre compte et générez une clé API.
   - Ajoutez `RENDER_API_KEY` dans les secrets de votre dépôt GitHub :
     - Allez dans les paramètres de votre dépôt GitHub.
     - Cliquez sur "Secrets and variables" > "Actions".
     - Cliquez sur "New repository secret" et ajoutez `RENDER_API_KEY`.

2. **Service ID (srv-ctdjf3ij1k6c73dqppb0)** :
   - Allez dans le tableau de bord de Render.
   - Sélectionnez votre service `vote-api`.
   - L'ID du service se trouve dans l'URL de la page du service (par exemple, `srv-ctdjf3ij1k6c73dqppb0`).

#### Netlify

1. **NETLIFY_AUTH_TOKEN** et **NETLIFY_SITE_ID** :
   - Créez un compte sur [Netlify](https://www.netlify.com/).
   - Allez dans les paramètres de votre compte et générez un token d'accès personnel.
   - Allez dans les paramètres de votre site et copiez l'ID du site.
   - Ajoutez `NETLIFY_AUTH_TOKEN` et `NETLIFY_SITE_ID` dans les secrets de votre dépôt GitHub :
     - Allez dans les paramètres de votre dépôt GitHub.
     - Cliquez sur "Secrets and variables" > "Actions".
     - Cliquez sur "New repository secret" et ajoutez `NETLIFY_AUTH_TOKEN` et `NETLIFY_SITE_ID`.

## Liens vers les images et applications déployées

- `web-client` : [https://tp-projet-finale-1.onrender.com](https://td-projet-final.onrender.com)
- `vote-api` : [https://tp-projet-finale.onrender.com/votes](https://tp-projet-finale.onrender.com)
- `docs` : [https://projetfinale.netlify.app](https://projetfinale.netlify.app/)

## Liens vers les images Docker

- `docs` : [diawara1/docs:latest](https://hub.docker.com/r/diawara1/docs/tags)
- `vote-api` : [diawara1/vote-api:latest](https://hub.docker.com/r/diawara1/vote-api/tags)
- `web-client` : [diawara1/web-client:latest](https://hub.docker.com/r/diawara1/web-client/tags)

## Comment un nouvel utilisateur devrait-il contribuer au projet

Pour contribuer à ce projet, suivez les étapes ci-dessous :

1. **Forker le dépôt :**
   - Allez sur la page du dépôt et cliquez sur le bouton "Fork" pour créer une copie du dépôt dans votre compte GitHub.

2. **Cloner le dépôt forké :**
   - Clonez le dépôt forké sur votre machine locale 

3. **Créer une branche pour votre fonctionnalité :**
   - Créez une nouvelle branche à partir de `develop` pour travailler sur votre fonctionnalité :
     ```sh
     git checkout develop
     git checkout -b feature/ma-fonctionnalite
     ```

4. **Faire des modifications et committer :**
   - Faites les modifications nécessaires et commitez-les en utilisant la convention Gitmoji :
     ```sh
     git add .
     git commit -m ":sparkles: Ajouter une nouvelle fonctionnalité"
     ```

5. **Pousser les modifications et créer une Pull Request :**
   - Poussez les modifications vers votre dépôt forké :
     ```sh
     git push origin feature/ma-fonctionnalite
     ```
   - Allez sur la page de votre dépôt forké sur GitHub et créez une Pull Request vers le dépôt original.

### Utilisation des branches `release/` et `hotfix/`

#### Branches `release/` pour préparer une nouvelle version

Les branches `release/` sont utilisées pour finaliser une nouvelle version. Elles permettent d'ajouter les derniers correctifs et tests avant de fusionner dans `main`.

1. **Créer une branche `release/` depuis `develop` :**
   - Créez une branche `release/vX.Y.Z` :
     ```sh
     git checkout develop
     git checkout -b release/vX.Y.Z
     ```

2. **Apporter des modifications nécessaires :**
   - Résolvez des bugs, mettez à jour les fichiers de version, ou préparez la documentation.
   - Commitez les modifications :
     ```sh
     git add .
     git commit -m "fix: corrections pour la version vX.Y.Z"
     ```

3. **Fusionner dans `main` et marquer la version :**
   - Fusionnez la branche `release/vX.Y.Z` dans `main` et marquez la version avec un tag :
     ```sh
     git checkout main
     git merge release/vX.Y.Z
     git tag -a vX.Y.Z -m "Release version vX.Y.Z"
     git push origin main --tags
     ```

4. **Fusionner dans `develop` pour synchroniser les modifications :**
   - Fusionnez la branche `release/vX.Y.Z` dans `develop` :
     ```sh
     git checkout develop
     git merge release/vX.Y.Z
     git push origin develop
     ```

#### Branches `hotfix/` pour des correctifs urgents

Les branches `hotfix/` sont utilisées pour résoudre rapidement des problèmes critiques dans la branche `main`, sans attendre une nouvelle version complète.

1. **Créer une branche `hotfix/` depuis `main` :**
   - Créez une branche `hotfix/vX.Y.Z` :
     ```sh
     git checkout main
     git checkout -b hotfix/vX.Y.Z
     ```

2. **Corriger les erreurs urgentes :**
   - Apportez les correctifs nécessaires et commitez-les :
     ```sh
     git add .
     git commit -m "hotfix: correction critique dans la version vX.Y.Z"
     ```

3. **Fusionner dans `main` et marquer la version corrigée :**
   - Fusionnez la branche `hotfix/vX.Y.Z` dans `main` et marquez la version corrigée avec un tag :
     ```sh
     git checkout main
     git merge hotfix/vX.Y.Z
     git tag -a vX.Y.Z -m "Hotfix version vX.Y.Z"
     git push origin main --tags
     ```

4. **Fusionner dans `develop` pour inclure les corrections dans la prochaine version :**
   - Fusionnez la branche `hotfix/vX.Y.Z` dans `develop` :
     ```sh
     git checkout develop
     git merge hotfix/vX.Y.Z
     git push origin develop
     ```