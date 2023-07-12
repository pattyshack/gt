### Bash auto-completion:
To enable bash auto-completion for the "example" binary, set the following
in your `.bashrc` or create a file in `/etc/bash_completion.d/` with the following content (and restart bash shell):

```
_example()
{
  local words
  local cword

  # XXX: consider disabling all COMP_WORDBREAKS characters: "'><=;|&(:
  _get_comp_words_by_ref -n = cword words

  IFS=$'\n'
  COMPREPLY=( $(${words[0]} --auto-suggest -- $cword ${words[@]}) )
}
complete -o nospace -F _example example
```

Self-reminder: cword (the current cursor position) can be beyond words's (the argument list) length, requesting the next argument.

### Resources:

[Intro to bash autocompletion](http://web.archive.org/web/20180304191616/https://debian-administration.org/article/317/An_introduction_to_bash_completion_part_2)
[Stackoverflow q&a](https://stackoverflow.com/questions/10528695/how-to-reset-comp-wordbreaks-without-affecting-other-completion-script/12495480)

[Bash manual - Bash variables COMP\*](https://www.gnu.org/software/bash/manual/bash.html#Bash-Variables)

[Bash manual - Programmable Completion](https://www.gnu.org/software/bash/manual/bash.html#Programmable-Completion)
