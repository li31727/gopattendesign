package _interface
type Observer interface {
	Update(temp,humidity,pressure float32)
}
