Утилита tree.

Выводит дерево каталогов и файлов (если указана опция -f).
```
go run main.go testdata -f
└───testdata
   	├───project  (4096b)
   	│	├───file.txt  (19b)
   	│	└───gopher.png  (70372b)
   	├───static  (4096b)
   	│	├───a_lorem  (4096b)
   	│	│	├───dolor.txt  (empty)
   	│	│	├───gopher.png  (70372b)
   	│	│	└───ipsum  (4096b)
   	│	│	   	└───gopher.png  (70372b)
   	│	├───css  (4096b)
   	│	│	└───body.css  (28b)
   	│	├───empty.txt  (empty)
   	│	├───html  (4096b)
   	│	│	└───index.html  (57b)
   	│	├───js  (4096b)
   	│	│	└───site.js  (10b)
   	│	├───z_lorem  (4096b)
   	│	│	├───dolor.txt  (empty)
   	│	│	├───gopher.png  (70372b)
   	│	│	└───ipsum  (4096b)
   	│	│	   	└───gopher.png  (70372b)
   	├───zline  (4096b)
   	│	├───empty.txt  (empty)
   	│	└───lorem  (4096b)
   	│	   	├───dolor.txt  (empty)
   	│	   	├───gopher.png  (70372b)
   	│	   	└───ipsum  (4096b)
   	│	   	   	└───gopher.png  (70372b)
   	└───zzfile.txt  (empty)
```

```
go run main.go testdata
└───testdata
   	├───project
   	├───static
   	│	├───a_lorem
   	│	│	└───ipsum
   	│	├───css
   	│	├───html
   	│	├───js
   	│	└───z_lorem
   	│	   	└───ipsum
   	└───zline
   	   	└───lorem
   	   	   	└───ipsum
```