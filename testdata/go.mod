// This go.mod exists only to keep this directory out of the published module zip.
//
// A Go module zip contains every file under the module root except directories that hold
// their own go.mod. Without this file, the RPM fixtures in this directory ship in every
// release. Xray unpacks the zip of each dependency and indexes those RPMs as CentOS and
// EPEL components of the consuming build.
//
// Do not remove it, and do not add any requirement: nothing builds from this module.
// Tests read the fixtures by relative path, which is unaffected.
module github.com/jfrog/go-rpm/v2/testdata

go 1.15
