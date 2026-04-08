package app

import env "github.com/atharvyadav96k/SPOTNEARR_SHARED/app/Env"

func Init() *App {
	return &App{
		Env: env.Init(),
	}
}
