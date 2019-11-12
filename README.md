# VOTING API IN Golang

## Resume


Après le Brexit l'Europe a pris un coup dans l'aile politiquement et économiquement.
Plusieurs mouvement séparatistes tentent des actions afin de continuer à affaiblir le pouvoir.
​

Un mouvement populaire non violent dont vous faites parti est en train d'émerger.
Entant que développeur vous avez la tâche de créer des outils permettant de voter des propositions de loi permettant aux citoyens de sortir de cette crise.


## Pre-requis

For building the project :
- ``Docker  & Docker-compose``

For test and use the project :
- Tools like ``Postman`` or ``Insomnia`` or ``Curl`` :)

## Utilisation

#### Run the project 
``` docker-compose up ```

#### Update .env
```
# Postgres Live
API_SECRET=voteapi #Used when creating a JWT. It can be anything
DB_HOST=db
DB_DRIVER=postgres
DB_USER=mehdi
DB_PASSWORD=123456789
DB_NAME=vote
DB_PORT=5432 #Default postgres port

# Postgres Test
TestApiSecret=
TestDbHost=
TestDbDriver=
TestDbUser=
TestDbPassword=
TestDbName=
TestDbPort=
```