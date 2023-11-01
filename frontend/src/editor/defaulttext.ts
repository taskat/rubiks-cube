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
    step yellow_up:
        do: upsideDown
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
            if yellow(Up 1 1):
                do: U
            if yellow(Front 0 1):
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