## checkmail

Output number of unread emails for Maildir format inboxes.

### Usage

``` sh
checkmail -c checkmailrc
```

It can be used in `tmux` status bar:

``` sh
set-option -g status-interval 20
set -g status-right '#(exec checkmail -c /home/foo/.checkmailrc)'
```
