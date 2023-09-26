const beginnerWithParenthesis = `puzzle: cube
size: 3
state description: beginner
state: {
    front:
        (b b b)
        (b b b)
        (b b b)
    up: 
        (r r r)
        (r r r)
        (r r r)
    right: 
        (w w w)
        (w w w)
        (w w w)
    left: 
        (y y y)
        (y y y)
        (y y y)
    back: 
        (g g g)
        (g g g)
        (g g g)
    down: 
        (o o o)
        (o o o)
        (o o o)
}`

const advanced = `puzzle: cube
size: 3
state description: advanced
state:
    corners:
        up: boy wbr ygr gow
        down: wrg gyo obw rby
    edges:
        up: rb bo bw og
        middle: wr yr rg by 
        down: gw gy yo ow`

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
    `

export const beginnerConfig = beginnerWithParenthesis;
export const advancedConfig = advanced;
export const algorithm = algo;