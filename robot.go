package main

import (
	"fmt"
)


const (
	TABLE_DIM_X	= 5
	TABLE_DIM_Y = 5
)


// direction type related defs
var directionEnums []string;

type DirectionType uint8

func (d DirectionType) String() string {
	return directionEnums[int(d)]
}

func ciota(s string) DirectionType {
	directionEnums = append(directionEnums, s)
	return DirectionType(len(directionEnums) - 1)
}

var (
	NORTH = ciota("NORTH") 
	EAST = ciota("EAST")
	SOUTH = ciota("SOUTH")
	WEST = ciota("WEST")

	NODIR = ciota("NODIR")
)

func stringToDirectionType(dir string) DirectionType {
	if dir == directionEnums[NORTH] {
		return NORTH
	} else if dir == directionEnums[EAST] {
		return EAST
	} else if dir == directionEnums[SOUTH] {
		return SOUTH
	} else if dir == directionEnums[WEST] {
		return WEST
	} else {
		return NODIR
	} 
}

// easy mechanism for moving robot using
// pre calculated increments on axis based on direction
type MovePair struct {
	x, y interface{}
}

var movePairArray [4]MovePair


type robot struct {
	x			int
	y			int
	direction 	DirectionType
}

// Craete and initialize robot 
func newRobot() robot {
	// TODO: this should be better done in init, once per app
	movePairArray[NORTH] = MovePair{0, +1}
	movePairArray[EAST] = MovePair{+1, 0}
	movePairArray[SOUTH] = MovePair{0, -1}
	movePairArray[WEST] = MovePair{-1, 0}

	r := robot{-1, -1, NODIR}
	return r
}

// Validate given x,y to be on the table
func (r *robot) validateXY(x int, y int) bool {
	return x > -1 && x < TABLE_DIM_X && y > -1 && y < TABLE_DIM_Y;
}

// PLACE will put the toy robot on the table in position X,Y 
// and facing NORTH, SOUTH, EAST or WEST. 
func (r *robot) place(x int, y int, dir string) {
	if direction := stringToDirectionType(dir); r.validateXY(x, y) && direction != NODIR {
		r.x = x
		r.y = y
		r.direction = direction
	}
}

// MOVE will move the toy robot one unit forward in the 
// direction it is currently facing.
func (r *robot) move() {
	if r.direction == NODIR {
		return
	}

	nextX := r.x + movePairArray[r.direction].x.(int)
	nextY := r.y + movePairArray[r.direction].y.(int)

	if r.validateXY(nextX, nextY) {
		r.x = nextX;
		r.y = nextY;
	}
}

// LEFT will rotate the robot 90 degrees in the specified 
// direction without changing the position of the robot.
func (r *robot) left() {
	r.direction--
	r.direction %= 4;
}

// RIGHT will rotate the robot 90 degrees in the specified 
// direction without changing the position of the robot.
func (r *robot) right() {
	r.direction++
	r.direction %= 4;
}

// REPORT will announce the X,Y and orientation of the robot.
func (r *robot) report() {
	if r.validateXY(r.x, r.y) {
		fmt.Printf("%d,%d,%s\n", r.x, r.y, r.direction.String())
	}
}