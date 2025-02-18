# Hex Editor

This is a simple command-line hex editor written in Go. It allows you to view the contents of a file in hexadecimal format and also edit it interactively in a hex view using Neovim. You can use the -w flag to edit the file in a hex editor mode and save the changes back to the original file.
Features

  * View a file in hexadecimal format.
  * Edit files interactively in Neovim, while displaying the content in hexadecimal.
  * Write back the modified content to the original file.

# Prerequisites

Make sure you have the following installed on your system:

  * Go (version 1.16 or later)
  - Neovim (for the -w flag to work... Well you can replace "nvim" with any other editor in the code and it will still work)

# Installation
Clone the repository to your local machine:

    git clone https://github.com/prajjwal000/hex_editor.git

Navigate to the project directory:

    cd hex-editor/cmd/term

Build the program:

    go build .

Or if you want to run it without building, you can directly use:

    go run .

# Usage
To view the contents of a file in hexadecimal format, run the following command:

    go run . filename
To open a file in Neovim for editing in hexadecimal mode, use the -w flag:

    go run . -w filename
  **_Make sure to use tabs (\t) between hexadecimal byte values while editing and avoid spaces to maintain the correct format._**

This will:

  * Show the file content in hexadecimal format.
  * Open the file in Neovim for editing.
  * Allow you to modify the file in hex.
  * After editing and closing Neovim, the program will save the changes back to the original file.
