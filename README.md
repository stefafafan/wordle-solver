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
-----Try number 1-----
There are 14849 candidates. Here are the first 10 candidates:
[aahed aalii aargh aaron abaca abaci aback abada abaff abaft]

Here are the first 10 recommendations to help narrow down the word:
[{arose 30949} {seora 30949} {serai 30623} {arise 30623} {raise 30623} {aesir 30623} {aries 30623} {aloes 30089} {osela 30089} {solea 30089}]

Enter your guess: arose
Enter the result (b: blank, y: yellow, g: green): bbyyg
Guess: arose, Result: bbyyg

-----Try number 2-----
There are 16 candidates. Here are the first 10 candidates:
[cosie essee fosie goloe hoboe issue josie moste oside socle]

Here are the first 10 recommendations to help narrow down the word:
[{toise 59} {sotie 59} {idose 58} {cosie 58} {moise 58} {diose 58} {noise 58} {eidos 58} {eosin 58} {oside 58}]

Enter your guess: toise
Enter the result (b: blank, y: yellow, g: green): bgbyg
Guess: toise, Result: bgbyg

-----Try number 3-----
There are 7 candidates. Here are the candidates:
[goloe hoboe socle solve somne sonde sowle]

Here are the first 10 recommendations to help narrow down the word:
[{solen 27} {slone 27} {lenos 27} {enols 27} {noels 27} {socle 26} {lowse 26} {hosel 26} {sheol 26} {solve 26}]

Enter your guess: solen
Enter the result (b: blank, y: yellow, g: green): gggyb
Guess: solen, Result: gggyb

-----Try number 4-----
There is 1 candidate. This is probably the answer: solve

Enter your guess: solve
Enter the result (b: blank, y: yellow, g: green): ggggg
Guess: solve, Result: ggggg

GJ!
Game over.
```