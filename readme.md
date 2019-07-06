# GraphQL + Postgres + DataLoader, in Go.
A simple example of using GraphQL to effectively retrieve data from a Postgres database.

## How to run

1. Build a Docker image for a sample `store` database

    ```bash
    $ docker build -t store .
    ```

2. And then run it
    ```bash
    $ docker run --rm -p 5432:5432 store
    ```

3. Build the project
    ```bash
    $ go build
    ```

4. Then run the binary
    ```bash
    $ ./graphqldb
    ```

Now you should be able to able to send GraphQL reqests to `http://localhost:8080/graphql`. For example

```bash
$ cat <<EOF | curl -X POST -d @- localhost:8080/graphql | python -m json.tool
{
   "query": "{ 
        orders(first: 2) { 
            id, 
            time,
            customer {
                name,
            },
            products {
                name,
                quantity,
                price,
            },
        }
    }"
}
EOF
```
