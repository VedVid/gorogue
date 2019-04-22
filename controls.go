/*
Copyright (c) 2018, Tomasz "VedVid" Nowakowski
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	blt "bearlibterminal"
)

const (
	StrMoveNorth = "MOVE_NORTH"
	StrMoveWest  = "MOVE_WEST"
	StrMoveEast  = "MOVE_EAST"
	StrMoveSouth = "MOVE_SOUTH"

	StrTarget = "TARGET"
	StrLook   = "LOOK"
	StrPickup = "PICKUP"

	StrInventory = "INVENTORY"
	StrEquipment = "EQUIPMENT"
)

var Actions = []string{
	StrMoveNorth,
	StrMoveWest,
	StrMoveEast,
	StrMoveSouth,
	StrTarget,
	StrLook,
	StrPickup,
	StrInventory,
	StrEquipment,
}

var CommandKeys = map[int]string{
	blt.TK_UP:    StrMoveNorth,
	blt.TK_RIGHT: StrMoveEast,
	blt.TK_DOWN:  StrMoveSouth,
	blt.TK_LEFT:  StrMoveWest,
	blt.TK_F:     StrTarget,
	blt.TK_L:     StrLook,
	blt.TK_G:     StrPickup,
	blt.TK_I:     StrInventory,
	blt.TK_E:     StrEquipment,
}

var CustomCommandKeys = map[int]string{}

func Command(com string, p *Creature, b *Board, c *Creatures, o *Objects) bool {
	turnSpent := false
	switch com {
	case StrMoveNorth:
		turnSpent = p.MoveOrAttack(0, -1, *b, *c)
	case StrMoveEast:
		turnSpent = p.MoveOrAttack(1, 0, *b, *c)
	case StrMoveSouth:
		turnSpent = p.MoveOrAttack(0, 1, *b, *c)
	case StrMoveWest:
		turnSpent = p.MoveOrAttack(-1, 0, *b, *c)

	case StrTarget:
		turnSpent = p.Target(*b, *o, *c)
	case StrLook:
		p.Look(*b, *o, *c)
	case StrPickup:
		turnSpent = p.PickUp(o)
	case StrInventory:
		turnSpent = p.InventoryMenu(o)
	case StrEquipment:
		turnSpent = p.EquipmentMenu(o)
	}
	return turnSpent
}

func Controls(k int, p *Creature, b *Board, c *Creatures, o *Objects) bool {
	turnSpent := false
	var command string
	if CustomControls == false {
		command = CommandKeys[k]
	} else {
		command = CustomCommandKeys[k]
	}
	turnSpent = Command(command, p, b, c, o)
	return turnSpent
}

func ReadInput() int {
	key := blt.Read()
	for _, v := range HardcodedKeys {
		if key == v {
			return v
		}
	}
	var r rune
	if blt.Check(blt.TK_WCHAR) != 0 {
		r = rune(blt.State(blt.TK_WCHAR))
	}
	return KeyMap[r]
}
