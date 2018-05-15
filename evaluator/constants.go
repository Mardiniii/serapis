package evaluator

import (
	"context"
)

var ctx = context.Background()

var images = map[string]string{
	"node": "node:latest",
	"ruby": "ruby:latest",
}

var extensions = map[string]string{
	"node": "js",
	"ruby": "rb",
}

var packageManagers = map[string]map[string]string{
	"node": map[string]string{
		"installer": "npm install ",
		"versioner": "@",
	},
	"ruby": map[string]string{
		"installer": "gem install ",
		"versioner": ":",
	},
}
