package integration

import (
	"testing"
)

func TestRunner(t *testing.T) {
	// <setup code>
	t.Run("Account not initialized", AccountNotInitialized)
	t.Run("Initialize account", InitializeAccount)
	t.Run("Account already initialized", AccountAlreadyInitialized)
	t.Run("Insufficiente Limie", InsufficientLimit)
	t.Run("Process transaction", ProcessTransaction)
	t.Run("High FrequencySmall Interval", HighFrequencySmallInterval)
	t.Run("Card Not Active", CardNotActive)
	// <tear-down code>
}
