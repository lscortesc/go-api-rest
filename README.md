# Simple API Rest with Go, Mux, Gorm & MySql

This api is based on this [post](https://dev.to/aspittel/how-i-built-an-api-with-mux-go-postgresql-and-gorm-5ah8)

## How to use?

1. Clone repo
2. Create project database
3. Create `.env` file based on `.env-example`
4. Configure `.env` file
5. Run `go run main.go`
6. Listen on `localhost:4000`

### Endpoints


| Endpoint | Method | Params | Description |
| -------- | ------ | ------ | ----------- |
| /people  | GET | none   | Get all People |
| /people/{id} | GET | id | Get Person with ID |
| /people | POST | Firstname, Lastname | Create a person |
| /people/{id} | DELETE | id | Delete a person |
| /people/list | POST | none | Create 1000 records |

### Smart commits with Jira

This is another example with Jira and Github


