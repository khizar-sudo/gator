package commands

import "context"

func handlerReset(s *state, cmd command) error {
	return s.db.DeleteAllUsers(context.Background())
}
