# PicoCTF 2025

## Fantasy CTF

> Play this short game to get familiar with terminal applications and some of the most important rules in scope for picoCTF.  
> Author: SYREAL  
> [`Fantasy CTF`](general_skills)

**Tags:** general_skills

---

## 🧾 Overview

This challenge provided an interactive text-based game where you guide Eibhilin, a young student, through the picoCTF registration process and learn key competition rules with the help of her AI assistant Nyx. By making the correct choices, you complete the simulation and reveal the flag.

---

## ✅ Solution

I connected to the challenge using:

```bash
nc verbal-sleep.picoctf.net 52486
```

The game began with several press Enter to continue prompts introducing the characters and setting. Eventually, Nyx presented the first major choice:

```
Options:
A) Register multiple accounts
B) Share an account with a friend
C) Register a single, private account
[a/b/c] >
```

I selected:
```
c
```

because creating a single, private account is the ethical and rule-abiding choice.

Next, Nyx introduced the sanity challenge:

```
Options:
A) Play the game
B) Search the Ether for the flag
[a/b] >
```

I selected:
```
a
```
because the challenge was designed to be played interactively, not bypassed.

After completing the game, Eibhilin announced she had found the flag, and Nyx reminded about good practices like not sharing flags or publishing writeups before the competition ends.

The simulation concluded by displaying the flag:

```
picoCTF{m1113n1um_3d1710n_dc6af91c}
```
---

## Flag
picoCTF{m1113n1um_3d1710n_dc6af91c}
