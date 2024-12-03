package main

import (
	"strconv"
	"strings"
)

type hail struct {
	p vector
	v vector
}

func newHail(s string) *hail {
	strs := strings.Split(s, "@")
	pStrs := strs[0]
	vStrs := strs[1]

	ps := strings.Split(pStrs, ",")
	pxStr := strings.TrimSpace(ps[0])
	pyStr := strings.TrimSpace(ps[1])
	pzStr := strings.TrimSpace(ps[2])

	vs := strings.Split(vStrs, ",")
	vxStr := strings.TrimSpace(vs[0])
	vyStr := strings.TrimSpace(vs[1])
	vzStr := strings.TrimSpace(vs[2])

	px, _ := strconv.Atoi(pxStr)
	py, _ := strconv.Atoi(pyStr)
	pz, _ := strconv.Atoi(pzStr)

	vx, _ := strconv.Atoi(vxStr)
	vy, _ := strconv.Atoi(vyStr)
	vz, _ := strconv.Atoi(vzStr)

	return &hail{
		p: vector{
			x: px,
			y: py,
			z: pz,
		},
		v: vector{
			x: vx,
			y: vy,
			z: vz,
		},
	}
}

func (h1 *hail) intersect(h2 *hail) (vector, bool) {
	vCross := h1.v.cross(h2.v)

	if vCross.x == 0 && vCross.y == 0 && vCross.z == 0 {
		return vector{}, false
	}

	pDiff := h2.p.subtract(h1.p)
	tNumerator := pDiff.cross(h2.v).dot(vCross)
	tDenominator := vCross.dot(vCross)
	t := tNumerator / tDenominator

	sNumerator := pDiff.cross(h1.v).dot(vCross)
	sDenominator := vCross.dot(vCross)
	s := sNumerator / sDenominator

	h1p := h1.p.add(h1.v.scale(t))
	h2p := h2.p.add(h2.v.scale(s))

	if h1p.equals(h2p) {
		return h1p, true
	}

	return vector{}, false
}

func (v1 vector) equals(v2 vector) bool {
	return v1.x == v2.x && v1.y == v2.y && v1.z == v2.z
}

func (p *vector) within(xyMin, xyMax int) bool {
	return p.x >= xyMin && p.x <= xyMax && p.y >= xyMin && p.y <= xyMax
}

type vector struct {
	x int
	y int
	z int
}

func (v1 vector) subtract(v2 vector) vector {
	return vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
		z: v1.z - v2.z,
	}
}

func (v1 vector) cross(v2 vector) vector {
	return vector{
		x: v1.y*v2.z - v1.z*v2.y,
		y: v1.z*v2.x - v1.x*v2.z,
		z: v1.x*v2.y - v1.y*v2.x,
	}
}

func (v1 vector) dot(v2 vector) int {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

func (v vector) scale(scalar int) vector {
	return vector{
		x: v.x * scalar,
		y: v.y * scalar,
		z: v.z * scalar,
	}
}

func (v1 vector) add(v2 vector) vector {
	return vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}
