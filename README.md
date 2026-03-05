# Castle Defender

A turn-based tower defense game written in Go for the console. Defend your base by strategically placing towers to stop waves of enemies from reaching the end of the path!

## Game Features

- **Turn-based gameplay** - Make strategic decisions each turn
- **Tower placement** - Position defensive towers to protect the path
- **Automatic enemy spawning** - Enemies appear every 2 turns
- **Strategic targeting** - Towers prioritize enemies closest to the end
- **Resource management** - Earn gold by defeating enemies, spend it on towers
- **Health system** - Enemies require multiple hits to defeat

## Quick Start

### Prerequisites
- Go 1.25.0 or later

### Installation & Running

1. Clone or download the project
2. Navigate to the project directory
3. Run the game:
   ```bash
   go run .
   ```

## How to Play

### Objective
Prevent enemies from reaching the end of the path (`*`). You lose a life each time an enemy reaches the end. The game ends when you run out of lives.

### Controls
- **`n`** or **Enter** - Advance to next turn (enemies move, towers attack)
- **`p`** - Place a tower (costs 100 gold)
- **`q`** - Quit the game

### Game Mechanics

#### Tower Placement
- Towers cost **100 gold** each
- Cannot be placed on the path (marked with `1`) or end point (`*`)
- Cannot be placed on top of other towers
- Use coordinates **1-10 for rows, 1-15 for columns**
- Example: `5,8` places a tower at row 5, column 8

#### Tower Stats
- **Range**: 3 tiles (Manhattan distance)
- **Damage**: 1 per attack
- **Targeting**: Each tower targets the farthest enemy in range
- **Attack Rate**: 1 attack per turn

#### Enemy Stats
- **Health**: 3 HP (requires 3 tower hits to defeat)
- **Movement**: 1 tile per turn along the path
- **Reward**: 50 gold per defeated enemy

#### Enemy Spawning
- **Frequency**: Every 2 turns
- **Wave Size**: 1 enemy per wave
- **Starting Position**: Beginning of the path

### Game Symbols
- `E` - Enemy
- `T` - Tower
- `1` - Path
- `*` - End point
- `0` - Empty space

### Track Format
The game reads track layouts from `tracks/track.txt`:
- `0`: Empty space (valid for tower placement)
- `1`: Path tiles (enemies follow this route)
- `*`: End point (enemies reaching here cost a life)

Example track:
```
100000000000000
111111111110000
000000000010000
000000000010000
000111111110000
000100000000000
000100000000000
000111111111100
000000000000100
000000000000*00
```

## Project Structure

```
├── main.go          # Main game loop and logic
├── enemy.go         # Enemy struct and behavior
├── tower.go         # Tower struct and combat logic
├── entity.go        # Base entity struct
├── track.go         # Track parsing and pathfinding
├── tracks/
│   └── track.txt    # Default track layout
├── go.mod           # Go module file
└── README.md        # This file
```