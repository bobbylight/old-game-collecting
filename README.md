# Old Game Collecting
A webapp allowing you to track your NES and SNES game collection.
Small and simple.

## Hacking
The backend is in go and the frontend is in Vue.  Data is stored in
Postgres.

To set up a local Postgres instance with the game data:

```bash
cp util/env.sh.orig util/env.sh
vi util/env.sh # Make changes to point to your local Postgres
./util/import-from-csv.sh
```

To start the backend server:

```go
go run hello.go
```

To start the frontend compiling with a watch for development purposes:

```bash
cd frontend
npm install
npm run watch
```
