This app is used to illustrate path-search algorithms of graphs.

# Step 1: install the bundler

Run the following command:

    $ go get -u github.com/asticode/go-astilectron-bundler/...
    
And don't forget to add `$GOPATH/bin` to your `$PATH`.

# Step 2: set up the building environments

To bundle the app for more environments, add an `environments` key to the bundler configuration (`bundler.json`):

```json
"environments": [
  {"arch": "amd64", "os": "linux"},
  {"arch": "386", "os": "windows"}
]
```
    
# Step 3: bundle the app

Run the following commands:

    $ astilectron-bundler -v

The result is in the `output` folder.
