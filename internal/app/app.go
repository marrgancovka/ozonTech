package app

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"ozonTech/graph"
	"ozonTech/internal/pkg/auth"
	repoAuthInmemory "ozonTech/internal/pkg/auth/repo/in-memory"
	repoAuthPostgres "ozonTech/internal/pkg/auth/repo/postgres"
	usecaseAuth "ozonTech/internal/pkg/auth/usecase"
	"ozonTech/internal/pkg/comment"
	repoCommentInmemory "ozonTech/internal/pkg/comment/repo/in_memory"
	repoCommentPostgres "ozonTech/internal/pkg/comment/repo/postgres"
	usecaseComment "ozonTech/internal/pkg/comment/usecase"
	"ozonTech/internal/pkg/middleware"
	"ozonTech/internal/pkg/post"
	repoPostInmemory "ozonTech/internal/pkg/post/repo/in_memory"
	repoPostPostgres "ozonTech/internal/pkg/post/repo/postgres"
	usecasePost "ozonTech/internal/pkg/post/usecase"
)

type App struct {
	logger *logrus.Logger
}

func NewApp(logger *logrus.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Run() error {
	_ = godotenv.Load()

	var storageType string
	flag.StringVar(&storageType, "storage", "in-memory", "storage type (in-memory or postgres)")
	flag.Parse()

	var postRepo post.PostRepository
	var commentRepo comment.CommentRepository
	var authRepo auth.AuthRepository

	switch storageType {
	case "in-memory":
		a.logger.Info("using in-memory storage")
		postRepo = repoPostInmemory.NewInMemoryPostRepo()
		commentRepo = repoCommentInmemory.NewInMemoryCommentRepo()
		authRepo = repoAuthInmemory.NewInMemoryAuthRepo()
	case "postgres":
		db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")))
		if err != nil {
			a.logger.Error("failed to connect database " + err.Error())
			break
		}

		postRepo = repoPostPostgres.NewPostRepository(db)
		commentRepo = repoCommentPostgres.NewCommentRepository(db)
		authRepo = repoAuthPostgres.NewAuthRepository(db)

		if err = db.Ping(); err != nil {
			a.logger.Error("failed to ping database " + err.Error())
			return err
		}
		defer db.Close()
		a.logger.Info("successfully connected to database")
	default:
		log.Fatalf("Unknown storage type: %s", storageType)
	}

	postUsecase := usecasePost.NewPostUsecase(postRepo, commentRepo)
	commentUsecase := usecaseComment.NewCommentUsecase(commentRepo)
	authUsecase := usecaseAuth.NewAuthUsecase(authRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	resolver := &graph.Resolver{
		PostUsecase:    postUsecase,
		CommentUsecase: commentUsecase,
		AuthUsecase:    authUsecase,
	}
	c := graph.Config{Resolvers: resolver}
	c.Directives.Auth = middleware.AuthDirective

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleware(srv, a.logger))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
