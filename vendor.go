package vendors

import (
	"github.com/lastbackend/vendors/vcs/github"
	"github.com/lastbackend/vendors/vcs/bitbucket"
	"github.com/lastbackend/vendors/vcs/gitlab"
	"github.com/lastbackend/vendors/notify/slack"
	"github.com/lastbackend/vendors/notify/wechat"
)

func GetGitHub(clientID, clientSecretID, redirectURI string) *github.GitHub {
	return github.GetClient(clientID, clientSecretID, redirectURI)
}

func GetBitBucket(clientID, clientSecretID, redirectURI string) *bitbucket.BitBucket {
	return bitbucket.GetClient(clientID, clientSecretID, redirectURI)
}

func GetGitLab(clientID, clientSecretID, redirectURI string) *gitlab.GitLab {
	return gitlab.GetClient(clientID, clientSecretID, redirectURI)
}

func GetSlack(clientID, clientSecretID, redirectURI string) *slack.Slack {
	return slack.GetClient(clientID, clientSecretID, redirectURI)
}

func GetWeChat(clientID, clientSecretID, redirectURI string) *wechat.WeChat {
	return wechat.GetClient(clientID, clientSecretID, redirectURI)
}