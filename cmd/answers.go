package cmd

import (
	"context"

	glean "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"github.com/gleanwork/glean-cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdAnswers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "answers",
		Short: "Manage Glean answers",
		Long: `Manage Glean answers.

Answers are curated Q&A pairs that surface authoritative responses in search results.

Example:
  glean answers list
  glean answers get --json '{"answerId":123}'`,
	}
	cmd.AddCommand(
		newAnswersListCmd(),
		newAnswersGetCmd(),
		newAnswersCreateCmd(),
		newAnswersUpdateCmd(),
		newAnswersDeleteCmd(),
	)
	return cmd
}

func newAnswersListCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.ListAnswersRequest]{
		Use:   "list",
		Short: "List answers (deprecated: answer boards are being removed by Glean in October 2026)",
		Run: func(ctx context.Context, sdk *glean.Glean, req components.ListAnswersRequest) (any, error) {
			resp, err := sdk.Client.Answers.List(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.ListAnswersResponse, nil
		},
	})
}

func newAnswersGetCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.GetAnswerRequest]{
		Use:          "get",
		Short:        "Get an answer",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.GetAnswerRequest) (any, error) {
			resp, err := sdk.Client.Answers.Retrieve(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.GetAnswerResponse, nil
		},
	})
}

func newAnswersCreateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.CreateAnswerRequest]{
		Use:          "create",
		Short:        "Create an answer",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.CreateAnswerRequest) (any, error) {
			resp, err := sdk.Client.Answers.Create(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Answer, nil
		},
	})
}

func newAnswersUpdateCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.EditAnswerRequest]{
		Use:          "update",
		Short:        "Update an answer",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.EditAnswerRequest) (any, error) {
			resp, err := sdk.Client.Answers.Update(ctx, req, nil)
			if err != nil {
				return nil, err
			}
			return resp.Answer, nil
		},
	})
}

func newAnswersDeleteCmd() *cobra.Command {
	return cmdutil.Build(cmdutil.Spec[components.DeleteAnswerRequest]{
		Use:          "delete",
		Short:        "Delete an answer",
		JSONRequired: true,
		Run: func(ctx context.Context, sdk *glean.Glean, req components.DeleteAnswerRequest) (any, error) {
			_, err := sdk.Client.Answers.Delete(ctx, req, nil)
			return nil, err
		},
	})
}
