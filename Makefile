.PHONY: doc
doc:
	go get github.com/lanre-ade/godoc2md \
	&& godoc2md github.com/kaluza-tech/urls > ./DOCUMENTATION.md
