package main

func RunPart1(plans []DigPlan) int {
	vertices := getDigVertices(plans)
	area := getArea(vertices)
	return getInterior(area, len(vertices))
}
