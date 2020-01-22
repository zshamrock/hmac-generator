# hmac-generator

HMAC HTTP Authorization header generator.

```
NAME:
   hmac-generator - Generates HMAC authorization HTTP Header

USAGE:
   hmac-generator --id <key id> --secret/--secret-file <value>

VERSION:
   1.0.0

AUTHOR:
   (c) Aliaksandr Kazlou

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --id value                       authorization key id
   --secret value, -s value         authorization secret
   --secret-file value, --sf value  file from which to read authorization secret
   --help, -h                       show help
   --version, -v                    print the version

```

## Description

```
$ hmac-generator --id foo --secret bar
Authorization: HMAC ts=1579862657754,id=foo,nonce=3396422525437371841,mac=l4MFVlY2zYiGk1bhMME/4TDr9k6U85ATwIySP0+F4GQ=
```

The tool produces HMAC custom line which could be used in the HTTP `Authorization` header to secure API calls, for example.

It will although require the implementation on the chosen backend system to parse that value, generate and verify the 
signature against the one received in the request.

Format of the generated authorization line:

- `HMAC` - Authorization type.
- `ts` - Unix timestamp (in milliseconds) value of the time when the signature/token has been generated. Could be used 
    to control the expiration/validity of the generated token.
- `id` - Key/client id, used to identify the client/caller on the backend, fetch the corresponding secret for this client,
    build the token and verify with the one from the request.
- `nonce` - Random generated value to "salt" the generated token.
- `mac` - Produced token. See [Algorithm](#algorithm) below on how the token is generating.

## Algorithm

1. Build `HMAC` hash (see you chosen language for the available implementation, below are given sample implementations 
in Go and Java) using obtained secret for the client `id`.
2. Concatenate `ts` and `nonce` together, as a string value, i.e. `1579518463570` + `5696149536374835586`
will result into `15795184635705696149536374835586`. 
3. Generate resulting token by appending concatenated above value into the HMAC.
           

### Go

```
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "strconv"
)
...
mac := hmac.New(sha256.New, []byte(secret))
mac.Write([]byte(strconv.Itoa(int(timestamp)) + strconv.Itoa(nonce)))
sum := base64.StdEncoding.EncodeToString(mac.Sum(nil))
...
```

### Java

```
import java.nio.charset.StandardCharsets;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

...
final Mac macSHA256;
try {
    macSHA256 = Mac.getInstance(HMAC_SHA_256);
    macSHA256.init(new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), HMAC_SHA_256));
} catch (final NoSuchAlgorithmException | InvalidKeyException ex) {
    // handle error
}
final String data = timestamp + "" + nonce;
final String sum = Base64.getEncoder().encodeToString(
    macSHA256.doFinal(data.getBytes(StandardCharsets.UTF_8)));
...
```

## Installation

```
$ go get github.com/zshamrock/hmac-generator
```

## Copyright                                                                                                                                                 
                                                                                                                                                             
Copyright (C) 2020 by Aliaksandr Kazlou.                                                                                                                     
                                                                                                                                                             
hmac-generator is released under MIT License.                                                                                                                       
See [LICENSE](https://github.com/zshamrock/hmac-generator/blob/master/LICENSE) for details.
