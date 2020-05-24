all: reconstruct invert rand tree_depth tree_depth2 tree_paths \
	drawtree bous order search cousins readtree bous prune_tree \
	minimal_ht_tree minimal_ht_tree2 testmin lca

images: reconstruct.png invert.png rand.png

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

reconstruct.png: reconstruct.dot
	dot -Tpng -o reconstruct.png reconstruct.dot
reconstruct.dot: reconstruct
	./reconstruct > reconstruct.dot
reconstruct: reconstruct.go
	go build reconstruct.go

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

clean:
	-rm -rf invert invert.dot invert.png
	-rm -rf reconstruct reconstruct.png reconstruct.dot
	-rm -rf rand rand.png rand.dot
	-rm -rf drawtree *.png *.dot 
	-rm -rf tree_depth tree_depth2 tree_paths bous order search
	-rm -rf cousins prune_tree readtree minimal_ht_tree minimal_ht_tree2
	-rm -rf testmin lca
