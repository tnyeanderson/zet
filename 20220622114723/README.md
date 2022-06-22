# Renumber windows in tmux

Sometimes window numbers get crazy in `tmux`. Use the `move-window` command
with the `-r` (renumber) flag to automatically make them sane!

This can be accomplished using `[tmux prefix] + :` to open the command line.
Then type `move-window -r` and press enter. Done!

> NOTE: For me, [tmux prefix] is CTRL+a. The default is CTRL+b.

    #tmux #tips #tools
