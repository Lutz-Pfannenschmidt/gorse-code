# Go Morsecode - Gorsecode
It's a simple Morsecode translator. It can translate Morsecode to any language that uses the Latin alphabet.
It's Tree based, so it's very fast and (somewhat) efficient (it rebuilts the tree every execution).

## Usage
```bash
go run . -m=".... . .-.. .-.. ---  .-- --- .-. .-.. -.." # HELLO WORLD
```
Where `-m` is the Morsecode string you want to translate.

Words are separated by 2 spaces and letters by 1 space.

![Morse table](table.png)