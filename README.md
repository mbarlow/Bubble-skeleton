# Bubble Tea Skeleton

A reusable Go project skeleton for building terminal user interfaces with [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Charm](https://charm.sh/), and [Lipgloss](https://github.com/charmbracelet/lipgloss).

## Features

- 🏗️ **Well-structured architecture** - Modular design with separate packages for models, views, commands, config, and styles
- 🎨 **Beautiful styling** - Pre-configured Lipgloss styles with customizable themes
- ⌨️ **Input handling** - Full keyboard navigation and text input with cursor support
- 🔧 **Configuration system** - JSON-based configuration with sensible defaults
- 📦 **Build automation** - Makefile with common development tasks
- 🚀 **Ready to extend** - Clean separation of concerns makes adding features easy

## Project Structure

```
bubble-skeleton/
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── models/           # Bubble Tea models and state management
│   ├── views/            # UI components and rendering logic
│   ├── commands/         # Bubble Tea commands for async operations
│   ├── config/           # Configuration management
│   └── styles/           # Lipgloss styles and theming
├── build/                # Build output directory
├── pkg/                  # Public packages (if needed)
├── Makefile              # Build automation
├── go.mod                # Go module definition
├── go.sum                # Dependency lock file
└── .gitignore           # Git ignore rules
```

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone this repository as a template for your new project:

```bash
git clone https://github.com/mbarlow/Bubble-skeleton.git my-app
cd my-app
rm -rf .git
git init
```

2. Update the module name in `go.mod`:

```bash
go mod edit -module github.com/yourusername/my-app
```

3. Update import paths throughout the project to match your new module name.

4. Install dependencies:

```bash
go mod download
```

### Building and Running

```bash
# Build the application
make build

# Run directly
make run

# Or without Make:
go run cmd/main.go
```

## Usage

### Keyboard Shortcuts

- `h` or `?` - Show help screen
- `i` - Enter input mode
- `r` - Refresh/reload
- `q` or `Ctrl+C` - Quit
- `Esc` - Return to main view

### Development Commands

```bash
# Format code
make fmt

# Run linter (requires golangci-lint)
make lint

# Run tests
make test

# Clean build artifacts
make clean

# Install globally
make install

# Run with live reload (requires air)
make dev

# Show all available commands
make help
```

## Customization Guide

### Adding New Views

1. Define a new state in `internal/models/model.go`:

```go
const (
    StateMain State = iota
    StateHelp
    StateInput
    StateYourNewView  // Add your state here
)
```

2. Create the view rendering in `internal/views/view.go`:

```go
func (v *View) YourNewView() string {
    // Your view rendering logic
}
```

3. Handle the state in the model's `View()` method.

### Adding Commands

Create new async operations in `internal/commands/commands.go`:

```go
func YourAsyncCommand() tea.Cmd {
    return func() tea.Msg {
        // Perform async operation
        return YourCommandCompleteMsg{}
    }
}
```

### Customizing Styles

Modify colors and styles in `internal/styles/styles.go`:

```go
primary := lipgloss.Color("#7D56F4")    // Change primary color
secondary := lipgloss.Color("#F25D94")  // Change secondary color
```

### Configuration

Extend the configuration in `internal/config/config.go`:

```go
type Config struct {
    AppName    string `json:"app_name"`
    Version    string `json:"version"`
    Debug      bool   `json:"debug"`
    // Add your configuration fields here
    YourField  string `json:"your_field"`
}
```

## Configuration File

The application supports a JSON configuration file. By default, it looks for:

```
~/.config/bubble-skeleton/config.json
```

Example configuration:

```json
{
  "app_name": "My App",
  "version": "1.0.0",
  "debug": true,
  "theme": {
    "color_scheme": "default",
    "use_emoji": true
  },
  "keybindings": {
    "quit": ["q", "ctrl+c"],
    "help": ["h", "?"],
    "input": ["i"],
    "refresh": ["r"]
  }
}
```

## Debug Mode

Enable debug mode to write logs to `debug.log`:

```bash
# Via command line flag
./build/bubble-skeleton --debug

# Or via config file
{
  "debug": true
}
```

## Examples

### Creating a List View

```go
// In internal/views/view.go
func (v *View) ListView(items []string, selected int) string {
    return v.styles.List(items, selected)
}
```

### Adding a Network Request

```go
// In internal/commands/commands.go
func FetchData(url string) tea.Cmd {
    return func() tea.Msg {
        resp, err := http.Get(url)
        if err != nil {
            return ErrorMsg{Err: err}
        }
        defer resp.Body.Close()
        
        // Process response
        return DataFetchedMsg{Data: processedData}
    }
}
```

## Contributing

This is a template project designed to be forked and modified. Feel free to use it as a starting point for your Bubble Tea applications!

## License

MIT - See LICENSE file for details

## Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The delightful TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for nice terminal layouts
- [Charm](https://charm.sh/) - The wonderful team behind these tools

## Support

If you find this skeleton helpful, please give it a star ⭐

For questions or issues, please open an issue on GitHub.