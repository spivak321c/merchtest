package dto


type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}


 type ProfileResponse struct {
  	ID       uint     `json:"id"`
  	Email    string   `json:"email"`
  	Name     string   `json:"name"`
  	Country  string   `json:"country"`
  	Addresses []string `json:"addresses,omitempty"`
  }	


   type UserUpdateRequest struct {
  	Email    string   `json:"email,omitempty"`
  	Name     string   `json:"name,omitempty"`
  	Country  string   `json:"country,omitempty"`
  	Addresses []string `json:"addresses,omitempty"`
  }	
