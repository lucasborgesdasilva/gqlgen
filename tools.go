//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)

// O _ serve para indicar que esses pacotes não serão necessários na build final,
// mas são necessários para garantir que as ferramentas estejam disponíveis
// durante o desenvolvimento e a geração de código.