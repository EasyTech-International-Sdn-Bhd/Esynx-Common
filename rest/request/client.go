package request

type ClientForm struct{}

type ClientRegisterForm struct {
	Username        string `form:"username" json:"username" binding:"required,v-naming"`
	Password        string `form:"password" json:"password" binding:"required,min=3,max=50"`
	ClientCompany   string `form:"clientCompany" json:"clientCompany" binding:"required,v-naming"`
	Metadata        string `form:"metadata" json:"metadata" binding:"required,v-metadata"`
	Server          string `form:"server" json:"server" binding:"required,v-addr"`
	BiDealer        string `form:"biDealer" json:"biDealer" binding:"-"`
	BiSubscriptions string `form:"biSubscriptions" json:"biSubscriptions" binding:"-"`
	BiState         string `form:"biState" json:"biState" binding:"-"`
	BiIndustry      string `form:"biIndustry" json:"biIndustry" binding:"-"`
}

type ClientLoginForm struct {
	Username string `form:"username" json:"username" binding:"required,v-naming"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=50"`
}

type ClientRefreshTokenForm struct {
	RefreshToken string `form:"refreshToken" json:"refreshToken" binding:"required"`
}

type ClientAccessTokenForm struct {
	AccessToken string `form:"accessToken" json:"accessToken" binding:"required"`
}
