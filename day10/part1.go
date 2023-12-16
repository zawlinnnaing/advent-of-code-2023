package main

func RunPart1(maze [][]string) int {
	startPoint := findStartPoint(maze)
	seen := make([][]bool, 0)
	for i := 0; i < len(maze[0]); i++ {
		innerSeen := make([]bool, len(maze[0]))
		seen = append(seen, innerSeen)
	}
	path := Path{points: make([]Point, 0)}
	walk(maze, startPoint, seen, true, &path)
	return path.Len() / 2
}
