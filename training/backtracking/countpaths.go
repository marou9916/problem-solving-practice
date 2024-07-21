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

	for _, room := range currentRoom.linkedRooms { //you can do a better choice of variable name for 'room'. Instead, why not "nextRoom"?
		if room.name == endRoom.name {
			pathCounterBetweenTwoRooms++
		} else {
			if IsAlreadyInTheCurrentPath(currentPath, room) {
				continue
			} else {
				currentPath = append(currentPath, room.name)
				pathCounterBetweenTwoRooms += CountPaths(room, secondRoom)  //updating the counter properly after each recursive call
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
