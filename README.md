Defaults
========

Enforce default values on struct fields.

```go
type User struct {
	Name     string  `default:"Goku"`
	Power    float64 `default:"9000.01"`
}

var u User

err := defaults.Apply(&u)
if err != nil {
	log.Fatal("Uh oh: %v", err)
}

fmt.Print(u.Name)  // Goku
fmt.Print(u.Power) // 9000.01
```

Defaults are only applied to fields at their zero value.

```go
type Config struct {
	Host  *string `default:"0.0.0.0"`
	Port  *int    `default:"8000"`
}

var cfg Config
json.Unmarshal([]byte(`{Host: "charm.sh"}`), &cfg)

if err := defaults.Apply(&cfg); err != nil {
	log.Fatal("Rats: %v", err)
}

fmt.Print(cfg.Host) // charm.sh
fmt.Print(cfg.Port) // 8000
```

Works well with JSON, Yaml, and so on.

```go
type Config struct {
    Host  string `json:"host" default:"0.0.0.0"`
    Port  int    `json:"port" default:"8000"`
    Debug bool   `json:"debug" default:"true"`
}
```

## Supported Types

The following types are supported:

* `string`
* `bool`
* `int`
* `int8`
* `int16`
* `int32` (and `rune`, with some caveats)
* `int64`
* `uint`
* `uint8`
* `uint16`
* `uint32`
* `uint64`
* `float32`
* `float64`
* `[]byte`/`[]uint8`

…as well as pointers to those types.

### Embedded Structs

Embedded structs are supported. The following will parse as expected:

```go
type GroceryList struct {
	Fruit struct {
		Bananas int `default:"8"`
		Pears   int `default:"12"`
	}
	Vegetables *struct {
		Artichokes    int `default:"4"`
		SweetPotatoes int `default:"16"`
	}
}
```

Embedded structs do not need a `default` tag in order to be parsed. Embedded
structs that are `nil` will be initialized with their zero value so they can be
parsed accoringly.

### Runes and Int32s

In Go `rune` is just an alias for `int32`. This presents some ambiguity when
parsing default values. For example, should `"1"` be parsed as a literal `1` or
as a unicode `'1'` (which has the `int32` value of `49`)?

Because of this ambiguity we recommend avoiding setting defaults on `rune`s.
That said, this package defaults to parsing `int32` as integers. Failing that,
it tries to parse them as a `rune`.

```go
// This works as expected...
type Cool struct {
	Fave32BitInteger int32 `default:"12"`
	FaveChar         rune  `default:"a"`
}

// ...but these will not.
type UhOh struct {
	FaveChar rune `default:"3"`  // this is a unicode ETX or ctrl+c
	FaveChar rune `default:"97"` // this is a unicode `a`
}
```

## Licnse

[MIT](https://github.com/meowgorithm/defaults/raw/master/LICENSE)
