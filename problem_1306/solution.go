package main

func canReach(arr []int, start int) bool {
	visited := make(map[int]bool)

	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if arr[current] == 0 {
			return true
		}

		front := current + arr[current]
		behind := current - arr[current]

		if front < len(arr) && !visited[front] {
			visited[front] = true
			queue = append(queue, front)
		}

		if behind >= 0 && !visited[behind] {
			visited[behind] = true
			queue = append(queue, behind)
		}
	}

	return false
}
