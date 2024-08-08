load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

root = "github.com/pattyshack/bt/"

def bt_go_package(
    name,
    path,  # relative to root
    deps = [],  # lib or main deps
    lib_visibility=["//visibility:public"],
    test_deps = [],
    test_size="small",
    is_binary=False):

  importpath = root + path

  lib_name = name
  if is_binary:
    lib_name += "_lib"

  go_library(
    name = lib_name,
    importpath = importpath,
    srcs = native.glob(include=["*.go"], exclude=["*_test.go"]),
    deps = deps,
    visibility = lib_visibility,
  )

  test_glob = native.glob(include=["*_test.go"])
  if test_glob:
    go_test(
      name = name + "_test",
      importpath = importpath,
      srcs = test_glob,
      embed = [
        ":"+lib_name,
      ],
      deps = deps + test_deps,
      size = test_size,
    )

  if is_binary:
    go_binary(
      name = name,
      srcs = native.glob(include=["*.go"], exclude=["*_test.go"]),
      deps = deps,
    )
