A simple tools to generate kubernetes resource documentation。 Output result see `doc/`

**Easy to use**
```bash
# generate documents for all default resource 
go run main.go 
# generate a use-defined 
go run main.go -s "pod deployment"
```