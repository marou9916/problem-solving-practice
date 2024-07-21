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

	for _, room := range currentRoom.linkedRooms {
		if room.name == endRoom.name {
			pathCounterBetweenTwoRooms++
		} else {
			if IsAlreadyInTheCurrentPath(currentPath, room) {
				continue
			} else {
				currentPath = append(currentPath, room.name)
				currentRoom = room                  // no need for this line
				CountPaths(currentRoom, secondRoom) // each recursive call explore a different option, need to store the sum of all that options into the counter variable
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
