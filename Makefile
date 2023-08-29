a:
	rm out/main && go build -o out/main && ./out/main

f:
	echo 123 1> static/thunder/123.txt && rm static/thunder/* && make a

or:
	open https://github.com/guimassoqueto/scrap-fiber