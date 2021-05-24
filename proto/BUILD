load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "ibsen_proto",
    srcs = ["ibsen.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "grpcApi_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "/api/grpcApi",
    proto = ":ibsen_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "grpcApi",
    embed = [":grpcApi_go_proto"],
    importpath = "/api/grpcApi",
    visibility = ["//visibility:public"],
)
