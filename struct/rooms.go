package _struct

type Room struct {
	width  float32
	length float32
	height float32
}

func (room Room) getSquare() float32 {
	return room.width * room.length
}

func (room Room) getVolume() float32 {
	return room.width * room.length * room.height
}

func (room Room) getPerimeter() float32 {
	return 2 * (room.width + room.length)
}

func createRooms() []Room {
	var rooms []Room

	rooms[0] = Room{width: 5.4, length: 3.4, height: 2.5}
	rooms[0] = Room{width: 3.2, length: 4.1, height: 3}
	rooms[0] = Room{width: 4.7, length: 3.2, height: 2.8}

	return rooms
}
