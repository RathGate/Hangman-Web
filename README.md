# Hangman-Web

Web version of the [Hangman Game](https://github.com/RathGate/Hangman-Go) written in Golang, based on [these instructions](https://github.com/Lyon-Ynov-Campus/YTrack/tree/master/subjects/hangman/hangman-web).
Authors are, once again, [Eva Chibane](https://github.com/evzs) and [Marianne Corbel](https://github.com/RathGate).

**HOW WE ORGANISED OUR WORK:** [HERE (GDrive)](https://drive.google.com/file/d/17SgibsAbYrRuBQR-Zo6qQP6UNfq2dp8L/view?usp=sharing)

## About

**Disclaimer:** This exercise must be submitted for notation before the **19th of December 2022**. As always, it is possible to see one or several commits after this date (notably if we discover issues in the code), but if any, they must not be taken into consideration by the examiner.

The goal of the exercice is to provide a web interface for [this Hangman command-line game](https://github.com/RathGate/Hangman-Go), using all or at least most of the program we wrote beforehand. For clarity and functionality purposes, some functions have been slightly altered and the entire program has been cleared of all the "print" functions and CLI elements.

## Technical Specifications

 - Back-end: Golang 
 - Front-end: HTML, CSS, JS (mostly used to send minor
   ajax requests that would not require to reload the entire page).

**NOTE:** The game has been tested mainly on Google Chrome, Mozilla Firefox, and Edge, all of them in their latest version, and should appear totally responsive. If you encounter any trouble, being with the game or the website itself, tell us and we'll try to fix the issue as soon as possible.

## How to use the program

For now, the website has not been hosted anywhere, though we might find a way to host it in the future. 

In order to see the website and play the game, you must clone the repository on your computer using:

    git clone https://github.com/RathGate/HANGMAN_WEB_CHIBANE_CORBEL.git

Then, launch a terminal at the root of the newly cloned folder:

    go run main.go

The website should be available in any web browser, on `localhost:8080`. If not, there might be a port collision, i.e. the assigned port (:8080) is already being used by something else on the computer. In that case, look for this line of code at the bottom of `./main.go` :

    preferredPort :=  ":8080"

And change the numerical value after the colon with any port number you want to use to play the game.

Enjoy ! ♪
