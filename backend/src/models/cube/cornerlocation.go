package cube

import (
	"sort"
	"strings"
)

type cornerLocation struct {
	sides [3]CubeSide
	hash  string
}

func newCornerLocation(s1, s2, s3 CubeSide) cornerLocation {
	return cornerLocation{sides: [3]CubeSide{s1, s2, s3}, hash: ""}
}

func (c cornerLocation) getHash() string {
	if c.hash != "" {
		return c.hash
	}
	hashes := []string{c.sides[0].getHash(), c.sides[1].getHash(), c.sides[2].getHash()}
	sort.Strings(hashes)
	return strings.Join(hashes, "")
}
