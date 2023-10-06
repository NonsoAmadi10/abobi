# abobi

---

A command line tool that parses your environment variables written in a file into a kubernetes base64 encoded secrets.

[![abobi workflow](https://github.com/NonsoAmadi10/abobi/actions/workflows/main.yml/badge.svg)](https://github.com/NonsoAmadi10/abobi/actions/workflows/main.yml)

## Getting started

To use this package, simply download the binary file:

```bash
wget https://github.com/NonsoAmadi10/abobi/releases/download/v0.1.0/abobi
```

Make the binary executable:

```bash
chmod +x ./abobi
```

Move the kubectl binary to a file location on your system `PATH`.

```bash
sudo mv ./abobi /usr/local/bin/abobi
sudo chown root: /usr/local/bin/abobi
```

Test to ensure the version you installed is up-to-date:

```bash
abobi --version
```

## Usage

To use abobi to generate a base64 encode kube secret, simply create your `.env` file, here is a sample example:

```.env
#.env

PORT=3000
JWT_SECRET_KEY=condo-orange
AWS_SECRET_KEY=897dgkkh-jdtkl009-007-djknln
AWS_ACCESSS_ID=sHcKVkslDiBBqNcD00i
```

Then run:

```bash
 abobi generate -f env                                                                                                                  21:02:48  ☁  main ☀ 𝝙 𝝙
2023/01/31 21:46:06 Generated secret yaml file: secret.yaml                                                     
```

A sample output format of the `secret.yaml` generated:

```yaml

apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  PORT: MzAwMA==
  JWT_SECRET_KEY: Y29uZG8tb3Jhbmdl
  AWS_SECRET_KEY: ODk3ZGdra2gtamR0a2wwMDktMDA3LWRqa25sbg==
  AWS_ACCESSS_ID: c0hjS1Zrc2xEaUJCcU5jRDAwaQ==


```

## Contributions

This is tool is currently in beta mode and hopefully it gets to alpha and a final release.
If you feel you have ideas on how we can improve this, shoot me a mail at `nonsoamadi@aol.com` and raise a PR or open an issue. I respond fast!

