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
[{arose 30949} {seora 30949} {aries 30623} {raise 30623} {arise 30623} {aesir 30623} {serai 30623} {solea 30089} {aloes 30089} {osela 30089}]
Enter your guess: arose
Enter the result (b: blank, y: yellow, g: green): bbyyg
Guess: arose, Result: bbyyg

Try number 2
There are 16 candidates. Here are the first 10 candidates:
[{sotie 59} {cosie 58} {oside 58} {fosie 57} {josie 57} {socle 56} {sowle 56} {moste 55} {solve 55} {sowte 55}]
Enter your guess: sotie
Enter the result (b: blank, y: yellow, g: green): ggbbg
Guess: sotie, Result: ggbbg

Try number 3
There are 5 candidates. Here are the candidates:
[{socle 19} {solve 19} {sowle 19} {somne 18} {sonde 18}]
Enter your guess: socle
Enter the result (b: blank, y: yellow, g: green): ggbyg
Guess: socle, Result: ggbyg

Try number 4
There are 1 candidates. Here are the candidates:
[{solve 5}]
Enter your guess: solve
Enter the result (b: blank, y: yellow, g: green): ggggg
Guess: solve, Result: ggggg

GJ!
Game over.
```