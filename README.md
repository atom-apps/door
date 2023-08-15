// send code check csrf or clientID





prompt:








type SessionItem struct {
	ID        int64     `json:"id,omitempty"`         //
	CreatedAt time.Time `json:"created_at,omitempty"` //
	UpdatedAt time.Time `json:"updated_at,omitempty"` //
	UserID    int64     `json:"user_id,omitempty"`    //
	SessionID string    `json:"session_id,omitempty"` //
}


// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"SessionID"
//	@Success		200	{object}	dto.SessionItem
//	@Router			/sessions/{id} [get]
func (c *SessionController) Show(ctx *fiber.Ctx, id int64) (*dto.SessionItem, error)