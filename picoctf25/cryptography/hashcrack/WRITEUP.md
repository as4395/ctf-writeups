# PicoCTF 2025

## hashcrack

> A company stored a secret message on a server, but the admin used weakly hashed passwords. Can you gain access to the secret stored within the server?  
> Author: Nana Ama Atombo-Sackey  
> [`hashcrack`](hashcrack)

**Tags:** cryptography

---

## 🧾 Overview

In this challenge, I connected to a remote server that displayed a sequence of password hashes. Each hash had to be cracked correctly in order to proceed to the next stage. The final reward was a flag. The task involved identifying hash types based on their format and length, cracking them manually where possible, and using external tools when needed.

---

## ✅ Solution

I started by connecting to the server using netcat:

```bash
nc verbal-sleep.picoctf.net 51759
```

The server then prompted me with three different hashes: one MD5, one SHA-1, and one SHA-256.

---

### 🔐 Step 1: Cracking the MD5 hash

The first hash was:

```
482c811da5d5b4bc6d497ffa98491e38
```

I recognized this as an MD5 hash because it was 32 characters long and entirely hexadecimal. I tested a common password manually in my terminal:

```bash
echo -n 'password123' | md5sum
```

This returned:

```
482c811da5d5b4bc6d497ffa98491e38  -
```

The hash matched. I entered `password123` into the server and moved to the next stage.

---

### 🔐 Step 2: Cracking the SHA-1 hash

The second hash was:

```
b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3
```

This hash was 40 characters long, which fits the SHA-1 format. I tested another common password:

```bash
echo -n 'letmein' | sha1sum
```

The output was:

```
b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3  -
```

The hash matched again. I entered `letmein` and proceeded to the final hash.

---

### 🔐 Step 3: Cracking the SHA-256 hash

The last hash was:

```
916e8c4f79b25028c9e467f1eb8eee6d6bbdff965f9928310ad30a8d88697745
```

This hash was 64 characters long, clearly indicating SHA-256.

I tried to crack it manually using `rockyou.txt`. I used a Bash loop to hash each word in the list and compare it to the target hash:

```bash
target="916e8c4f79b25028c9e467f1eb8eee6d6bbdff965f9928310ad30a8d88697745"

while IFS= read -r word; do
  hash=$(printf "%s" "$word" | shasum -a 256 | awk '{print $1}')
  if [ "$hash" = "$target" ]; then
    echo "Match found: $word"
    break
  fi
done < rockyou.txt
```

To speed up the process, I even tried trimming the wordlist to the top 100,000 and top 1,000,000 entries, but I still didn’t find a match.

---

### 🔎 Final Step: Using CrackStation

Since the password didn’t appear in my local attempts, I turned to [CrackStation](https://crackstation.net), a public online hash-cracking tool. I pasted the SHA-256 hash into the input field, completed the CAPTCHA, and got the result:

```
qwerty098
```

I returned to the challenge and entered `qwerty098`. The server accepted it and revealed the final flag.

---

## Flag

```
picoCTF{UseStr0nG_h@shEs_&PaSswDs!_ce730f64}
```
