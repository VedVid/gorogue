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

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

	blt "bearlibterminal"
)

const (
	// Colors.
	colorPlayer     = "white"
	colorPlayerDark = "white"
)

func NewPlayer(layer, x, y int, character, color, colorDark string,
	alwaysVisible, blocked, blocksSight bool, ai, hp, attack,
	defense int, equipment EquipmentComponent) (*Creature, error) {
	/* Function NewPlayer takes all values necessary by its struct,
	   and creates then returns pointer to Creature;
	   so, it's basically NewCreature function. */
	var err error
	if layer < 0 {
		txt := LayerError(layer)
		err = errors.New("Player layer is smaller than 0." + txt)
	}
	if x < 0 || x >= MapSizeX || y < 0 || y >= MapSizeY {
		txt := CoordsError(x, y)
		err = errors.New("Player coords is out of window range." + txt)
	}
	if utf8.RuneCountInString(character) != 1 {
		txt := CharacterLengthError(character)
		err = errors.New("Player character string length is not equal to 1." + txt)
	}
	if ai != PlayerAI {
		txt := PlayerAIError(ai)
		err = errors.New("Warning: Player AI is supposed to be " +
			strconv.Itoa(PlayerAI) + "." + txt)
	}
	if hp < 0 {
		txt := InitialHPError(hp)
		err = errors.New("Player HPMax is smaller than 0." + txt)
	}
	if attack < 0 {
		txt := InitialAttackError(attack)
		err = errors.New("Player attack value is smaller than 0." + txt)
	}
	if defense < 0 {
		txt := InitialDefenseError(defense)
		err = errors.New("Player defense value is smaller than 0." + txt)
	}
	playerBasicProperties := BasicProperties{layer, x, y, character, color,
		colorDark}
	playerVisibilityProperties := VisibilityProperties{alwaysVisible}
	playerCollisionProperties := CollisionProperties{blocked, blocksSight}
	playerAIProperties := AIProperties{ai}
	playerFighterProperties := FighterProperties{hp, hp, attack, defense}
	playerNew := &Creature{playerBasicProperties, playerVisibilityProperties,
		playerCollisionProperties, playerAIProperties, playerFighterProperties,
		equipment}
	return playerNew, err
}

func (p *Creature) InventoryMenu(o *Objects) bool {
	/* Inventory menu is method of Creature (that is supposed to be a player).
	   It calls PrintInventoryMenu (that have much better docstring).
	   It returns false as printing menu doesn't spent turn, but it may change
	   in near future, because using / equipping items will spent turn. */
	PrintInventoryMenu(UIPosX, UIPosY, "Inventory:", p.Inventory)
	turnSpent := p.HandleInventory(o, KeyToOrder(blt.Read())) //it's ugly one-liner
	return turnSpent
	}

func (p *Creature) EquipmentMenu() bool {
	/* Same as InventoryMenu method. */
	PrintEquipmentMenu(UIPosX, UIPosY, "Equipment:", Objects{p.Slot})
	return false
}

func (p *Creature) HandleInventory(o *Objects, option int) bool {
	turnSpent := false
	if option <= len(p.Inventory) { //valid input
		turnSpent = p.InventoryActions(o, option)
	}
	return turnSpent
}

func (p *Creature) InventoryActions(o *Objects, option int) bool {
	//it's very basic example; it should create additional menu
	//to choose to drop or use item or whatever is possible to do with it
	//but it won't just now as it's kind of proof-of-concept
	turnSpent := false
	fmt.Println("You are working with: " + p.Inventory[option].Char)
	for {
		key := blt.Read()
		if key == blt.TK_ENTER {
			p.Drop(o, option)
			turnSpent = true
			break
		} else if key == blt.TK_ESCAPE {
			fmt.Println("You closed the menu.")
			turnSpent = false
			break
		}
	}
	return turnSpent
}
