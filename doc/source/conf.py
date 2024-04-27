# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

project = 'WebGPU Headers'
copyright = '2024, WebGPU Contributors'
author = 'WebGPU Contributors'
release = 'v0.1.0'

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = [
    # The breathe extension is used to get docstring data as parsed by Doxygen
    'breathe'
]

templates_path = ['_templates']
exclude_patterns = []

# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

html_theme = 'alabaster'
html_static_path = ['_static']

# -- Options for breathe -------------------------------------------------
# https://github.com/breathe-doc/breathe

breathe_projects = {
    "webgpu": "../generated/xml"
}
breathe_default_project = "webgpu"
