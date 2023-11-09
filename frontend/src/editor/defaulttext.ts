const advanced = `puzzle: cube
size: 4
state description: beginner
state:
    Front:
        y y y r
        g r b y
        r w w o
        b r y g
    Up: 
        g r y b
        b o y b
        y y g y
        r o b w
    Right: 
        b r w y
        b y r w
        b o g b
        o w w r
    Left: 
        w w g g
        o w r w
        g b b g
        y b g o
    Back: 
        o o w r
        o b r g
        r y g w
        b g g g
    Down: 
        w b r w
        r g o o
        o w o r
        o y o y`

const algo = `helpers:
upsideDown: x2
steps:
step white_center:
    goal: white([Up 1 1, Up 1 2, Up 2 1, Up 2 2])
    runs: 12
    branches:
        if white(Up 2 2):
            do: U'
        if white(Front 2 2):
            do: (2D) (2R') (2D') (2R)
        if white(Front 1 1):
            do: F2
        if white(Front 1 2):
            do: F
        if white(Front 2 1):
            do: F'
        if any(white(?), [Right 1 1, Right 1 2, Right 2 1, Right 2 2]):
            do: (2U) (2D')
        if any(white(?), [Left 1 1, Left 1 2, Left 2 1, Left 2 2]):
            do: (2U') (2D)
        if any(white(?), [Back 1 1, Back 1 2, Back 2 1, Back 2 2]):
            do: (2U2) (2D2)
        if any(white(?), [Down 1 1, Down 1 2, Down 2 1, Down 2 2]):
            do: (2R) F2 (2R2) D2 (2R)
step yellow_up:
    do: upsideDown
step yellow_center:
    goal: yellow([Up 1 1, Up 1 2, Up 2 1, Up 2 2])
    runs: 12
    branches:
        if yellow(Up 2 2):
            do: U'
        if yellow(Front 2 2):
            do: (2D) (2R') B (2D') B' (2R)
        if yellow(Front 1 1):
            do: F2
        if yellow(Front 1 2):
            do: F
        if yellow(Front 2 1):
            do: F'
        if any(yellow(?), [Right 1 1, Right 1 2, Right 2 1, Right 2 2]):
            do: (2U) (2D')
        if any(yellow(?), [Left 1 1, Left 1 2, Left 2 1, Left 2 2]):
            do: (2U') (2D)
        if any(yellow(?), [Back 1 1, Back 1 2, Back 2 1, Back 2 2]):
            do: (2U2) (2D2) 
step blue_up:
    do: x 

`

export const advancedConfig = advanced;
export const algorithm = algo;