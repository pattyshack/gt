### Bash auto-completion:
To enable bash auto-completion for the "example" binary, set the following
in your `.bashrc` or create a file in `/etc/bash_completion.d/` with the following content (and restart bash shell):

```
_example()
{
  COMPREPLY=( $(${COMP_WORDS[0]} --auto-suggest -- $COMP_CWORD ${COMP_WORDS[@]}) )
}
complete -F _example example
```

Self-reminder: `COMP_CWORD` (the current cursor position) can be beyond `COMP_CWORD`'s (the argument list) length, requesting the next argument.

### Resources:

[Intro to bash autocompletion](http://web.archive.org/web/20180304191616/https://debian-administration.org/article/317/An_introduction_to_bash_completion_part_2)
