package services

import (
	"context"

	"github.com/google/go-github/github"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

// IRepositoryServices : This interface will be used to provide light coupling between github.RepositoryServices and its consumer
//
type IRepositoryServices interface {
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
}

type IGithub interface {
	GetPackageRepoInfo(owner string, repositoryName string) (Package, error)
}

// Github : This struct will be used to get Github related information
type Github struct {
	repositoryServices IRepositoryServices
	context            context.Context
}

// NewGithub : It will intialized Github class
func NewGithub(context context.Context, repositoryServices *github.RepositoriesService) Github {
	return Github{
		repositoryServices: repositoryServices,
		context:            context,
	}
}

// GetPackageRepoInfo : This receiver provide Github related repository information
func (service *Github) GetPackageRepoInfo(owner string, repositoryName string) (*Package, error) {
	repo, _, err := service.repositoryServices.Get(service.context, owner, repositoryName)
	if err != nil {
		return nil, err
	}
	pack := &Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}
	return pack, nil
}
