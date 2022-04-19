# Schemas
1) Key names
1a) required
2) key:value types
3) acceptable values ( support both enum or [x..y] notation, ideally )


Schemas should be in toml. each header can be a key?

```toml
[title]
type = string

```

at the end of the process, i can check the diff between the set of all metadata keys and the set of keys configured in the schema and output a list of all "untracked" keys. this will catch typos on optional keys
