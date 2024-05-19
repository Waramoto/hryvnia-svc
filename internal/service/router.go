package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/Waramoto/hryvnia-svc/internal/config"
	pg "github.com/Waramoto/hryvnia-svc/internal/data/postgres"
	"github.com/Waramoto/hryvnia-svc/internal/service/handlers"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(pg.NewDB(cfg.DB())),
		),
	)

	r.Route("/api", func(r chi.Router) {
		r.Get("/rate", handlers.GetRate)
		r.Post("/subscribe", handlers.Subscribe)
	})

	return r
}
