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
step green_up:
    do: x y'
step green_center:
    goal: green([Up 1 1, Up 1 2, Up 2 1, Up 2 2])
    runs: 12
    branches:
        if green([Up 1 2, Up 2 2]):
            do: U2
        if green(Up 2 2):
            do: (2R') F' (2R)
        if green(Front 1 1) and green(Up 1 2):
            do: (2R') F (2R) U2
        if green(Front 1 1) and none(green(?), [Up 1 2, Up 2 2]):
            do: F (2R)
        if green(Front 1 2):
            do: F'
        if green(Front 2 1):
            do: F
        if green(Front 2 2):
            do: F2
        if any(green(?), [Down 1 2, Down 2 2]):
            do: (2R) F2 (2R')
        if any(green(?), [Down 1 1, Down 2 1]):
            do: D2 (2R) F2 (2R') 
        if any(green(?), [Back 1 1, Back 2 1]):
            do: (2R2) F2 (2R2)
        if any(green(?), [Back 1 2, Back 2 2]):
            do: B2 (2R2) F2 (2R2)
step red_up:
    do: x
step red_center:
    goal: red([Up 1 1, Up 1 2, Up 2 1, Up 2 2])
    runs: 12
    branches:
        if red([Up 1 2, Up 2 2]):
            do: U2
        if red(Up 2 2):
            do: (2R') F' (2R)
        if red(Front 1 1) and red(Up 1 2):
            do: (2R') F (2R) U2
        if red(Front 1 1) and none(red(?), [Up 1 2, Up 2 2]):
            do: (2R') F (2R)
        if red(Front 1 2):
            do: F'
        if red(Front 2 1):
            do: F
        if red(Front 2 2):
            do: F2
        if any(red(?), [Down 1 2, Down 2 2]):
            do: (2R) F2 (2R')
        if any(red(?), [Down 1 1, Down 2 1]):
            do: D2 (2R) F2 (2R') 
step blue_up:
    do: x
step last_centers:
    goal: blue([Up 1 1, Up 1 2, Up 2 1, Up 2 2])
    runs: 7
    branches:
        if none(blue(?), [Up 1 1, Up 1 2, Up 2 1, Up 2 2]):
            do: (2R) U2 (2R') (2L') U2 (2L)
        if blue([Up 1 2, Up 2 2]):
            do: U2
        if blue([Up 1 1, Up 1 2]) and not blue([Up 1 1, Up 2 1]):
            do: U'
        if blue([Up 2 1, Up 2 2]) and not blue([Up 1 1, Up 2 1]):
            do: U
        if blue(Up 2 2):
            do: (2R') F' (2R)
        if blue(Up 1 1) and not blue(Up 2 1):
            do: (2L) F (2L') F' (2R') F (2R)
        if blue(Up 2 1) and not blue([Up 1 2, Up 1 1]):
            do: (2U)
        if blue([Up 1 2, Front 1 1]):
            do: (2R') F (2R) U2
        if blue(Front 1 2):
            do: F'
        if blue(Front 2 1):
            do: F
        if blue(Front 2 2):
            do: F2
step pair_edges:
    goal: same([Up 0 1, Up 0 2]) and same([Up 1 0, Up 2 0]) and same([Up 1 3, Up 2 3]) and same([Up 3 1, Up 3 2]) and
          same([Front 0 1, Front 0 2]) and same([Front 1 0, Front 2 0]) and same([Front 1 3, Front 2 3]) and same([Front 3 1, Front 3 2]) and
          same([Right 0 1, Right 0 2]) and same([Right 1 0, Right 2 0]) and same([Right 1 3, Right 2 3]) and same([Right 3 1, Right 3 2]) and
          same([Back 0 1, Back 0 2]) and same([Back 1 0, Back 2 0]) and same([Back 1 3, Back 2 3]) and same([Back 3 1, Back 3 2]) and
          same([Left 0 1, Left 0 2]) and same([Left 1 0, Left 2 0]) and same([Left 1 3, Left 2 3]) and same([Left 3 1, Left 3 2]) and
          same([Down 0 1, Down 0 2]) and same([Down 1 0, Down 2 0]) and same([Down 1 3, Down 2 3]) and same([Down 3 1, Down 3 2])
    runs: 40
    helpers:
        pair: (2L') U2 (2L') U2 F2 (2L') F2 (2R) U2 (2R') U2 (2L2)
    branches:
        if same([Up 0 1, Up 0 2]) and same([Back 0 1, Back 0 2]) and
           same([Up 1 0, Up 2 0]) and same([Left 0 1, Left 0 2]) and
           same([Up 1 3, Up 2 3]) and same([Right 0 1, Right 0 2]) and
           same([Up 3 1, Up 3 2]) and same([Front 0 1, Front 0 2]) and
           same([Down 0 1, Down 0 2]) and same([Front 3 1, Front 3 2]) and
           same([Down 1 0, Down 2 0]) and same([Left 3 1, Front 3 2]) and
           same([Down 1 3, Down 2 3]) and same([Right 3 1, Right 3 2]) and
           same([Down 3 1, Down 3 2]) and same([Back 3 1, Back 3 2]):
            do: z
        if same([Up 0 1, Up 0 2]) and same([Back 0 1, Back 0 2]) and
           same([Up 1 0, Up 2 0]) and same([Left 0 1, Left 0 2]) and
           same([Up 1 3, Up 2 3]) and same([Right 0 1, Right 0 2]) and
           same([Up 3 1, Up 3 2]) and same([Front 0 1, Front 0 2]):
            do: x2
        if same([Up 0 1, Up 0 2]) and same([Back 0 1, Back 0 2]):
            do: U
        if same([Up 0 2, Front 0 1]) and same([Back 0 1, Up 3 1]):
            do: pair U
        if same([Up 0 2, Up 3 2]) and same([Back 0 1, Front 0 2]):
            do: F U' R U
        if same([Up 0 2, Left 0 1]) and same([Back 0 1, Up 1 0]):
            do: B' U' B
        if same([Up 0 2, Up 2 0]) and same([Back 0 1, Left 0 2]):
            do: L F
        if same([Up 0 2, Right 0 1]) and same([Back 0 1, Up 2 3]):
            do: B U B'
        if same([Up 0 2, Up 1 3]) and same([Back 0 1, Right 0 2]):
            do: R' F'
        if same([Up 0 2, Front 2 0]) and same([Back 0 1, Left 2 3]):
            do: F
        if same([Up 0 2, Left 1 3]) and same([Back 0 1, Front 1 0]):
            do: L D F2
        if same([Up 0 2, Front 1 3]) and same([Back 0 1, Right 1 0]):
            do: F'
        if same([Up 0 2, Right 2 0]) and same([Back 0 1, Front 2 3]):
            do: R' D' F2
        if same([Up 0 2, Right 1 3]) and same([Back 0 1, Back 1 0]):
            do: R' B U B'
        if same([Up 0 2, Back 2 0]) and same([Back 0 1, Right 2 3]):
            do: R2 F'
        if same([Up 0 2, Back 1 3]) and same([Back 0 1, Left 1 0]):
            do: L2 F
        if same([Up 0 2, Left 2 0]) and same([Back 0 1, Back 2 3]):
            do: L' D F2
        if same([Up 0 2, Front 3 2]) and same([Back 0 1, Down 0 2]):
            do: F2
        if same([Up 0 2, Down 0 1]) and same([Back 0 1, Front 3 1]):
            do: D' L' F
        if same([Up 0 2, Right 3 2]) and same([Back 0 1, Down 3 2]):
            do: D' F2
        if same([Up 0 2, Down 1 3]) and same([Back 0 1, Right 3 1]):
            do: R F'
        if same([Up 0 2, Back 3 2]) and same([Back 0 1, Down 3 1]):
            do: D2 F2
        if same([Up 0 2, Down 3 2]) and same([Back 0 1, Back 3 1]):
            do: D L' F
        if same([Up 0 2, Left 3 2]) and same([Back 0 1, Down 1 0]):
            do: D F2
        if same([Up 0 2, Down 2 0]) and same([Back 0 1, Left 3 1]):
            do: L' F    
        prepare: D

        
        

         
`

export const advancedConfig = advanced;
export const algorithm = algo;