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
        prepare:
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
        prepare:
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
        prepare:
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
        prepare:
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
        prepare:
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
        prepare:
    step white_up:
        goal: white(Up 1 1)
        runs: 4
        branches:
            if white(Down 1 1):
                do: upsideDown
            if white(Front 1 1):
                do: x
            prepare: y
    step white_cross:
        goal: orientation([(Up, Front), (Up, Right), (Up, Back), (Up, Left)])
        runs: 20
        branches:
            if orientation(Up, Front):
                do: y
            if piece(Up, Front) at (Up, Front):
                do: F R' D' R F2 y
            if piece(Up, Front) at (Front, Down):
                do: F2
            if piece(Up, Front) at (Up, Right):
                do: R2 D'
            if piece(Up, Front) at (Up, Back):
                do: B2 D2
            if piece(Up, Front) at (Up, Left):
                do: L2 D
            if piece(Up, Front) at (Right, Front):
                do: F
            if piece(Up, Front) at (Front, Left):
                do: F'
            if piece(Up, Front) at (Left, Back):
                do: L' D L
            if piece(Up, Front) at (Back, Right):
                do: R D' R'
            prepare: D
    step white_side:
        goal: orientation([(Up, Front, Right), (Up, Right, Back), (Up, Back, Left), (Up, Left, Front)])
        runs: 16
        branches:
            if orientation(Up, Front, Right):
                do: y
            if piece(Up, Front, Right) like position(Right, Up, Front):
                do: R' D' R D
            if piece(Up, Front, Right) like pos(Front, Right, Up):
                do: R' D R
            if (Up, Front, Right) like position(Down, Right, Front):
                do: R' D2 R D 
            if (Up, Front, Right) like (Front, Down, Right):
                do: D' R' D R y
            if piece(Up, Front, Right) like (Right, Front, Down):
                do: D F D' F' y
            if piece(Up, Front, Right) at (Up, Back, Right):
                do: B' D' B
            if piece(Up, Front, Right) at (Up, Back, Left):
                do: L' D2 L
            if piece(Up, Front, Right) at (Up, Left, Front):
                do: L D L'
            prepare: D
    step second_layer:
        goal: orientation([(Front, Right), (Right, Back), (Back, Left), (Left, Front)])
        helpers:
            right: D' R' D R D F D' F'
            left: D L D' L' D' F' D F
        runs: 16
        branches:
            if orientation(Front, Right):
                do: y
            if place(Front, Right) or piece(Front, Right) like (Front, Down):
                do: right
            if piece(Front, Right) like (Down, Front):
                do: D y left
            if piece(Front, Right) at (Front, Left):
                do: left D2
            if piece(Front, Right) at (Left, Back):
                do: y' left y D'
            if piece(Front, Right) at (Back, Right):
                do: y right y' D
            prepare: D
    step yellow_up_:
        do: upsideDown
    step yellow_cross:
        goal: yellow([Up 0 1, Up 0 2, Up 1 0, Up 1 3, Up 2 0, Up 2 3, Up 3 1, Up 3 2])
        helpers:
            fromL: F U R U' R' F'
            fromDash: F R U R' U' F'
        runs: 6
        branches:
            if yellow([Up 0 1, Up 0 2, Up 1 0, Up 2 0, Up 1 3, Up 2 3]) and yellow([Front 0 1, Front 0 2]):
                do: (2R2) B2 U2 (2L) U2 (2R') U2 (2R) U2 F2 (2R) F2 (2L') B2 (2R2)
            if yellow([Up 0 1, Up 0 2, Up 1 3, Up 2 3, Up 3 1, Up 3 2]) and yellow([Left 0 1, Left 0 2]):
                do: U'
            if yellow([Up 1 0, Up 2 0, Up 1 3, Up 2 3, Up 3 1, Up 3 2]) and yellow([Back 0 1, Back 0 2]):
                do: U2
            if yellow([Up 0 1, Up 0 2, Up 1 0, Up 2 0, Up 3 1, Up 3 2]) and yellow([Right 0 1, Right 0 2]):
                do: U
            if yellow([Up 0 1, Up 0 2, Up 1 0, Up 2 0]):
                do: fromL
            if yellow([Up 1 0, Up 2 0, Up 1 3, Up 2 3]):
                do: fromDash
            prepare: U
    step yellow_edges:
        goal: orientation([(Up, Front), (Up, Right), (Up, Back), (Up, Left)])
        helpers:
            swapTwo: R U R' U R U2 R' U
        runs: 8
        branches:
            if orientation([(Up, Back), (Up, Right)]):
                do: swapTwo
            if orientation([(Up, Back), (Up, Left)]):
                do: y
            if orientation([(Up, Front), (Up, Right)]):
                do: y'
            if orientation([(Up, Front), (Up, Left)]):
                do: y2
            prepare:
                do: U
                consecutive: 3
            prepare: swapTwo
    step yellow_corners:
        goal: place([(Up, Front, Right), (Up, Right, Back), (Up, Back, Left), (Up, Left, Front)])
        helpers:
            cycleCorners: U R U' L' U R' U' L
        runs: 9
        branches:
            if place([(Up, Front, Right), (Up, Front, Left)]):
                do: y2 F2 R2 B' D' B R2 F' U (2F2) F' L2 (2F2) L2 (2L2) (2F2) (2L2) U'
            if place(Up, Front, Right):
                do: cycleCorners
            prepare: 
                do: y
                consecutive: 3
            prepare: cycleCorners
    step yellow_corners_orient:
        goal: orientation([(Up, Front, Right), (Up, Right, Back), (Up, Back, Left), (Up, Left, Front)])
        runs: 4
        branches:
            if yellow(Up 3 3):
                do: U
            if yellow(Front 0 3):
                do: 4(R' D' R D) U
            if yellow(Right 0 0):
                do: 2(R' D' R D) U
            prepare:
    step orient_top_layer:
        goal: orientation(Up, Front, Right)
        runs: 3
        do: U        

        
        

        
`

export const advancedConfig = advanced;
export const algorithm = algo;