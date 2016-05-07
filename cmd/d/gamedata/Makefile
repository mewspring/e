all: init cut_scenes towners

.PHONY: all

# === [ init ] =================================================================

init: diabdat cel_dump

diabdat:
	@if [ ! -d "$@/" ]; then \
		echo "Please extract \"diabdat.mpq\" to \"diabdat/\" or provide a symlink; e.g."; \
		echo; \
		echo "   ln -s /path/to/extracted_diabdat_mpq/ diabdat"; \
		echo; \
		exit 1; \
	fi

cel_dump:
	@if ! which $@ &> /dev/null ; then \
		echo "Please install the \"cel_dump\" command"; \
		echo; \
		echo "   go get github.com/sanctuary/formats/cmd/cel_dump"; \
		echo; \
		exit 1; \
	fi

.PHONY: init cel_dump

# === [ cut_scenes ] ===========================================================

cut_scenes: cut_scenes/tristram.png

cut_scenes/tristram.png: _dump_/gendata/cutstart/cutstart.png
	mkdir -p $(dir $@)
	cp $< $@

.PHONY: cut_scenes

# === [ towners ] ==============================================================

towners: towners/cow.png

towners/cow.png: _dump_/towners/animals/cow
	mkdir -p $(dir $@)
	montage $</cow_*/*.png -background none -tile x8 $@

# === [ clean ] ================================================================

clean:
	rm -rf _dump_/ cut_scenes/ towners/

.PHONY: clean

# === [ _dump_ ] ===============================================================

_dump_/gendata/cutstart/cutstart.png: diabdat/gendata/cutstart.cel
	cel_dump $(<:diabdat/%=%)

_dump_/towners/animals/cow: diabdat/towners/animals/cow.cel
	cel_dump $(<:diabdat/%=%)
