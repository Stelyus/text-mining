# TextMining - Correction orthographique

Le but est de construire un outil en ligne de commande de correction orthographique rapide et stable en utilisant une distance de Damerau-Levenshtein.
Le projet a été réalisé par **Franck THANG** (EPITA 2019) et **Hadrien BERTHIER** (EPITA 2019) dans le cadre du cours TEXT MINING.

## Introduction

Concernant le langage, nous avons décidé d'utiliser __Go__. C'est un langage dont nous avons jamais eut l'opportunité de tester. C'est un langage de plus en plus populaire, une forme de C et C++ modernisée. Il est reconnu pour avoir une performance comparable au C++.
Go possede un Garbage Collector permettant de nous faciliter la tache concernant les contraintes de memoires.

Le but du projet est de produire deux binaires:
- __TextMiningCompiler__: Il s'agit d'un binaire prenant en argument un fichier contenant des mots et leurs fréquences associés et le nom du fichier binaire qui sera géneré et qui contiendra un Radix Tree sérialisé.
- __TextMiningApp__: Il s'agit d'un binaire qui prend  un fichier binaire contenant le Radix Tree sérialisé ainsi qu'un string à 'corriger' ainsi qu'un nombre représentant la distance de Damerau-Levenshtein.
- 
## Compilation

Il faut tout d'abord installer Go sur son ordinateur Unix. Une fois Go installé:
```
$ sh build.sh
```
Le script va produire deux binaires: __TextMiningCompiler__ et __TextMiningApp__ 

## Documentation

Le code est bien documenté. Go possede un builtin permettant de génerer une documentation sur un server local. Pour se faire, imaginons que je veuille utiliser le port 6060:

```
$ godoc -http=:6060
```
__Attention__: Il faut changer la variable `GOPATH` pour qu'il soit égale à `pwd`.
Ensuite, aller sur `localhost:6060/pkg` pour trouver tous les packages qu'utilisent notre projet. Vous trouverez notamment `app`, `compiler ` et `radix`, ou pouvez directement y accéder en allant sur `localhost:6060/pkg/app` etc.

Si cela ne marche pas, vous pourrez toujours regarder le code par vous même.

## Architecture du projet
```
TEXT_MINING_PROJECT
│   README.md
│   AUTHORS
│   build.sh
└───src
│   │
│   └───app
│   |   │  damerau.go
│   |   │  file.go
|   |   |
|   └───main_app
|   |   |   main.go
|   |
│   └───compiler
|   |   | file.go
|   |   |
|   └───main_compiler
|   |   | main.go
|   |
│   └─── radix
|       | radix.go
───ressources
    │   words.txt
    │   test.txt
    |   subject.txt
```

**build.sh**: Un script bash permettant de creer les deux binaires __TextMiningCompiler__ et __TextMiningApp__

**ressources**: Contient des ressources et les fichiers textes pouvant être passé en paramètre aux binaires.

**app**: Contient les fichiers sources de __TextMiningApp__
**main_app**: Contient le main permettant la compilation de __TextMiningApp__

**compiler**: Contient les fichiers sources de __TextMiningCompiler__
**main_compiler**: Contient le main permettant la compilation de __TextMiningCompiler__

**radix**: Contient le __Radix tree__ utilisé par le compiler et l'app

## Reponses aux questions

###  1.	Decrivez les choix de design de votre programme
### 2.	Listez l’ensemble des tests effectués sur votre programme (en plus des units tests)
### 3.	Avez-vous détecté des cas où la correction par distance ne fonctionnait pas (même avec une distance élevée) ?
### 4.	Quelle est la structure de données que vous avez implémentée dans votre projet, pourquoi ?
### 5.	Proposez un réglage automatique de la distance pour un programme qui prend juste une chaîne de caractères en entrée, donner le processus d’évaluation ainsi que les résultats
### 6.	Comment comptez vous améliorer les performances de votre programme
### 7.	Que manque-t-il à votre correcteur orthographique pour qu’il soit à l’état de l’art ?

