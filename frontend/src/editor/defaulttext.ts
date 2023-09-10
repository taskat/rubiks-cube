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
        up: wrg wgo wob wbr
        down: yrb ybo yog ygr
    edges:
        up: wg wo wb wr
        middle: rg go ob br
        down: yb yo yg yr`

const algo = `helpers:
    upsideDown: x2
steps:
    step white_up:
        goal: red(Up 1 1)
        runs: 4
        branches:
            if red(Down 1 1):
                do: upsideDown
            if red(Front 1 1):
                do: x
            prepare: y
    `

export const beginnerConfig = beginnerWithParenthesis;
export const advancedConfig = advanced;
export const algorithm = algo;