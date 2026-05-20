# todoist-api-go

> [!WARNING]
> This project has been archived since I no longer use Todoist.

[![Go Reference](https://pkg.go.dev/badge/github.com/CnTeng/todoist-api-go.svg)](https://pkg.go.dev/github.com/CnTeng/todoist-api-go)

This repository provides a Go library that implements Todoist API v1.

## 📦 Packages

- Package [todoist](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/todoist) implements the base client.

- Package [sync](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/sync) defines models and helpers for the Sync API.

- Package [rest](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/rest) defines models for the RESTful API.

- Package [ws](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/ws) implements a WebSocket client for real-time notifications.

## 🚧 Implementation Notes

Not all APIs have been implemented.
Most of the Sync APIs and a small portion of the RESTful APIs have been implemented.

### Feature Status

#### Sync API

| Feature            | Status         |
| ------------------ | -------------- |
| Sections           | ✅ Implemented |
| Reminders          | ✅ Implemented |
| Projects           | ✅ Implemented |
| Labels             | ✅ Implemented |
| Tasks              | ✅ Implemented |
| Filters            | ✅ Implemented |
| User               | 🔄 In progress |
| Workspace          | ❌ No plan     |
| Workspace users    | ❌ No plan     |
| View Options       | ❌ No plan     |
| Sharing            | ❌ No plan     |
| Comments           | ❌ No plan     |
| Live Notifications | ❌ No plan     |

#### RESTful API

| Feature                            | Status         |
| ---------------------------------- | -------------- |
| Projects Get Archived              | ✅ Implemented |
| Tasks Completed By Completion Date | ✅ Implemented |
| Tasks Completed By Due Date        | ✅ Implemented |

#### WebSocket API

| Feature          | Status         |
| ---------------- | -------------- |
| WebSocket Client | ✅ Implemented |

## 🧪 Examples

### Sync API Basic Usage

```go
token := os.Getenv("API_TOKEN")
cli := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)
taskSvc := todoist.NewTaskService(cli)

_, err := cli.SyncWithAutoToken(context.Background(), true)
if err != nil {
    panic(err)
}

args := &sync.TaskAddArgs{Content: "test"}
_, err = taskSvc.AddTask(context.Background(), args)
if err != nil {
    panic(err)
}
```

### Sync API Multiple Commands

```go
token := os.Getenv("API_TOKEN")
cli := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)

projCmd := sync.NewCommand(&sync.ProjectAddArgs{Name: "test"})

tempID := projCmd.TempID.String()
taskCmd := sync.NewCommand(&sync.TaskAddArgs{Content: "test", ProjectID: &tempID})

_, err := cli.ExecuteCommands(context.Background(), sync.Commands{projCmd, taskCmd})
if err != nil {
    panic(err)
}
```

### RESTful API Basic Usage

```go
token := os.Getenv("API_TOKEN")
cli := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)
taskSvc := todoist.NewTaskService(cli)

req := &rest.TaskQuickAddRequest{Text: "test"}
_, err := taskSvc.QuickAddTask(context.Background(), req)
if err != nil {
    panic(err)
}
```

### Custom Handler

```go
func (h *customHandler) SyncToken(ctx context.Context) (*string, error) {
	// Full sync every time.
	st := sync.DefaultSyncToken
	return &st, nil
}

func (h *customHandler) ResourceTypes(ctx context.Context) (*sync.ResourceTypes, error) {
	// Only sync tasks.
	return &sync.ResourceTypes{sync.Tasks}, nil
}

func (h *customHandler) HandleResponse(ctx context.Context, resp any) error {
	// Print the tasks.
	switch r := resp.(type) {
	case *sync.SyncResponse:
		for _, task := range r.Tasks {
			fmt.Printf("Task ID: %s, Content: %s\n", task.ID, task.Content)
		}
	case *rest.TaskGetCompletedResponse:
		for _, task := range r.Tasks {
			fmt.Printf("Completed Task ID: %s, Content: %s\n", task.ID, task.Content)
		}
	}

	return nil
}
```

## 📚 More Examples

You can find more usage examples in [todoist-cli](https://github.com/CnTeng/todoist-cli),
a Todoist CLI client with an integrated WebSocket daemon for real-time syncing.

## 🔍 Comparison

| Feature                              | todoist-api-go | [todoist/cli](https://pkg.go.dev/github.com/sachaos/todoist/lib) |
| ------------------------------------ | -------------- | ---------------------------------------------------------------- |
| Exec command and sync in one request | ✅             | ❌                                                               |
| WebSocket support                    | ✅             | ❌                                                               |
| Local filter implementation          | ❌             | ✅                                                               |
| API version                          | Todoist API v1 | Sync API v9                                                      |

## ⭐ Credit

- [todoist](https://github.com/sachaos/todoist)
- [todoist-net](https://github.com/olsh/todoist-net)
