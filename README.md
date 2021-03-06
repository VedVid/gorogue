# RAWIG

RAWIG is planned as architecture, ready to modify and expand. It won't be tutorial material nor full game.  

### Prerequisites

To work with RAWIG, you need:  
 - Go compiler  
 - C compiler  
 - BearLibTerminal library  

### Usage

RAWIG is standalone application, not a library. You can modify files and compile project by `go build` in root.

### Disclaimer

"Master" branch is for releases only. To try bleeding edge versions, check "development" branch.

### Roadmap

RAWIG is small, personal project, and doesn't have proper, detailed roadmap. However, TODO list is maintained.

Every revision before v 0.1 is potentially unstable. 

**Implemented already:**  
- support for QWERTY, QWERTZ, AZERTY and Dvorak keyboards  
- configurable keybindings  
- rendering system  
- maps, monsters, items  
- fov, pathfinding  
- melee combat  
- menus  
- inventory system  
- ranged combat  
- looking command  
- save / load system  
- json for data storage  

**TODO:**  
- level generation algorithms

### Influences

RAWIG is built from scratch. Yet, due historical reasons, it is influenced by [roguebasin python+libtcod tutorial](http://www.roguebasin.com/index.php?title=Complete_Roguelike_Tutorial,_using_python%2Blibtcod). It's fair to give Jotaf, original author of this tutorial, a credit here.

### Contributors

[Adam](https://github.com/adam-weiler)  
[Travis Yoder](https://github.com/trayo)
