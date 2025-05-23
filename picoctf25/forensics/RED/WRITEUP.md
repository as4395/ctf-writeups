# PicoCTF 2025

## Red

> RED, RED, RED, RED  
> Author: SHUAILIN PAN (LECONJUROR)  
> [`Red`](red.png)

**Tags:** forensics

---

## 🧾 Overview

This challenge gave us a file called `red.png`, which initially appeared to be a plain red square. However, the repetitive emphasis on "RED" and the reference to poetry hinted that there might be hidden data embedded in the image. The goal was to extract the flag hidden using steganography.

---

## ✅ Solution

I began by inspecting the file using `exiftool`:

```bash
exiftool red.png
