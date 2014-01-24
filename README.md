# Go Parody Generator

This is an example of using [Markov Chains][0] to generate 'real' sounding text segments.

To run, use `cat word-seed.txt | go run markov.go`.
Make sure `word-seed.txt` is pretty large.
This generator does not do too well with small texts.

This is my first try at Go.
So it may not be as Go-like as it should be in some places.
Feedback is welcomed!

## Customizing Output

You can adjust some settings via command line flags.

```
-words=100: Number of words per paragraph (roughly).
-paragraphs=2: Number of paragraphs.
```

## Example Output

Input file was taken from [Project Gutenberg][1].

```
nervously extended with augmented blazonry, as sagacious as copyists a moment. 
All who was always there;_--first in stationery on in vanity, but his chair with
all subsequent potations were deemed expressive of a solitary in haunting the
favor, and outrages; but highly plume myself at every page or nothing, while 
```

[0]: https://en.wikipedia.org/wiki/Markov_chain "Wikipedia Markov Chain"
[1]: http://gutenberg.org/ebooks/11231 "Bartleby, the Scrivener"
