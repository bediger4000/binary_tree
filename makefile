all: reconstruct reconstruct2 invert rand tree_depth tree_depth2 tree_paths \
	drawtree bous order search cousins readtree prune_tree \
	minimal_ht_tree minimal_ht_tree2 testmin lca nread bottomview \
	ht_balanced

images: reconstruct.png reconstruct2.png invert.png rand.png

test: all
	cd tree; go test -v .
	./runtests

minimal_ht_tree: minimal_ht_tree.go
	go build minimal_ht_tree.go

minimal_ht_tree2: minimal_ht_tree2.go
	go build minimal_ht_tree2.go

testmin: testmin.go
	go build testmin.go

prune_tree: prune_tree.go
	go build prune_tree.go

cousins: cousins.go
	go build cousins.go

nread: nread.go
	go build nread.go

readtree: readtree.go
	go build readtree.go

drawtree: drawtree.go
	go build drawtree.go

bous: bous.go
	go build bous.go

lca: lca.go
	go build lca.go

order: order.go
	go build order.go

search: search.go
	go build search.go

minht.png: testmin.load testmin
	./testmin > min.dat
	gnuplot < testmin.load

reconstruct.png: reconstruct.dot
	dot -Tpng -o reconstruct.png reconstruct.dot
reconstruct2.png: reconstruct2
	./reconstruct2 '(a(b(d(p)())(e()(q)))(c(f()(r))(g(s)(t))))' > reconstruct2.dot
	dot -Tpng -o reconstruct2.png reconstruct2.dot
reconstruct.dot: reconstruct
	./reconstruct > reconstruct.dot
reconstruct: reconstruct.go
	go build reconstruct.go
reconstruct2: reconstruct2.go
	go build reconstruct2.go

invert.png: invert.dot
	dot -Tpng -o invert.png invert.dot
invert.dot: invert
	./invert > invert.dot
invert: invert.go
	go build invert.go

rand.png: rand.dot
	dot -Tpng -o rand.png rand.dot
rand.dot: rand
	./rand > rand.dot
rand: rand.go
	go build rand.go

tree_depth: tree_depth.go
	go build tree_depth.go

tree_depth2: tree_depth2.go
	go build tree_depth2.go

tree_paths: tree_paths.go
	go build tree_paths.go

bottomview: bottomview.go
	go build bottomview.go

ht_balanced: ht_balanced.go
	go build ht_balanced.go

clean:
	-rm -rf invert invert.dot invert.png
	-rm -rf reconstruct reconstruct.png reconstruct.dot
	-rm -rf reconstruct2 reconstruct2.png reconstruct2.dot
	-rm -rf rand rand.png rand.dot
	-rm -rf drawtree *.dot 
	-rm -rf tree_depth tree_depth2 tree_paths bous order search
	-rm -rf cousins prune_tree readtree minimal_ht_tree minimal_ht_tree2
	-rm -rf testmin lca nread bottomview ht_balanced
	-rm -rf min.dat
