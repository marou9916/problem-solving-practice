package backtracking

type Room struct {
	name        string
	linkedRooms map[*Room]Room
}

func CountPaths(firstRoom, secondRoom Room) int {
	startRoom := firstRoom
	endRoom := secondRoom
	currentRoom := startRoom
	var currentPath []string
	pathCounterBetweenTwoRooms := 0

	currentPath = append(currentPath, startRoom.name)

	for _, nextRoom := range currentRoom.linkedRooms {
		if nextRoom.name == endRoom.name {
			pathCounterBetweenTwoRooms++
		} else {
			if IsAlreadyInTheCurrentPath(currentPath, nextRoom) {
				continue
			} else {
				currentPath = append(currentPath, nextRoom.name)
				pathCounterBetweenTwoRooms += CountPaths(nextRoom, secondRoom)
				currentPath = currentPath[:len(currentPath)-1] //now you do backtrack
			}
		}
	}
	return pathCounterBetweenTwoRooms

}

func IsAlreadyInTheCurrentPath(currentPath []string, roomToCheck Room) bool {
	for _, name := range currentPath {
		return name == roomToCheck.name
	}
	return false
}
