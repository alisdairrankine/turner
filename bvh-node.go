package turner

import (
	"math/rand"
	"sort"
)

type bvhNode struct {
	Left  Hitable
	Right Hitable
	Box   *AABB
}

func BVHNode(hitables []Hitable, time0, time1 float64) Hitable {
	node := &bvhNode{}

	if len(hitables) == 1 {
		hit := hitables[0]
		node.Left = hit
		node.Right = hit
		if box, ok := hit.BoundingBox(time0, time1); ok {
			node.Box = box
		}
		return node
	}

	list := axisSort(hitables, int(3*rand.Float64()), time0, time1)
	if len(list) == 2 {
		node.Left = list[0]
		node.Right = list[1]
	} else {
		node.Left = BVHNode(list[:len(list)/2], time0, time1)
		node.Right = BVHNode(list[len(list)/2:], time0, time1)
	}

	leftBox, leftOK := node.Left.BoundingBox(time0, time1)
	rightBox, rightOK := node.Right.BoundingBox(time0, time1)
	if leftOK && rightOK {
		node.Box = SurroundingBox(leftBox, rightBox)
	} else {
		panic("sort your boxes out")
	}

	return node
}

func axisSort(hitables []Hitable, axis int, time0, time1 float64) []Hitable {

	comparator := func(i, j int) bool {
		a := hitables[i]
		b := hitables[j]
		leftBox, leftOK := a.BoundingBox(time0, time1)
		rightBox, rightOK := b.BoundingBox(time0, time1)
		if !leftOK || !rightOK {
			panic("cannot sort hitable without bounding box")
		}
		value := leftBox.Min.X - rightBox.Min.X
		if axis == 1 {
			value = leftBox.Min.Y - rightBox.Min.Y
		} else if axis == 2 {
			value = leftBox.Min.Z - rightBox.Min.Z
		}

		if value < 0 {
			return true
		}
		return false
	}

	sort.Slice(hitables, comparator)
	return hitables
}

func (b bvhNode) Hit(r *Ray, tMin, tMax float64) (*HitRecord, bool) {
	record := &HitRecord{}
	if _, hit := b.Box.Hit(r, tMin, tMax); hit {

		leftRec, leftHit := b.Left.Hit(r, tMin, tMax)
		rightRec, rightHit := b.Right.Hit(r, tMin, tMax)

		if leftHit && rightHit {
			if leftRec.T < rightRec.T {
				record = leftRec
			} else {
				record = rightRec
			}
			return record, true
		} else if leftHit {
			return leftRec, true
		} else if rightHit {
			return rightRec, true
		}

	}
	return nil, false
}
func (b bvhNode) BoundingBox(t0, t1 float64) (*AABB, bool) {
	return b.Box, true
}
