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

export const beginnerConfig = beginnerWithParenthesis;
export const advancedConfig = advanced;