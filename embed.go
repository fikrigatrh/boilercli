package main

import "embed"

//go:embed template/**
var TemplateFiles embed.FS
