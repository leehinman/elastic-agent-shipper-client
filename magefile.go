// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	devtools "github.com/elastic/elastic-agent-libs/dev-tools/mage"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	protoDest = "./pkg/proto"
)

var (

	// Add here new packages that have to be compiled.
	// Vendor packages are not included since they already have compiled versions.
	// All `.proto` files in the listed directories will be compiled to Go.
	protoPackagesToCompile = []string{
		"api",
		"api/messages",
	}

	// List all the protobuf packages that need to be included
	protoPackages = append(
		protoPackagesToCompile,
		"api/vendor",
	)

	// Add here files that have their own license that must remain untouched
	goLicenserExcluded = []string{
		"api/vendor",
		"api/messages/struct.proto",
		"pkg/proto/messages/struct.pb.go",
	}
)

// Update regenerates the Go files out of .proto files
func Update() error {
	var (
		importFlags []string
		toCompile   []string
	)

	for _, p := range protoPackages {
		importFlags = append(importFlags, "-I"+p)
	}

	for _, p := range protoPackagesToCompile {
		log.Printf("Listing the %s package...\n", p)

		files, err := ioutil.ReadDir(p)
		if err != nil {
			return fmt.Errorf("failed to read the proto package directory %s: %w", p, err)
		}
		for _, f := range files {
			if path.Ext(f.Name()) != ".proto" {
				continue
			}
			toCompile = append(toCompile, path.Join(p, f.Name()))
		}
	}

	args := append(
		[]string{
			"--go_out=" + protoDest,
			"--go-grpc_out=" + protoDest,
			"--go_opt=paths=source_relative",
			"--go-grpc_opt=paths=source_relative",
		},
		importFlags...,
	)

	args = append(args, toCompile...)

	log.Printf("Compiling %d packages...\n", len(protoPackages))
	err := sh.Run("protoc", args...)
	if err != nil {
		return fmt.Errorf("failed to compile protobuf: %w", err)
	}

	log.Println("Adding license headers...")

	return License()
}

// License applies the right license header.
func License() error {
	mg.Deps(devtools.InstallGoLicenser)
	return licenser(rewriteHeader)
}

// Check runs all the checks
func Check() error {
	mg.Deps(devtools.Deps.CheckModuleTidy, CheckLicenseHeaders)
	mg.Deps(devtools.CheckNoChanges)
	return nil
}

// CheckLicenseHeaders checks ASL2 headers in .go files
func CheckLicenseHeaders() error {
	mg.Deps(devtools.InstallGoLicenser)
	return licenser(checkHeader)
}

type licenserMode int

var (
	rewriteHeader licenserMode = 1
	checkHeader   licenserMode = 2
)

func licenser(mode licenserMode) error {
	var args []string

	switch mode {
	case checkHeader:
		args = append(args, "-d")
	}

	for _, e := range goLicenserExcluded {
		args = append(args, "-exclude", e)
	}

	args = append(args, "-license", "Elastic")

	// go-licenser does not support multiple extensions at the same time,
	// so we have to run it twice

	err := sh.RunV("go-licenser", append(args, "-ext", ".go")...)
	if err != nil {
		return fmt.Errorf("failed to process .go files: %w", err)
	}

	err = sh.RunV("go-licenser", append(args, "-ext", ".proto")...)
	if err != nil {
		return fmt.Errorf("failed to process .proto files: %w", err)
	}

	return nil
}
