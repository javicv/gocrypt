![License](https://img.shields.io/github/license/javicv/gocrypt.svg)
[![Build Status](https://travis-ci.com/javicv/gocrypt.svg?branch=main)](https://travis-ci.com/javicv/gocrypt)
![GitHub release](https://img.shields.io/github/release/javicv/gocrypt.svg)
# gocrypt
Symmetric encryption/decryption command line utility

## Instructions
To encrypt or decrypt symply execute like this

```$ gocrypt [-e|-d] PASSWORD VALUE```

where:
* **-e:** Encrypt (the output will be base64-encoded)
* **-d:** Decrypt (the value should be base64-encoded)