# todoist-api-go

[![Go][go-shield]][go-url]

This repository provides a Go library that implements Todoist API v1.

## Packages

- Package [todoist](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/todoist) implements the base client.

- Package [sync](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/sync) implements the Sync API models.

- Package [rest](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/rest) implements the RESTful API models.

- Package [ws](https://pkg.go.dev/github.com/CnTeng/todoist-api-go/ws) implements the websocket client.

## Implementation Notes

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

[go-shield]: https://img.shields.io/github/go-mod/go-version/CnTeng/todoist-api-go?style=for-the-badge&logo=go
[go-url]: https://golang.org
