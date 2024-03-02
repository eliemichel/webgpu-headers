#!/usr/bin/python3

#
# This file is part of the WebGPU native specification.
#   https://github.com/webgpu-native/webgpu-headers
#
# Copyright 2019-2024 WebGPU-Native developers
#
# SPDX-License-Identifier: BSD-3-Clause
#

import yaml
from os.path import dirname, join
import logging

# -----------------------------------------------------------------------------

DEFAULT_YAML_PATH = join(dirname(dirname(__file__)), "webgpu.yml")
DEFAULT_MKDOCS_SRC_PATH = join(dirname(__file__), "src")

logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")

# -----------------------------------------------------------------------------

def makeArgParser():
    import argparse

    parser = argparse.ArgumentParser(description="""
    Generate the WebGPU native documentation source from the webgpu.yml spec.
    """)

    parser.add_argument("-y", "--yaml", type=str,
                        default=DEFAULT_YAML_PATH,
                        help="Path to the WebGPU native specification YAML file (webgpu.yml).")

    parser.add_argument("-m", "--mkdocs", type=str,
                        default=DEFAULT_MKDOCS_SRC_PATH,
                        help="Path to directory of mkdocs source where this script must write auto-generated documentation files.")

    return parser

# -----------------------------------------------------------------------------

def main(args):
    with open(args.yaml, 'r', encoding="utf-8") as f:
        spec = yaml.safe_load(f)

    api_path = join(args.mkdocs, "api.md")
    logging.info(f"Writing doc source file '{api_path}'...")
    with open(api_path, 'w', encoding="utf-8") as f:
        f.write("# API\n\n")

        f.write("## Constants\n\n")

        for entry in spec["constants"]:
            name = entry["name"]
            value = entry["value"]
            doc = entry["doc"]
            f.write(f" - `WGPU_{constantCase(name)}` = `{cValue(value)}`\n")

# -----------------------------------------------------------------------------
# Formatters

def constantCase(name):
    return name.upper()

def cValue(value):
    return {
        "usize_max": "SIZE_MAX",
        "uint32_max": "0xffffffffUL",
        "uint64_max": "0xffffffffffffffffULL",
    }.get(value, value)

# -----------------------------------------------------------------------------

if __name__ == "__main__":
    args = makeArgParser().parse_args() 
    main(args)
