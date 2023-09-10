# Archiver

## Cli

[Command Line Interface documentation](./docs/cli.md)

## Usage

1) Build a binary file

```bash
go build -o ./build
```

2) Create file with some data

```bash
rm -rf ./tests
mkdir ./tests
touch ./tests/test_1.txt
echo "Hello, world\!" >> ./tests/test_1.txt
```

3) Encode the file

```bash
./build -input ./tests/test_1.txt -output ./tests/test_1_encoded.txt -method shannon-fano -operation encode
```

4) Decode the file

```bash
./build -input ./tests/test_1_encoded.txt -output ./tests/test_1_decoded.txt -method shannon-fano -operation decode
```

5) Compare the source file and the decoded one

```bash
vimdiff ./tests/test_1.txt ./tests/test_1_decoded.txt
```
