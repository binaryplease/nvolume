[![Build Status](https://travis-ci.org/binaryplease/nvolume.svg?branch=master)](https://travis-ci.org/binaryplease/nvolume)
# nvolume
--
    import "gitlab.com/binaryplease/NVolume"


## Usage

#### type NElement

```go
type NElement interface {
}
```


#### type NVolume

```go
type NVolume struct {
	Dimensions int
	Shape      []int
}
```

NVolume is a type for multi-dimensional volumes

#### func  New

```go
func New(size ...int) NVolume
```
New creates a new n-dimensional NVolume with given sizes. The number of
dimensions is implied with the number of given sizes

#### func (NVolume) Get

```go
func (nvol NVolume) Get(dim ...int) NElement
```
Get the Value at a given point in the NVolume

#### func (*NVolume) Set

```go
func (nvol *NVolume) Set(val NElement, dim ...int)
```
Set the Value at a given point in the NVolume
