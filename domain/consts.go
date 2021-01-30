package domain

// Violations
const (
	Account                    string = "account"
	Transaction                string = "transaction"
	AccountAlreadyInitialized  string = "account-already-initialized"
	AccountNotInitialized      string = "account-not-initialized"
	DoubleTransaction          string = "double-transaction"
	InsufficientLimit          string = "insufficient-limit"
	CardNotActive              string = "card-not-active"
	HighFrequencySmallInterval string = "high-frequency-small-interval"
	ContainViolations          string = "contain-violations"
)
