# Treefy
This is a small CLI that creates a tree view for the directory. In addition, if ```-f``` flag is provided, it can print out files and their size. See the examples below:

With ```-f``` flag
```
treefy testdata -f
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
```
```
treefy testdata
└───testdata
	├───project
	├───static
	│	├───css
	│	├───html
	│	└───js
	└───zline
```
## Installation
* Make sure that you have ```Go``` installed.
* In the project root run:
```
go install
```
## Usage
The program accepts arguments in the following format:
```
treefy [path] [-f]
```
```[path]``` -- path to the directory
```[-f]``` -- optional flag: output files
