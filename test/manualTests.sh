# test new
fwatch . 'make install && note new hello.md --tags hey,you -n test && cat $DOTFILES/notes/test/hello.md'

# test find
fwatch . 'make install && note find --tags hey,you -n test'
