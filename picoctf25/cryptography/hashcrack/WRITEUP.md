# PicoCTF 2025

## hashcrack

> A company stored a secret message on a server which got breached due to the admin using weakly hashed passwords. Can you gain access to the secret stored within the server?  
> Author: Nana Ama Atombo-Sackey  
> [`hashcrack`](hashcrack)

**Tags:** cryptography

---

## 🧾 Overview

In this challenge, the goal was to connect to a remote server and crack a series of password hashes: an MD5 hash, a SHA-1 hash, and a SHA-256 hash. Only after successfully cracking all three would the server reveal the flag. I cracked the first two manually using `shasum`/`md5sum` on macOS, and used CrackStation to handle the final SHA-256 hash when it wasn’t easily found in my local attempts.

---

## ✅ Solution

I connected to the challenge server using:

```bash
nc verbal-sleep.picoctf.net 51759
```

The server gave me the first hash.

---

### 🔐 Step 1: MD5 Hash

**Hash received:**
```
482c811da5d5b4bc6d497ffa98491e38
```

I identified this as MD5 based on its 32-character hexadecimal format. I tested common values manually using:

```bash
echo -n 'password123' | md5sum
```

The result matched the challenge hash:

```
482c811da5d5b4bc6d497ffa98491e38  -
```

✅ So I entered `password123` in the prompt and it accepted the input.

---

### 🔐 Step 2: SHA-1 Hash

Next, the server gave me:

```
b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3
```

This is a 40-character hex hash — clearly SHA-1. I repeated the same local process:

```bash
echo -n 'letmein' | sha1sum
```

Output:

```
b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3  -
```

✅ This matched! So I entered `letmein`, which was also accepted by the server.

---

### 🔐 Step 3: SHA-256 Hash

The final hash was:

```
916e8c4f79b25028c9e467f1eb8eee6d6bbdff965f9928310ad30a8d88697745
```

This 64-character hex string indicated a SHA-256 hash. I attempted to crack it manually using Bash and `rockyou.txt`:

```bash
while read word; do
  hash=$(printf "%s" "$word" | shasum -a 256 | awk '{print $1}')
  if [ "$hash" = "916e8c4f79b25028c9e467f1eb8eee6d6bbdff965f9928310ad30a8d88697745" ]; then
    echo "Match found: $word"
    break
  fi
done < rockyou.txt
```

I even trimmed `rockyou.txt` to its top 100k and 1M lines to speed up the process, but still no match appeared locally. After exhausting manual attempts, I switched to CrackStation:

### 🔎 Using CrackStation

I visited [https://crackstation.net](https://crackstation.net), pasted the SHA-256 hash into the input box, solved the CAPTCHA, and clicked “Crack Hashes.” The site returned:

```
qwerty098
```

✅ I returned to the challenge and entered `qwerty098` as the final password — and the flag was revealed.

---

## Flag

```
picoCTF{UseStr0nG_h@shEs_&PaSswDs!_ce730f64}
```
