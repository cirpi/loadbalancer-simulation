package primitives

type Response struct {
	Response string
	Receiver chan Response
}
