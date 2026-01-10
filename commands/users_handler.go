package commands

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, val := range users {
		if val.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", val.Name)
		} else {
			fmt.Printf("* %s\n", val.Name)
		}
	}
	return nil
}
