package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/notnil/chess"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Chess")
	game := chess.NewGame()
	// we need to know the current state of the board
	grid := createGrid(game.Position().Board())
	// we need a container for floating pieces
	floatPiece := canvas.NewImageFromResource(nil)
	floatPiece.Hide()
	w.SetContent(container.NewMax(grid, container.NewWithoutLayout(floatPiece)))
	w.Resize(fyne.NewSize(480, 480))

	// get available moves
	go func() {
		rand.Seed(time.Now().Unix())
		// keep playing until the game not over
		for game.Outcome() == chess.NoOutcome {
			time.Sleep(1 * time.Second)
			validMoves := game.ValidMoves()
			// pick random move
			randMove := validMoves[rand.Intn(len(validMoves))]
			move(randMove, game, grid, floatPiece)
		}

	}()
	w.ShowAndRun()
}

func move(randMove *chess.Move, game *chess.Game, grid *fyne.Container, floatPiece *canvas.Image) {
	offset := squareToOffset(randMove.S1())
	// get the cell in the container
	cell := grid.Objects[offset].(*fyne.Container)
	// the image of the piece that moved
	startImg := cell.Objects[1].(*canvas.Image)
	startPos := cell.Position()

	floatPiece.Resource = startImg.Resource
	floatPiece.Move(startPos)
	// set the float piece to be on top of where the prev image was
	floatPiece.Resize(startImg.Size())

	floatPiece.Show()
	startImg.Resource = nil
	startImg.Refresh()

	offset = squareToOffset(randMove.S2())
	cell = grid.Objects[offset].(*fyne.Container)
	// the image of the piece that moved
	endPos := cell.Position()

	// make animation

	a := canvas.NewPositionAnimation(startPos, endPos, time.Millisecond*500, func(p fyne.Position) {
		floatPiece.Move(p)
		floatPiece.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 550)

	// make the move
	game.Move(randMove)
	// change the grid to reflect a new move
	refreshGrid(grid, game.Position().Board())
	floatPiece.Hide()

}

func squareToOffset(square chess.Square) int {
	x := square % 8
	// the board is upside down
	y := 7 - ((square - x) / 8)

	return int(x + y*8)
}
func refreshGrid(grid *fyne.Container, board *chess.Board) {
	y, x := 7, 0
	for _, cell := range grid.Objects {
		// we want to give the correct image to every cell
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		piece := board.Piece(chess.Square(x + y*8))

		img.Resource = resourceForPiece(piece)
		img.Refresh()
		x++
		if x == 8 {
			x = 0
			y--
		}
	}
}
func createGrid(board *chess.Board) *fyne.Container {
	grid := container.NewGridWithColumns(8)
	// color the board
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bg := canvas.NewRectangle(color.Gray{0x30})
			if x%2 == y%2 {
				bg.FillColor = color.Gray{0xE0}
			}
			// figure in whice squre we are
			piece := board.Piece(chess.Square(x + (7-y)*8))
			// add image to board
			img := canvas.NewImageFromResource(resourceForPiece(piece))
			// maintin the ratio
			img.FillMode = canvas.ImageFillContain
			grid.Add(container.NewMax(bg, img))

		}
	}
	return grid
}
