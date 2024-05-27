# RSS Aggregator in Go
This is a project to build a fully fledged back-end server in Go from scratch on your local machine. The purpose of the server is to aggregate data from RSS feeds. RSS feeds are a way to distribute information like podcasts and blog posts.

## The server will allow users to:

- Add different RSS feeds to its database
- Automatically collect all the posts from those feeds
- Download and save the posts in the database
- View the posts later

## Running the Server

Create a .env file to store configuration secrets for your project, such as the port the server will run on. Add the following line to the .env file, replacing 8000 with your desired port number:

`PORT=8000`

#### Setup Postgres database and add config to .env file:

`DB_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable`

#### Executing the binaries

For Windows:

`go build`

`.\rss_scraper.exe`

For Linux:

`go build && ./rss_scraper`

##### Use an HTTP client like Thunder Client to send requests to the server's API endpoints.
