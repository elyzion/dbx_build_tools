package(default_visibility = ["//visibility:public"])

load("//build_tools/py:py.bzl", "dbx_py_pypi_piplib")

dbx_py_pypi_piplib(
    name = "importlib-metadata",
    pip_version = "1.6.0",
    provides = ["importlib_metadata"],
    setup_requires = ["//pip/setuptools_scm"],
    deps = [
        "//pip/configparser",
        "//pip/contextlib2",
        "//pip/pathlib2",
        "//pip/zipp",
    ],
)
