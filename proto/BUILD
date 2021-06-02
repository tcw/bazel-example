load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_proto_grpc//go:defs.bzl", "go_grpc_library")

proto_library(
    name = "echo_proto",
    srcs = ["echo.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "echo_go_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/tcw/bazel-example/go/grpc",
    protos = [":echo_proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "echo_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "client-go/proto",
    proto = ":echo_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "proto",
    embed = [":echo_go_proto"],
    importpath = "client-go/proto",
    visibility = ["//visibility:public"],
)
