//go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"                              //nolint
	_ "github.com/cloudspannerecosystem/wrench"                      //nolint
	_ "github.com/cosmtrek/air"                                      //nolint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"          //nolint
	_ "github.com/kauche/splanter"                                   //nolint
	_ "github.com/spf13/cobra-cli"                                   //nolint
	_ "github.com/volatiletech/sqlboiler/v4"                         //nolint
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql" //nolint
	_ "go.mercari.io/yo"                                             //nolint
	_ "go.uber.org/mock/mockgen"                                     //nolint
	_ "golang.org/x/tools/cmd/goimports"                             //nolint
)
