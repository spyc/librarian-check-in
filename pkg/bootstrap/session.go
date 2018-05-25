package bootstrap

import (
	"crypto/rand"
	"github.com/facebookgo/inject"
	"github.com/gorilla/sessions"
)

func injectSessionStore(graph *inject.Graph) error {
	keyPair := make([]byte, 16)
	rand.Read(keyPair)

	return graph.Provide(
		&inject.Object{Value: sessions.NewCookieStore(keyPair)},
	)
}
