package day20

type Transition int

const (
	Rotate0Deg Transition = iota + 1
	Rotate90Deg
	Rotate180Deg
	Rotate270Deg
	FlipHorizontal
	FlipVertical
	EndOfTransition
)

type Tile struct {
	ID           int
	BorderTop    string
	BorderBottom string
	BorderLeft   string
	BorderRight  string
}

func (t *Tile) Clone() *Tile {
	tile := &Tile{}
	tile.ID = t.ID
	return tile
}

func (t *Tile) Transposition(transition Transition) *Tile {
	switch transition {
	case Rotate0Deg:
		return t.Rotate0Deg()
	case Rotate90Deg:
		return t.Rotate90Deg()
	case Rotate180Deg:
		return t.Rotate180Deg()
	case Rotate270Deg:
		return t.Rotate270Deg()
	case FlipHorizontal:
		return t.FlipHorizontal()
	case FlipVertical:
		return t.FlipVertical()
	default:
		return nil
	}
}

func (t *Tile) Rotate0Deg() *Tile {
	tile := t.Clone()
	tile.BorderTop = t.BorderTop
	tile.BorderBottom = t.BorderBottom
	tile.BorderLeft = t.BorderLeft
	tile.BorderRight = t.BorderRight
	return tile
}

func (t *Tile) Rotate90Deg() *Tile {
	tile := t.Clone()
	tile.BorderTop = t.BorderRight
	tile.BorderLeft = reverse(t.BorderTop)
	tile.BorderBottom = t.BorderLeft
	tile.BorderRight = reverse(t.BorderBottom)
	return tile
}

func (t *Tile) Rotate180Deg() *Tile {
	tile := t.Clone()
	tile.BorderTop = reverse(t.BorderBottom)
	tile.BorderLeft = reverse(t.BorderRight)
	tile.BorderBottom = reverse(t.BorderTop)
	tile.BorderRight = reverse(t.BorderLeft)
	return tile
}

func (t *Tile) Rotate270Deg() *Tile {
	tile := t.Clone()
	tile.BorderTop = t.BorderLeft
	tile.BorderLeft = reverse(t.BorderBottom)
	tile.BorderBottom = t.BorderRight
	tile.BorderRight = reverse(t.BorderTop)
	return tile
}

func (t *Tile) FlipVertical() *Tile {
	tile := t.Clone()
	tile.BorderTop = t.BorderBottom
	tile.BorderBottom = t.BorderTop
	tile.BorderLeft = reverse(t.BorderLeft)
	tile.BorderRight = reverse(t.BorderRight)
	return tile
}

func (t *Tile) FlipHorizontal() *Tile {
	tile := t.Clone()
	tile.BorderLeft = t.BorderRight
	tile.BorderRight = t.BorderLeft
	tile.BorderTop = reverse(t.BorderTop)
	tile.BorderBottom = reverse(t.BorderBottom)
	return tile
}

type Match struct {
	SourceID    int
	SourceState Transition
	DestID      int
	DestState   Transition
}

type UniqueMatch struct {
	SourceID int
	DestID   int
}

func reverse(a string) string {
	count := len(a)
	bytes := make([]byte, count)
	for i := 0; i < count; i++ {
		bytes[i] = a[count-1-i]
	}
	return string(bytes)
}
