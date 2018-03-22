package issues

import (
	"fmt"

	"github.com/jenkins-x/jx/pkg/auth"
	"github.com/jenkins-x/jx/pkg/gits"
)

type IssueProvider interface {
	GetIssue(key string) (*gits.GitIssue, error)

	SearchIssues(query string) ([]*gits.GitIssue, error)

	CreateIssue(issue *gits.GitIssue) (*gits.GitIssue, error)

	CreateIssueComment(key string, comment string) error
}

func CreateIssueProvider(kind string, server *auth.AuthServer, project string) (IssueProvider, error) {
	switch kind {
	case Jira:
		return CreateJiraIssueProvider(server, project)
	default:
		return nil, fmt.Errorf("Unsupported issue provider kind: %s", kind)
	}
}
