package main

import (
	"context"

	_ "github.com/benthosdev/benthos/v4/public/components/io"
	_ "github.com/benthosdev/benthos/v4/public/components/pure"
	"github.com/benthosdev/benthos/v4/public/service"
	_ "github.com/nibbleshift/mod2blob/test/bloblang"
)

func main() {
	service.RunCLI(context.Background())
}
