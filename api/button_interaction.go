package api

// ButtonInteraction is a specific Interaction when CLicked on Button(s)
type ButtonInteraction struct {
	*ComponentInteraction
	Data *ButtonInteractionData `json:"data,omitempty"`
}

// ButtonInteractionData is the Button data payload
type ButtonInteractionData struct {
	*ComponentInteractionData
}
