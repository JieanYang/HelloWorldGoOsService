# AnsysCSPAgentManagerServiceApp

We use this application to monitor the running status of the agent. If the agent application crashes or stops running, our application will restart the Agent.

This application will become a service in the system (service for Linux and windows service for Windows) and will run every time the system restarts to ensure the agent can run stably for a long period of time.

## Build App

```
go build -o ansysCSPAgentManagerServiceApp main.go
```

## Structure

- /main.go: This is the entry point of the project.
- \*.sh: The installation and uninstallation scripts for Linux.
- AnsysCSPAgentManagerService.service: configuration for Linux service.
- \*.ps1: The installation and uninstallation scripts for Windows.
