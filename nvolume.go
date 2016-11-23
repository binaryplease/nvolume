package nvolume

import ()

//NVolume is a type for multi-dimensional volumes
type NVolume struct {
	Dimensions int
	Shape      []int
	Elements   []NElement
}

//NElement empty interface to hold NVolume elements
type NElement interface{}

//New creates a new n-dimensional NVolume with given sizes.
//The number of dimensions is implied with the number of given sizes
func New(size ...int) NVolume {

	vol := NVolume{}

	numelem := 1
	vol.Dimensions = len(size)
	vol.Shape = []int{}
	for i := 0; i < vol.Dimensions; i++ {
		numelem *= size[i]
		vol.Shape = append(vol.Shape, size[i])
	}

	vol.Elements = make([]NElement, numelem, numelem)

	return vol
}

func (nvol NVolume) dimToPos(dims []int) int {

	pos := 0
	for i := 0; i < len(dims); i++ {
		pos += nvol.Shape[i] * dimsize(dims[i:])
	}

	//fmt.Println("dims: ", dims, " pos: ", pos)
	return pos
}

func dimsize(sizes []int) int {
	if len(sizes) == 1 {
		return sizes[0]
	}
	return (sizes[0]) * dimsize(sizes[1:])
}

//Set the Value at a given point in the NVolume
func (nvol *NVolume) Set(val NElement, dim ...int) {
	//fmt.Println(dim, dimT(dim))
	//nvol.Elements[nvol.dimToPos(dim)] = val
}

//Get the Value at a given point in the NVolume
func (nvol NVolume) Get(dim ...int) NElement {
	//return nvol.Elements[nvol.dimToPos(dim)]
	return 0
}
