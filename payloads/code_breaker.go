package payloads

type CodeBreakerCreatePayload struct {
	Columns int
	Colors  []string
}

type CodeBreakerGuessPayload struct {
	Colors []string
}
