# cp-bot (alpha)

A CLI tool using golang for simplifying the process of participating in a
competitive programming contest.

# How to use?

This tool provides the following commands:

## cp-bot contest-url

This command creates the following structure in `~/cp/<platform-name>` directory
by scraping the contest pages. It will also copy `~/cp/template.cpp` to this
directory so you can get started coding right away.

```
<contest-id>
├── <problem-id>
│   ├── exp[0..].txt
│   ├── inp[0..].txt
│   ├── out[0..].txt
│   ├── sol
│   └── sol.cpp
└─...
```

- `sol.cpp` is the main file containing the solution containing your template.
- `i[0..].text` are the files with the input for various test cases.
- `o[0..].cpp` will contain the output of running the respective `i` files with your program.
- `e[0..].cpp` is the expected output for the respective `i` file.

## cp-bot check

Once done with your solution run this command in that problem directory to check
your solution against all corresponding expected outputs and get a nice diff
view.

## cp-bot clean (wip)

After you're done with the contest

# Configuration

A `~/.config/cp-bot/config.yaml` is read to configure some basic settings.

1. _directory_, which directory to keep all the contests in. Defaults to `~/cp`.
2. _template_, which template to use. Defaults to `$cp/template.cpp`.

Example:

```yaml
directory: ~/pichu/raichu
template: ~/abra/kadabra/alkazam.cpp
```

# New Features?

This tool is still in alpha so write in your (https://github.com/ArchitBhonsle/cp-bot/issues)[issues].
Even better, if you have experience in Golang send in
[pull requests](https://github.com/ArchitBhonsle/cp-bot/pulls).

Some things I wanna add but am too bored to:

1. Support for other languages. (Say Java and Python)
2. Windows support. (This I may not be able to do on my own)
3. Support more cp platforms.
4. Individual problems instead of whole contests.
5. Instead of url, entering just the platform name and contest id.
