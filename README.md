# Hangman
[![Build Status](https://www.travis-ci.com/TheLe0/Hangman.svg?branch=main)](https://www.travis-ci.com/TheLe0/Hangman)

A simple console implemantation of the hangman game.

To run the app:
```bash
go run main.go
```

To build an executable of the app:
```bash
go build main.go
```

![Print](./images/game.PNG)

If you input more than one character, you are trying to guess the word. Only one character longer input will search in the word, like the tradicional game.

![Print](./images/won.PNG)


![Print](./images/lost.PNG)
