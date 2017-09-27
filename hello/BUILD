package(default_visibility = ["//visibility:public"])

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_library")

go_proto_library(
    name = "hello",
    protos = [
        "hello.proto",
    ],
    with_grpc = True,
)
