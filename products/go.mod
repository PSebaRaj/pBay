module github.com/psebaraj/pbay/products

go 1.17

require internal/config v1.0.0

replace internal/config => ./internal/config

require gopkg.in/yaml.v3 v3.0.1
