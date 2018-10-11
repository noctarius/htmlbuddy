## Wordpress Sanitizer

_Wordpress Sanitizer_ is a tool to fix HTML from imported .docx files.

For usage import the .docx as normal, however it is important to fix image filename before fixing the HTML code.

Afterwards copy the generated HTML out of Wordpress into a simple textfile.

Use the _Wordpress Sanitizer_ to automatically fix the HTML according to the chosen sanitization Javascript program.

Execute the _Wordpress Sanitizer_ as in the following snippet: 

```
./wordpress-sanitizer -configuration=<sanitization.js> </path/to/textfile>
```

The fixed HTML, by default, is printed to stdout. To redirect it into a file use the common OS way, like `> output.file`. 