package Interface
type Observer interface {
	Update(temp,humidity,pressure float32)
}
