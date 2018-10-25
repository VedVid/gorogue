/*
Copyright (c) 2018 Tomasz "VedVid" Nowakowski

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

import "math"

type Node struct {
	x, y   int
	weight int
}

func (c *Creature) MoveTowardsPath(b Board, tx, ty int) {
	const startWeight = -1
	const goalWeight = -2
	const emptyWeight = -3
	var nodesEmpty = []Node{}  //for unvisited cells
	var nodesFilled = []Node{} //for visited cells
	start := Node{c.X, c.Y, startWeight}
	goal := Node{tx, ty, goalWeight}
	goal.weight = 0                         //make it explicit
	nodesFilled = append(nodesFilled, goal) //start with goal
	//initialize nodesEmpty with board tiles
	for x := 0; x < WindowSizeX; x++ {
		for y := 0; y < WindowSizeY; y++ {
			if x == goal.X && y == goal.Y { //but skip the goal
				continue
			}
			n := Node{b[x][y].X, b[x][y].Y, emptyWeight}
			nodesEmpty = append(nodesEmpty, n)
		}
	}
	currentNode := goal
	for {
		if len(nodesEmpty) == 0 { //all map tiles visited
			break
		}
		for i := 0; i < len(nodesEmpty); i++ {
			newNode := &nodesEmpty[i]
			if newNode.X >= (currentNode.X-1) && newNode.X <= (currentNode.X+1) &&
				newNode.Y >= (currentNode.Y-1) && newNode.Y <= (currentNode.Y-1) {
				//it's neightbour Node to currentNode
			}
		}
	}
}

func (c *Creature) MoveTowardsDumb(b Board, tx, ty int) {
	/*MoveTowardsDumb is Creature method;
	  it is main part of creature pathfinding. It is very simple algorithm that
	  is not supposed to replace good, old A-Star.*/
	dx := tx - c.X
	dy := ty - c.Y
	ddx, ddy := 0, 0
	if dx > 0 {
		ddx = 1
	} else if dx < 0 {
		ddx = (-1)
	}
	if dy > 0 {
		ddy = 1
	} else if dy < 0 {
		ddy = (-1)
	}
	if b[c.X+ddx][c.Y+ddy].Blocked == false {
		c.Move(ddx, ddy, b)
	} else {
		if ddx != 0 {
			if b[c.X+ddx][c.Y].Blocked == false {
				c.Move(ddx, 0, b)
			}
		} else if ddy != 0 {
			if b[c.X][c.Y+ddy].Blocked == false {
				c.Move(0, ddy, b)
			}
		}
	}
}

func (c *Creature) DistanceTo(tx, ty int) int {
	/*DistanceTo is Creature method. It takes target x and target y as args;
	  computes then returns distance from receiver to target.*/
	dx := float64(tx - c.X)
	dy := float64(ty - c.Y)
	return RoundFloatToInt(math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2)))
}
