# Don

I use this package to manage my application donation tiers.

```go
dons, err := don.Parse("don.yml")
if err != nil {
    ...
}

d := dons.Get("premium")
// d.Level
// d.Scopes
// d.Meta

if d.Scope("offline") {
    // Download music on the disk.                
}

if d.Int("accounts") > 1 {
    // Enable multiple account support.
}
```
