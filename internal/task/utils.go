package task

func NextID(tasks []Task) int {
	maxID := 0

	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	return maxID + 1
}
