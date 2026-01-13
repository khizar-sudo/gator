# Feed Aggregator (Gator)

A command-line RSS feed aggregator built with Go. Gator allows you to track and browse RSS feeds from your terminal.

## Prerequisites

Before running this program, you'll need to have the following installed on your system:

- **Go** (version 1.20 or higher) - [Download and install Go](https://go.dev/doc/install)
- **PostgreSQL** - [Download and install PostgreSQL](https://www.postgresql.org/download/)
- **Goose** - Database migration tool

Make sure PostgreSQL is running and you have access to create databases.

### Installing Goose

Install goose using Go:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Installation

Install the `gator` CLI using Go's install command:

```bash
go install github.com/khizar-sudo/gator@latest
```

This will download, compile, and install the `gator` binary to your `$GOPATH/bin` directory (or `$GOBIN` if set).

Make sure your `$GOPATH/bin` is in your system's `PATH` to run `gator` from anywhere:

```bash
# Add this to your ~/.bashrc or ~/.zshrc
export PATH=$PATH:$(go env GOPATH)/bin
```

## Database Setup

1. Create a PostgreSQL database for the feed aggregator:

```sql
CREATE DATABASE gator;
```

2. Run the database migrations using goose:

```bash
cd sql/schema
goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

Replace `username` and `password` with your PostgreSQL credentials.

This will automatically run all migration files in the correct order. Goose tracks which migrations have been applied, so it's safe to run this command multiple times.

### Other Useful Goose Commands

- **Check migration status:**

  ```bash
  goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" status
  ```

- **Rollback the last migration:**

  ```bash
  goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" down
  ```

- **Reset database (rollback all migrations):**
  ```bash
  goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" reset
  ```

## Configuration

Before running the program, you need to create a configuration file in your home directory:

**File location:** `~/.gatorconfig.json`

**Contents:**

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` and `password` with your PostgreSQL credentials. The `current_user_name` field will be automatically populated when you register or login.

## Usage

The general syntax for running commands is:

```bash
gator <command> [arguments]
```

### Available Commands

#### User Management

- **Register a new user:**

  ```bash
  gator register <username>
  ```

  Creates a new user and automatically logs you in.

- **Login as an existing user:**

  ```bash
  gator login <username>
  ```

  Sets the current user in the config file.

- **List all users:**
  ```bash
  gator users
  ```
  Displays all registered users.

#### Feed Management

- **Add a new RSS feed:**

  ```bash
  gator addfeed <feed_name> <feed_url>
  ```

  Example: `gator addfeed "Blog" https://blog.example.com/feed.xml`

  This command automatically follows the feed after adding it.

- **List all feeds:**

  ```bash
  gator feeds
  ```

  Shows all RSS feeds in the system and their creators.

- **Follow a feed:**

  ```bash
  gator follow <feed_url>
  ```

  Start following an existing feed.

- **List feeds you're following:**

  ```bash
  gator following
  ```

  Shows all feeds that you're currently following.

- **Unfollow a feed:**
  ```bash
  gator unfollow <feed_url>
  ```
  Stop following a feed.

#### Aggregation and Browsing

- **Start the aggregator:**

  ```bash
  gator agg <time_duration>
  ```

  Example: `gator agg 1m` or `gator agg 30s`

  Fetches new posts from all feeds at the specified interval. This command runs continuously until you stop it (Ctrl+C).

- **Browse recent posts:**

  ```bash
  gator browse [limit]
  ```

  Example: `gator browse 10`

  Displays recent posts from feeds you're following. If no limit is specified, it shows 2 posts by default.

#### Utility

- **Reset the database:**
  ```bash
  gator reset
  ```
  Deletes all users and cascades to feeds and posts (use with caution!).

## Example Workflow

Here's a typical workflow to get started:

```bash
# 1. Register a new user
gator register john

# 2. Add some RSS feeds
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
gator addfeed "Golang Weekly" https://golangweekly.com/rss

# 3. View all feeds
gator feeds

# 4. Start the aggregator in the background (fetch every minute)
gator agg 1m &

# 5. Browse recent posts
gator browse 5
```

## Development

This project uses:

- [sqlc](https://sqlc.dev/) for type-safe SQL queries
- [goose](https://github.com/pressly/goose) for database migrations
- PostgreSQL for data storage
- Native Go RSS parsing

## License

This project is part of the Boot.dev curriculum.
