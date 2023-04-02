package cube

import "encoding/json"

type CubeJSON struct {
	Sides map[CubeSide]Side `json:"sides"`
	Size  int               `json:"size"`
}

func NewCubeJSON(cube *Cube) CubeJSON {
	return CubeJSON{
		Sides: cube.sides,
		Size:  cube.size,
	}
}

func (c CubeJSON) marshal() ([]byte, error) {
	return json.Marshal(c)
}
