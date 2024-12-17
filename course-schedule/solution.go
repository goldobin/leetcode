package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	var (
		pmap      = make(map[int][]int)
		traversed = make([]bool, numCourses)
		visited   = make([]bool, numCourses)
	)

	for _, p := range prerequisites {
		pmap[p[0]] = append(pmap[p[0]], p[1])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(pmap, traversed, visited, i) {
			return false
		}
	}

	return true
}

func dfs(pmap map[int][]int, traversed []bool, visited []bool, i int) bool {
	if traversed[i] {
		return true
	}

	visited[i] = true
	defer func() {
		visited[i] = false
		traversed[i] = true
	}()
	ps, ok := pmap[i]

	if !ok {
		return true
	}

	for _, next := range ps {
		if traversed[next] {
			continue
		}
		if visited[next] {
			return false
		}
		if !dfs(pmap, traversed, visited, next) {
			return false
		}
	}

	return true
}
