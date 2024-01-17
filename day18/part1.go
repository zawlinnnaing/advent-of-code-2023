package main

func RunPart1(plans []DigPlan) int {
	vertices := getDigVertices(plans)
	return getInterior(vertices)
}
