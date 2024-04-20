package facades

import "github.com/Firhan384/framework/contracts/mail"

func Mail() mail.Mail {
	return App().MakeMail()
}
