# Task-cli

Task-cli is a todo application available in the terminal. There are five main operations these are:

- Add, `task-cli add "task to todo"`
- Update, `task-cli update <id> "New description of task"`
- Delete, `task-cli delete <id>`
- List, `task-cli list`, this command can take additional argument for what status tasks to show, e.g. if only interested in non-started tasks run `task-cli list todo`. There are three total statuses `done`, `in-progress` and `todo`. Default behaviour is to show non-completed tasks.
- Mark, `task-cli mark 1 done`. This command changes the status of a commmand.
