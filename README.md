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

- [ ] ~~Workspace (no plan)~~
- [ ] ~~Workspace users (no plan)~~
- [ ] ~~View Options (no plan)~~
- [ ] User (in progress)
- [ ] ~~Sharing (no plan)~~
- [x] Sections
- [x] Reminders
- [x] Projects
- [ ] ~~Comments (no plan)~~
- [ ] ~~Live Notifications (no plan)~~
- [x] Labels
- [x] Tasks
- [x] Filters

#### RESTful API

- [x] Projects Get Archived
- [x] Tasks Completed By Completion Date
- [x] Tasks Completed By Due Date

#### WebSocket API

- [x] WebSocket Client

[go-shield]: https://img.shields.io/github/go-mod/go-version/CnTeng/todoist-api-go?style=for-the-badge&logo=go
[go-url]: https://golang.org
