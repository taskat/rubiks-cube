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

### Advanced state description

The advanced state description takes advantage of the Rubik's cube's special properties. It is based on the following facts:
- The middle elements of each side will always be the same color, relative to each other.
- The edge pieces's two side will always be adjacent to each other (the pieces can not break)
- The corner pieces's three side will always be adjacent to each other (the pieces can not break)
- The corner pieces's three side will always be in the same order (the pieces can not flip)

Based on these properties, we can conclude, that in order to get the cube's state, we do not need all the colors of the middle elements, we only need to know two middle elements's color, which are adjacent. Since the cube can always be rotated this state description always assumes that the up side is white, and the front side is blue. From this we can conclude, that the right side is orange, the back side is green, the left side is red, and the down side is yellow.

Secondly to describe the corners, we can takes advantage that there is only 8 possible places for a corner piece, and for every position, there is 3 different orientation. So we can have a two lists of the cube pieces, one for the upper layer and for the lower layer. Both list starts from the back left corner and goes around clockwise. For the lower layer, the cube should be temporarily flipped around the left-rigth axis, so that the down side becomes the up side. These list will decide the position of each corner piece, and orientation will be decided by the order of the colors in the element. The colors should start from the up side (for the lower layer, the down side) and move around clockwise. For example if the qhite-red-green corner's white side is on the up side, then this piece's notation will be wrg.

Thirdly the edge pieces's description is based on the same principle as the corner pieces. However there are a couple differences. First, there is a third layer for edge pieces, the middle layer. Second for the upper and lower layer, the starting element is the back edge (for the lower layer, the cube should be flipped the same way as for the corners), and for the middle layer it is the left-back edge. And third the orientation should start with the up side's color for the upper and lower layers, and from the left side of the left-back edge for the middle layer. In the middle layer, the pieces should traversed clockwise

```

## Credits

Thanks to [Jonathan Taylor](https://github.com/taylorjg), for licensing his project, so I can use his code!
