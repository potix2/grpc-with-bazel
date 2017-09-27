# golang configuration
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.2"
)
load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")

git_repository(
    name = "org_pubref_rules_protobuf",
    remote = "https://github.com/pubref/rules_protobuf.git",
    tag = "v0.7.2",
)

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")
go_proto_repositories()

go_repositories(
    go_version = "1.8.3",
)

go_repository(
    name = "com_github_grpc_grpc_go",
    importpath = "google.golang.org/grpc",
    commit = "fa1cb32dc4f81e23ab862dd5e7ac4f2920a33088",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    commit = "fa1cb32dc4f81e23ab862dd5e7ac4f2920a33088",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    commit = "09f6ed296fc66555a25fe4ce95173148778dfa85",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    commit = "bb807669a61aca6092d8137da1fab2150bb96ad7",
)
