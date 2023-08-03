package govatar

import (
	"github.com/disintegration/imaging"
	"golang.org/x/crypto/scrypt"

	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

var (
	blockX    int
	blockY    int
	rows      int
	columns   int
	numBlocks int
	keyLength int
)

type block struct{}

func CreateAvatarWithDefault(username string) (image.Image, error) {
	return CreateAvatar(200, 200, 20, 20, 20, username, "tiktok")
}

// Create the avatar and return as image.Image
func CreateAvatar(canvasWidth, canvasHeight, blockWidth, blockHeight int, vibrance uint8, userName, salt string) (image.Image, error) {
	var shuffleInt int64
	shuffleInt = 1

	// number of rows/columns
	rows = (canvasWidth / 2) / blockWidth
	columns = canvasHeight / blockHeight

	// number of blocks
	numBlocks = rows * columns

	keyLength = numBlocks

	var (
		hash   []uint8
		colors []uint8
		err    error
	)

	// Generate hash using Scrypt
	hash, err = Hash(userName, salt)

	if err != nil {
		return nil, err
	}

	// Create the avatar image
	r := image.Rect(0, 0, canvasWidth, canvasHeight)
	img := image.NewRGBA(r)

	// Create and draw a grey block as background color
	grayBlock := color.RGBA{R: 210, G: 210, B: 210, A: 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{C: grayBlock}, image.ZP, draw.Src)

	// Create 32 blocks to fill half of avatar canvas
	blocks := make([]block, numBlocks)

	// For each block decide if to draw anything and what color to draw based on hash
	for i := range blocks {

		var gC uint8
		var rC uint8
		var bC uint8
		var aC uint8

		var skip bool

		num0 := hash[i] % 10
		num1 := (hash[i] / 10) % 10
		num2 := (hash[i] / 100) % 10

		if i == 0 {
			colors = make([]uint8, 3)
		}

		if num2 == 0 {
			// If byte is < 100 then display nothing
			skip = true

			// Use this data to seed the RBG shuffle rand value
			if num1*num0 != 0 {
				shuffleInt *= int64(num1 * num0)
			} else {
				shuffleInt += int64(num1 + num0)
			}

		} else {
			skip = false
			aC = 255

			colors = []uint8{
				rC,
				gC,
				bC,
			}

			// Generate RGB values
			colors[0] = ((num0 * vibrance) * num1) + num2
			colors[1] = ((num1 * vibrance) * num2) + num0
			colors[2] = ((num2 * vibrance) * num0) + num1

			// Shuffle the colors based on rand
			Shuffle(colors, shuffleInt)
		}

		// Start position of block
		if i == 0 {
			blockX = 0
			blockY = 0
		} else {
			// New row
			if i%rows == 0 {
				blockX = 0
				blockY += blockHeight
				// New column
			} else {
				blockX += blockWidth
			}
		}

		if skip {
			continue
		}

		// If colors have somehow overstepped bounds reset them
		for i := range colors {

			if colors[i] > 255 {
				colors[i] = 255
			}

			if colors[i] < 0 {
				colors[i] = 0
			}
		}

		// Create a color based on hash in random order
		colorBlock := color.RGBA{R: colors[0], G: colors[1], B: colors[2], A: aC}

		// Print area is size of one block starting from position of block
		spMin := image.Point{X: blockX, Y: blockY}
		spMax := image.Point{X: blockX + blockWidth, Y: blockY + blockHeight}

		// Create a rectangle the size of the block
		blockRectangle := image.Rectangle{Min: spMin, Max: spMax}

		// Draw the block to the image
		draw.Draw(img, blockRectangle, &image.Uniform{C: colorBlock}, image.Point{X: 0, Y: 0}, draw.Src)
	}

	// Create a horizontal mirror of the generated image
	imgHFlip := imaging.FlipH(img)

	// Draw to the right side of the canvas
	r2 := image.Rect(canvasWidth, canvasHeight, canvasWidth/2, 0)
	draw.Draw(img, r2, imgHFlip, image.Point{X: canvasWidth / 2, Y: 0}, draw.Src)

	return img, nil
}

// Generates a hashed byte array based on username and salt
func Hash(userName, salt string) ([]byte, error) {
	var (
		hash []byte
		err  error
	)

	hash, err = scrypt.Key([]byte(userName), []byte(salt), 16384, 8, 1, keyLength)

	if err != nil {
		return nil, err
	}

	return hash, err
}

// Shuffles a uint8 array
func Shuffle(vals []uint8, colorInt int64) {
	r := rand.New(rand.NewSource(colorInt))

	r.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})
}
