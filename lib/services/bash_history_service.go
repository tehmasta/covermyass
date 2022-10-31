package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sundowndev/covermyass/v2/lib/find"
	"os"
)

type ShellHistoryService struct{}

func NewShellHistoryService() Service {
	return &ShellHistoryService{}
}

func (s *ShellHistoryService) Name() string {
	return "shell_history"
}

func (s *ShellHistoryService) Paths() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Error(err)
		return []string{}
	}
	return []string{
		fmt.Sprintf("%s/.bash_history", homeDir),
		fmt.Sprintf("%s/.zsh_history", homeDir),
		fmt.Sprintf("%s/.node_repl_history", homeDir),
		fmt.Sprintf("%s/.python_history", homeDir),
	}
}

func (s *ShellHistoryService) HandleFile(file find.FileInfo) error {
	return os.Truncate(file.Path(), 0)
}
