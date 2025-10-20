# Refactoring – Bonnes Pratiques de Conception
Groupe : Kelyan DANIS, Jeremy GAUDIN, Cassian JOLY

## Instructions

Pour utiliser ce projet :

1. Téléchargez-le depuis GitHub au format `.zip`.
2. Décompressez l’archive.
3. Ouvrez un terminal ou un **invite de commandes**.
4. Rendez-vous à la racine du dossier du projet, puis exécutez la commande :
go run main.go

text
(Assurez-vous que **Go (Golang)** est installé sur votre machine.)

## Objectif du projet

Réaliser le TP de refactoring en appliquant les bonnes pratiques de conception logicielle.

## Technologies utilisées

- **Golang** pour le serveur et la logique métier.  
- **HTML**, **CSS** et **JavaScript** pour le front-end.  
- **JSON** pour la gestion et le stockage des données.

## Organisation du projet

Le projet suit une architecture modulaire de type **MVC** :
1 dossier **data/** : contient les données (Modèle).  
2 dossier **static/** et **templates/** : contiennent les vues (Vue).  
et 1 **service/** pour les gouverners tous : gère la logique applicative et fait le lien entre les modèles et les vues (Contrôleur).