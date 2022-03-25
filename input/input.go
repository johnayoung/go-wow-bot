package input

type Input struct {
	KeyPresses  chan<- string
	AgentStatus chan<- string
}
