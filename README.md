# wordle-solver

Recommends guesses for https://www.powerlanguage.co.uk/wordle/

## How to use this

Prepare a dictionary file that has five-letter words in it.

Example:
```sh
aahed
aalii
aargh
aaron
abaca
...
```

go run
```sh
go run main.go --dictionary five-letter-words.txt
```

## Example
Let's say the answer is `solve` .
For each try, the number of candidates and the list of recommendations will be output.
Enter your guess and result each time.

```sh
$ go run main.go --dictionary five-letter-words.txt

Try number 1
There are 14848 candidates. Here are the first 10 candidates:
[{seora 30949} {arose 30949} {aries 30623} {arise 30623} {serai 30623} {aesir 30623} {raise 30623} {aloes 30089} {alose 30089} {osela 30089}]
Enter your guess: seora
Enter the result (b: blank, y: yellow, g: green): gyybb
Guess: seora, Result: gyybb

Try number 2
There are 565 candidates. Here are the first 10 candidates:
[{slite 1652} {stile 1652} {suite 1625} {stine 1610} {stein 1610} {snite 1610} {silen 1605} {sline 1605} {stipe 1589} {spite 1589}]
Enter your guess: slite
Enter the result (b: blank, y: yellow, g: green): gybbg
Guess: slite, Result: gybbg

Try number 3
There are 20 candidates. Here are the first 10 candidates:
[{souse 68} {shune 67} {spume 67} {sonde 67} {spuke 67} {somne 67} {socle 67} {scuse 66} {sowle 65} {skuse 64}]
Enter your guess: souse
Enter the result (b: blank, y: yellow, g: green): ggbyg
Guess: souse, Result: ggbyg

Try number 4
There are 5 candidates. Here are the candidates:
[{sowle 19} {socle 19} {solve 19} {somne 18} {sonde 18}]
Enter your guess: sowle
Enter the result (b: blank, y: yellow, g: green): ggbyg
Guess: sowle, Result: ggbyg

Try number 5
There are 3 candidates. Here are the candidates:
[{somne 12} {sonde 12} {solve 11}]
Enter your guess: somne
Enter the result (b: blank, y: yellow, g: green): ggbbg
Guess: somne, Result: ggbbg

Try number 6
There are 1 candidates. Here are the candidates:
[{solve 5}]
Enter your guess: solve
Enter the result (b: blank, y: yellow, g: green): ggggg
Guess: solve, Result: ggggg

GJ!
Game over.
```