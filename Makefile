.PHONEY: p
p:
	hugo new post/${TITLE}.md

.PHONEY: b
b:
	hugo
