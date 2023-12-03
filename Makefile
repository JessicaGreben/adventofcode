# Usage: To make a copy of the template
# 	$ make copy year=2023 day=01
copy:
	mkdir -p $(year)/$(day)/
	cp -r template/* $(year)/$(day)/
