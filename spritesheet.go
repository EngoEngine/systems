package common

import (
	"log"

	"engo.io/engo"
	"engo.io/engo/math"
	"engo.io/gl"
)

// Spritesheet is a class that stores a set of tiles from a file, used by tilemaps and animations
type Spritesheet struct {
	texture       *gl.Texture     // The original texture
	width, height float32         // The dimensions of the total texture
	cells         []Sprite        // The dimensions of each sprite
	cache         map[int]Texture // The cell cache cells
}

// Sprite holds the position data for each sprite on the sheet
type Sprite struct {
	Position      engo.Point
	Width, Height int
}

// NewAsymmetricSpriteSheetFromTexture creates a new AsymmetricSpriteSheet from a
// TextureResource. The data provided is the location and size of the sprites
func NewAsymmetricSpritesheetFromTexture(tr *TextureResource, sprites []Sprite) *Spritesheet {
	return &Spritesheet{
		texture: tr.Texture,
		width:   tr.Width,
		height:  tr.Height,
		cells:   sprites,
		cache:   make(map[int]Texture),
	}
}

// NewAsymmetricSpriteSheetFromFile creates a new AsymmetricSpriteSheet from a
// file name. The data provided is the location and size of the sprites
func NewAsymmetricSpritesheetFromFile(textureName string, sprites []Sprite) *Spritesheet {
	res, err := engo.Files.Resource(textureName)
	if err != nil {
		log.Println("[WARNING] [NewAsymmetricSpritesheetFromFile]: Received error:", err)
		return nil
	}

	img, ok := res.(TextureResource)
	if !ok {
		log.Println("[WARNING] [NewAsymmetricSpritesheetFromFile]: Resource not of type `TextureResource`:", textureName)
		return nil
	}

	return NewAsymmetricSpritesheetFromTexture(&img, sprites)
}

// NewSpritesheetFromTexture creates a new spritesheet from a texture resource.
func NewSpritesheetFromTexture(tr *TextureResource, cellWidth, cellHeight int) *Spritesheet {
	sprites := generateSprites(tr.Width, tr.Height, cellWidth, cellHeight, 0, 0)
	return NewAsymmetricSpritesheetFromTexture(tr, sprites)
}

// NewSpritesheetFromFile is a simple handler for creating a new spritesheet from a file
// textureName is the name of a texture already preloaded with engo.Files.Add
func NewSpritesheetFromFile(textureName string, cellWidth, cellHeight int) *Spritesheet {
	res, err := engo.Files.Resource(textureName)
	if err != nil {
		log.Println("[WARNING] [NewSpritesheetFromFile]: Received error:", err)
		return nil
	}

	img, ok := res.(TextureResource)
	if !ok {
		log.Println("[WARNING] [NewSpritesheetFromFile]: Resource not of type `TextureResource`:", textureName)
		return nil
	}

	return NewSpritesheetFromTexture(&img, cellWidth, cellHeight)
}

// NewSpritesheetWithBorderFromTexture creates a new spritesheet from a texture resource.
// This sheet has sprites of a uniform width and height, but also have borders around
// each sprite to prevent bleeding over
func NewSpritesheetWithBorderFromTexture(tr *TextureResource, cellWidth, cellHeight, borderWidth, borderHeight int) *Spritesheet {
	sprites := generateSprites(tr.Width, tr.Height, cellWidth, cellHeight, borderWidth, borderHeight)
	return NewAsymmetricSpritesheetFromTexture(tr, sprites)
}

// NewSpritesheetWithBorderFromFile creates a new spritesheet from a file
// This sheet has sprites of a uniform width and height, but also have borders around
// each sprite to prevent bleeding over
func NewSpritesheetWithBorderFromFile(textureName string, cellWidth, cellHeight, borderWidth, borderHeight int) *Spritesheet {
	res, err := engo.Files.Resource(textureName)
	if err != nil {
		log.Println("[WARNING] [NewSpritesheetWithBorderFromFile]: Received error:", err)
		return nil
	}

	img, ok := res.(TextureResource)
	if !ok {
		log.Println("[WARNING] [NewSpritesheetWithBorderFromFile]: Resource not of type `TextureResource`:", textureName)
		return nil
	}

	return NewSpritesheetWithBorderFromTexture(&img, cellWidth, cellHeight, borderWidth, borderHeight)
}

// Cell gets the region at the index i, updates and pulls from cache if need be
func (s *Spritesheet) Cell(index int) Texture {
	if r, ok := s.cache[index]; ok {
		return r
	}

	cell := s.cells[index]
	s.cache[index] = Texture{
		id:     s.texture,
		width:  float32(cell.Width),
		height: float32(cell.Height),
		viewport: engo.AABB{
			Min: engo.Point{
				X: cell.Position.X / s.width,
				Y: cell.Position.Y / s.height,
			},
			Max: engo.Point{
				X: (cell.Position.X + float32(cell.Width)) / s.width,
				Y: (cell.Position.Y + float32(cell.Height)) / s.height,
			},
		},
	}

	return s.cache[index]
}

// Drawable returns the drawable for a given index
func (s *Spritesheet) Drawable(index int) Drawable {
	return s.Cell(index)
}

// Drawables returns all the drawables on the sheet
func (s *Spritesheet) Drawables() []Drawable {
	drawables := make([]Drawable, s.CellCount())

	for i := 0; i < s.CellCount(); i++ {
		drawables[i] = s.Drawable(i)
	}

	return drawables
}

// CellCount returns the number of cells on the sheet
func (s *Spritesheet) CellCount() int {
	return len(s.cells)
}

// Cells returns all the cells on the sheet
func (s *Spritesheet) Cells() []Texture {
	cellsNo := s.CellCount()
	cells := make([]Texture, cellsNo)
	for i := 0; i < cellsNo; i++ {
		cells[i] = s.Cell(i)
	}

	return cells
}

// Width is the amount of tiles on the x-axis of the spritesheet
// only if the sprite sheet is symmetric with no border.
func (s Spritesheet) Width() float32 {
	return s.width / s.Cell(0).Width()
}

// Height is the amount of tiles on the y-axis of the spritesheet
// only if the sprite sheet is symmetric with no border.
func (s Spritesheet) Height() float32 {
	return s.height / s.Cell(0).Height()
}

func generateSprites(totalWidth, totalHeight float32, cellWidth, cellHeight, borderWidth, borderHeight int) []Sprite {
	var sprites []Sprite

	for y := 0; y <= int(math.Floor(totalHeight-1)); y += cellHeight + borderHeight {
		for x := 0; x <= int(math.Floor(totalWidth-1)); x += cellWidth + borderWidth {
			sprite := Sprite{
				Position: engo.Point{float32(x), float32(y)},
				Width:    cellWidth,
				Height:   cellHeight,
			}
			sprites = append(sprites, sprite)
		}
	}

	return sprites
}

/*
type Sprite struct {
	Position *Point
	Scale    *Point
	Anchor   *Point
	Rotation float32
	Color    color.Color
	Alpha    float32
	Region   *Region
}

func NewSprite(region *Region, x, y float32) *Sprite {
	return &Sprite{
		Position: &Point{x, y},
		Scale:    &Point{1, 1},
		Anchor:   &Point{0, 0},
		Rotation: 0,
		Color:    color.White,
		Alpha:    1,
		Region:   region,
	}
}
*/
