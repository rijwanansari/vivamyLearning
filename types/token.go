package types

type (
	Token struct {
		UserID        int    `json:"uid"`
		AccessToken   string `json:"act"`
		RefreshToken  string `json:"rft"`
		AccessUuid    string `json:"aid"`
		RefreshUuid   string `json:"rid"`
		AccessExpiry  int64  `json:"axp"`
		RefreshExpiry int64  `json:"rxp"`
	}
)
