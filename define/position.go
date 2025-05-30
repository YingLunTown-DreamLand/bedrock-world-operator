package define

import (
	"fmt"
)

// ChunkPos holds the position of a chunk. The type is provided as a utility struct for keeping track of a
// chunk's position. Chunks do not themselves keep track of that. Chunk positions are different from block
// positions in the way that increasing the X/Z by one means increasing the absolute value on the X/Z axis in
// terms of blocks by 16.
type ChunkPos [2]int32

// String implements fmt.Stringer and returns (x, z).
func (p ChunkPos) String() string {
	return fmt.Sprintf("(%v, %v)", p[0], p[1])
}

// X returns the X coordinate of the chunk position.
func (p ChunkPos) X() int32 {
	return p[0]
}

// Z returns the Z coordinate of the chunk position.
func (p ChunkPos) Z() int32 {
	return p[1]
}

// SubChunkPos holds the position of a sub-chunk. The type is provided as a utility struct for keeping track of a
// sub-chunk's position. Sub-chunks do not themselves keep track of that. Sub-chunk positions are different from
// block positions in the way that increasing the X/Y/Z by one means increasing the absolute value on the X/Y/Z axis in
// terms of blocks by 16.
type SubChunkPos [3]int32

// String implements fmt.Stringer and returns (x, y, z).
func (p SubChunkPos) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p[0], p[1], p[2])
}

// X returns the X coordinate of the sub-chunk position.
func (p SubChunkPos) X() int32 {
	return p[0]
}

// Y returns the Y coordinate of the sub-chunk position.
func (p SubChunkPos) Y() int32 {
	return p[1]
}

// Z returns the Z coordinate of the sub-chunk position.
func (p SubChunkPos) Z() int32 {
	return p[2]
}
