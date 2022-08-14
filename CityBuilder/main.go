package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	img.Init(img.INIT_JPG | img.INIT_PNG)

	window, err := sdl.CreateWindow("City Builder", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1920, 1080, sdl.WINDOW_RESIZABLE)

	if err != nil {
		log.Fatal(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		log.Fatal(err)
	}

	var exit = false

	tiles := CreateTiles(1920/100, 1080/100, renderer, window)

	for !exit {
		// Event Handling
		var event = sdl.PollEvent()
		if event != nil {
			var eventType = event.GetType()

			switch eventType {
			case sdl.QUIT:
				exit = true
			case sdl.MOUSEBUTTONDOWN:
				x, y, state := sdl.GetMouseState()
				var mousePos = sdl.Point{x, y}
				var tileNum int
				var tile *Tile
				for i, _ := range tiles {
					if mousePos.InRect(tiles[i].rect) {
						tileNum = i
						break
					}
				}
				tile = &tiles[tileNum]
				if state == sdl.ButtonLMask() {
					tile.texture, err = img.LoadTexture(renderer, "assets/white_square.png")
				} else if state == sdl.ButtonRMask() {
					tile.texture, err = img.LoadTexture(renderer, "assets/black_square.png")
				}

				if err != nil {
					log.Fatal(err)
				}
			}

		}
		// Rendering
		renderer.Clear()
		for _, t := range tiles {
			renderer.Copy(t.texture, nil, t.rect)
		}
		renderer.Present()

	}

}

type Tile struct {
	rect    *sdl.Rect
	texture *sdl.Texture
}

func CreateTiles(xRes, yRes int32, renderer *sdl.Renderer, window *sdl.Window) []Tile {
	var tiles []Tile
	displayMode, err := window.GetDisplayMode()
	if err != nil {
		log.Fatal(err)
	}
	w := displayMode.W
	h := displayMode.H
	var xSize, ySize int32
	xSize = w / xRes
	ySize = h / yRes
	tempTex, _ := img.LoadTexture(renderer, "assets/black_sqare.png")
	for x := 0; x < int(w); x += int(xSize) {
		for y := 0; y < int(h); y += int(ySize) {
			tiles = append(tiles, Tile{&sdl.Rect{int32(x), int32(y), xSize, ySize}, tempTex})
		}
	}
	return tiles
}
