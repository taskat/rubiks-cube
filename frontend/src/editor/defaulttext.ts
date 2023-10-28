const advanced = `puzzle: cube
size: 2
state description: advanced
state:
    corners:
        Up: boy wbr ygr gow
        Down: wrg gyo obw rby`

const algo = `helpers:
    upsideDown: x2
steps:
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
    step yellow_up:
        do: upsideDown
    step yellow_cross:
        goal: yellow([Up 0 1, Up 1 0, Up 1 2, Up 2 1])
        helpers:
            fromL: F U R U' R' F'
            fromDash: F R U R' U' F'
        runs: 4
        branches:
            if yellow([Up 0 1, Up 1 0]):
                do: fromL
            if yellow([Up 1 0, Up 1 2]):
                do: fromDash
            if none(yellow(?), [Up 0 1, Up 1 0, Up 1 2, Up 2 1]):
                do: fromDash U2 fromL
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
            if yellow(Up 2 2):
                do: U
            if yellow(Front 0 2):
                do: 4(R' D' R D) U
            if yellow(Right 0 0):
                do: 2(R' D' R D) U
    step orient_top_layer:
        goal: orientation(Up, Front, Right)
        runs: 3
        do: U      
    `

export const advancedConfig = advanced;
export const algorithm = algo;