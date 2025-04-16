ASCII Art Banner Generator
This Go project allows you to generate ASCII art representations of strings and save them into a file. It supports multiple banner styles and provides a flexible command-line interface (CLI) for generating, viewing, and saving banners.

Features
Generate ASCII art banners from text input.
Save the generated banners into a text file.
Supports multiple banner styles, e.g., "standard", "shadow", etc.
Simple command-line usage with customizable output filenames.
Requirements
Go 1.18 or higher
Go modules enabled
Installation
To install the project, you need to have Go installed on your machine. Then, clone the repository and navigate to the project folder.

bash
Copy
git clone <repository_url>
cd <project_directory>
go mod tidy
Usage
You can generate a banner and save it to a file by using the following command format:

bash
Copy
go run . --output=<fileName.txt> "<STRING>" <BANNER_STYLE>
Arguments:
--output=<fileName.txt>: The flag specifying the output file where the banner will be saved. The file must end with .txt.
<STRING>: The string you want to convert into an ASCII art banner.
<BANNER_STYLE>: The style of the banner. Some common styles are:
standard
shadow
Additional styles may be added as the project grows.
Examples
Example 1: Generate a "standard" banner with the text "hello" and save it to banner.txt:
bash
Copy
go run . --output=banner.txt "hello" standard
Output (in banner.txt):

txt
Copy
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
Example 2: Generate a "shadow" banner with the text "Hello There!" and save it to banner.txt:
bash
Copy
go run . --output=banner.txt "Hello There!" shadow
Output (in banner.txt):

txt
Copy
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
Example 3: Generate a banner with a single string and no style (defaults to standard):
bash
Copy
go run . --output=banner.txt "Simple Banner"
Output (in banner.txt):

txt
Copy
 _     _               _                 _       
| |   (_)             | |               | |      
| |__  _ _ __   __ _  | |__     __ _ ___| |_ ___ 
| '_ \| | '_ \ / _` | | '_ \   / _` / __| __/ _ \
| | | | | | | | (_| | | | | | | (_| \__ \ ||  __/
|_| |_|_|_| |_|\__,_| |_| |_|\__,_|___/\__\___|
Command-Line Arguments
--output=<fileName.txt>: Specifies the file to save the ASCII banner.
<STRING>: The string that will be converted into ASCII art.
<BANNER_STYLE>: Optional argument that specifies the style of the banner (e.g., standard, shadow). If no style is provided, the default is standard.
Development
To add new banner styles, you can extend the project by implementing additional ASCII art generators. Each banner style is defined in a separate function that takes in a string and outputs the ASCII art.

Testing
Unit tests are recommended for verifying the correctness of the code. You can write tests for different banner styles and file handling functions.

Run Tests
To run unit tests (if implemented), use:

bash
Copy
go test ./test