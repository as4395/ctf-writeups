# CTF@CIT 2025

## Ask Nicely

> I made this program, you just have to ask really nicely for the flag! 
> Author: ronnie  
> [`asknicely`](asknicely)

**Tags:** reverse_engineering

---

## 🧾 Overview

This challenge provided a 64-bit ELF binary named `asknicely`. Unlike [`read-only`](readonly), this binary was not stripped, which made it easier to analyze using static tools. The goal was to figure out the exact input program expected in order to reveal the flag.

---

## ✅ Solution

I started by inspecting the binary:

```bash
file asknicely
```

**Output:**
```
asknicely: ELF 64-bit LSB executable, x86-64, statically linked, not stripped
```

I then ran `strings` on the file to check for any embedded clues:

```bash
strings asknicely | grep -i please
```

This revealed a very specific input string:

```
pretty pretty pretty pretty pretty please with sprinkles and a cherry on top
```

When I executed the binary normally, it prompted for input:

```
How badly do you want the flag?
Ask nicely...
that's not quite what I'm lookng for.
```

So I entered the full string exactly as shown above:

```bash
$ ./asknicely
pretty pretty pretty pretty pretty please with sprinkles and a cherry on top
```

And got this response:

```
How badly do you want the flag?
Ask nicely...
Good job, I'm so proud of you!
CIT{2G20kX09yF3F}
```

---

## Flag

```
CIT{2G20kX09yF3F}
```
