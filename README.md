# TextMining - Correction orthographique

Le but est de construire un outil en ligne de commande de correction orthographique rapide et stable en utilisant une distance de Damerau-Levenshtein.
Le projet a été réalisé par **Franck THANG** (EPTIA 2019) et **Hadrien BERTHIER** (EPTIA 2019) dans le cadre du cours TEXT MINING.

## Introduction

Concernant le langage, nous avons décidé d'utiliser __Go__. C'est un langage dont nous avons jamais eut l'opportunité de tester. C'est un langage de plus en plus populaire, une forme de C et C++ modernisée.
Go possede un Garbage Collector permettant de nous faciliter la tache concernant les contraintes de memoires.

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
│   |   │  main.go
|   |
│   └───compiler
|   |   | file.go
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

**compiler**: Contient les fichiers sources de __TextMiningCompiler__

**app**: Contient les fichiers sources de __TextMiningApp__

**radix**: Contient le __Radix tree__ utilisé par le compiler et l'app


