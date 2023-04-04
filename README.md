# rubiks-cube
My thesis project about Rubik cubes and algorithms to solve it

## Frontend

The frontend consists of an editor and a simulator part. The editor is an integrated Monaco editor with two tabs.
The simulator is based on [this repository](https://github.com/taylorjg/rubiks-cube). I modified the code so it is easier to integrate into my project.

## Backend

The backend will consist of the two language servers, and an executor. The language servers will based on ANTLR4 generated lexers and parsers. Also, it is possible that in the future, I will implement the LSP for them, so they will be more compatible with other editors. The executor will be able to create a list of moves to solve the Rubik cube, from a given state, and a given algorithm description. The backend is written in Golang

## Usage

This section contains information about the usage of the software.

### Beginner state description

The beginner state description consists of 6 3x3 matrices (for a 3x3 cube), where one element of the matrix contains the color of the corresponding piece. The relative orientation of the matrices towards each other is the following:

- the upper row of the front side is adjacent with the lower row of the up side
- the lower row of the front side is adjacent with the upper row of the up side
- the front, left, back and right sides build up a "circle" around the cube, so each side's left column is adjacent to the next side's right column.

For easier understanding here are the following tables:

Table 1. One side

```
+---+---+---+
| 0 | 1 | 2 |
+---+---+---+
| 3 | 4 | 5 |
+---+---+---+
| 6 | 7 | 8 |
+---+---+---+
```

Where the 0-2 is the upper row (U), 6-8 is the lower row (D), 0, 3 and 6 is the left column (L), 2, 5 and 8 is the right column (C).

To simplify the tables, the previous table can be written in the following way:

```
+---------+
|    U    |
|L   F   R|
|    D    |
+---------+
```

Where F is for Front side

Table 2. The relative orientation of the sides

```
+---------+
|    U    |
|L   U   R|
|    D    |
+---------+---------+---------+---------+
|    U    |    U    |    U    |    U    |
|L   F   R|L   R   R|L   B   R|L   L   R|
|    D    |    D    |    D    |    D    |
+---------+---------+---------+---------+
|    U    |
|L   D   R|
|    D    |
+---------+
```

## Credits

Thanks to [Jonathan Taylor](https://github.com/taylorjg), for licensing his project, so I can use his code!
