# TextMining - Correction orthographique

Le but est de construire un outil en ligne de commande de correction orthographique rapide et stable en utilisant une distance de Damerau-Levenshtein.
Le projet a été réalisé par **Franck THANG** (EPITA 2019) et **Hadrien BERTHIER** (EPITA 2019) dans le cadre du cours TEXT MINING.

## Introduction

Concernant le langage, nous avons décidé d'utiliser __Go__. C'est un langage dont nous avons jamais eu l'opportunité de tester. C'est un langage de plus en plus populaire, une forme de C et C++ modernisée. Il est reconnu pour avoir une performance comparable au C++.
Go possede un Garbage Collector permettant de nous faciliter la tache concernant les contraintes de memoires.

Le but du projet est de produire deux binaires:
- __TextMiningCompiler__: Il s'agit d'un binaire prenant en argument un fichier contenant des mots et leurs fréquences associés et le nom du fichier binaire qui sera géneré et qui contiendra un Radix Tree sérialisé.
- __TextMiningApp__: Il s'agit d'un binaire qui prend  un fichier binaire contenant le Radix Tree sérialisé ainsi qu'un string à 'corriger' ainsi qu'un nombre représentant la distance de Damerau-Levenshtein.

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

Si cela ne marche pas, vous pourrez toujours regarder le code directement dans les fichiers sources.


## Test suite

Go possede un builtin (aussi) permettant de lancer des tests. Avant de lancer la commande, veuillez etre sur que la variable env `GOROOT`soit égale a `pwd`. La commande permettant de lancer les tests est:

```
$ go test compiler
$ go test app
```

Le script `compare.js` est utilisé pour lancer des tests concernant la correction orthographique.
```
$ node compare.js pathToRefBin pathToRefDict pathToMyBin pathToMyDict
PASSED approx 0 test
FAILED approx 1 test
Not same length: ref: 160, me: 159
Not the same on 62 th element
{ word: 'etst', freq: 5403, distance: 1 }
{ word: 'aest', freq: 4999, distance: 1 }
Print the 5 first elements
MISSING etst

...
```
__Attention__: Un path doit obligatoirement commencé par __./__ s'il se trouve dans le meme repertoire que le script.

## Architecture du projet
```
TEXT_MINING_PROJECT
│   README.md
│   AUTHORS
│   build.sh
│   compare.js
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
|       
| ───test_ressources
    | dict_test.bin

```

**build.sh**: Un script bash permettant de creer les deux binaires __TextMiningCompiler__ et __TextMiningApp__

**compare.js**: Lance une suite de tests sur la correction orthographique.

**ressources**: Contient des ressources et les fichiers textes pouvant être passé en paramètre aux binaires.

**test_ressources**: Contient des ressources utiles pour go test

**app**: Contient les fichiers sources de __TextMiningApp__
**main_app**: Contient le main permettant la compilation de __TextMiningApp__

**compiler**: Contient les fichiers sources de __TextMiningCompiler__
**main_compiler**: Contient le main permettant la compilation de __TextMiningCompiler__

**radix**: Contient le __Radix tree__ utilisé par le compiler et l'app

## Reponses aux questions

###  1.	Decrivez les choix de design de votre programme
Nous avons réalisé le projet en Go, le Go nous impose d'avoir une architecture du projet particulière. Nous devons produire deux executables, et de ce fait nous avons quatres dossiers.

Le dossier app contient des fichiers qui forment le `package app`. Ce package contient toutes les fonctions necessaires pour produire le binaire __TextMiningApp__. Les packages sont des sortes de namespaces.
C'est la meme chose concernant le dossier compiler qui forment le `package compiler`. Leurs "main" respectives sont dans d'autres dossiers car ils sont dans le `package main`.

Le dossier radix contient un fichier qui forme le `package radix`. Ce package est utilisé par les autres fichiers.

### 2.	Listez l’ensemble des tests effectués sur votre programme (en plus des units tests)

### 3.	Avez-vous détecté des cas où la correction par distance ne fonctionnait pas (même avec une distance élevée) ?

### 4.	Quelle est la structure de données que vous avez implémentée dans votre projet, pourquoi ?
Nous avons choisis d'utiliser un Radix tree comme structure de donnée qui est une optimisation du trie où chaque noeud qui est un fils unique est fusionné avec son père. Cette structure permet une optimisation mémoire incroyable et permet
d'effectuer une recherche d'un mots particulier en O(n) n étant la taille du mots. En faisant quelques recherches nous sommes tombés sur cette structure qui semblait être implémentable sans trop de difficulté et nous en avons fait une
première version avant de l'optimiser pour gagner en mémoire.

### 5.	Proposez un réglage automatique de la distance pour un programme qui prend juste une chaîne de caractères en entrée, donner le processus d’évaluation ainsi que les résultats

### 6.	Comment comptez vous améliorer les performances de votre programme

Nous pouvons essayer d'améliorer la sérialisation et la désérialisation qui ne sont pas encore optimal car on a dû baisser notre consommation de RAM. Nous pouvons implementer un Patricia Trie, qui selon nous coute moins cher en mémoire.
Un autre moyen d'ameliorer les performances de notre programmes serait d'utiliser un bloom filter (CF cours :))

### 7.	Que manque-t-il à votre correcteur orthographique pour qu’il soit à l’état de l’art ?


