# EstagTest-BZRGROUP

This repository contains three small Go exercises, each organized in its own folder:

- `ponto1` — reads a JSON file and sums all numeric values found in it.
- `ponto2` — compares two text files after converting one of them from ISO-8859-1 to UTF-8.
- `ponto3456` — calculates the number of days between two dates for a set of years and prints the results.

## Project structure

### `ponto1`
Files:
- `main.go` — program entry point.
- `test1.json` — input data used by the program.

What the code does:
- Reads `test1.json`.
- Parses the JSON as a list of values.
- Adds together values that are numeric or numeric strings.
- Prints the final sum.

### `ponto2`
Files:
- `main.go` — program entry point.
- `test2_text1.txt` — first text file.
- `test2_text2.txt` — second text file encoded in ISO-8859-1.

What the code does:
- Reads both text files.
- Converts the second file from ISO-8859-1 to UTF-8.
- Compares the two strings by visible text, not by raw bytes.
- Prints whether they are equal.

### `ponto3456`
Files:
- `main.go` — program entry point.

What the code does:
- Defines a list of years.
- Creates two dates for each year: February 15 and October 15.
- Calculates the number of days between those dates.
- Prints the result for each year.

## Notes

- The repository uses Go.
- Each folder is an independent exercise with its own `main.go`.
- The files in each folder are meant to be run from that folder so the relative file paths work correctly.

## How to run

For each folder, open the folder and run:

```bash
go run .
```

If needed, install dependencies first:

```bash
go mod tidy
```
