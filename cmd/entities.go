package cmd

import (
	"context"
	"fmt"
	"strings"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

const entityTypeValidValues = "PEOPLE, TEAMS, CUSTOM_ENTITIES"

func NewCmdEntities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entities",
		Short: "List and read Glean entities and people",
		Long: `List and read Glean entities and people.

Entities represent structured objects in your company's knowledge graph — people, teams, projects, and more.

Example:
  glean entities list --json '{"entityType":"PEOPLE","query":"engineering"}'
  glean entities read-people --json '{"emailIds":["user@example.com"]}'`,
	}
	cmd.AddCommand(
		newEntitiesListCmd(),
		newEntitiesReadPeopleCmd(),
	)
	return cmd
}

func newEntitiesListCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.ListEntitiesRequest]{
		Use:          "list",
		Short:        "List entities",
		JSONRequired: true,
		ErrTransform: func(err error) error {
			if strings.Contains(err.Error(), "ListEntitiesRequestEntityType") {
				return fmt.Errorf("invalid entityType: valid values are %s", entityTypeValidValues)
			}
			return fmt.Errorf("invalid --json: %w", err)
		},
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ListEntitiesRequest) (any, error) {
			resp, err := sdk.Client.Entities.List(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.ListEntitiesResponse, nil
		},
	})
}

func newEntitiesReadPeopleCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.PeopleRequest]{
		Use:          "read-people",
		Short:        "Read people records",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.PeopleRequest) (any, error) {
			resp, err := sdk.Client.Entities.ReadPeople(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.PeopleResponse, nil
		},
	})
}
