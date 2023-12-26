package main

func RunPart1(plane Plane) int {
	plane.Tilt()
	return plane.TotalLoad()
}
