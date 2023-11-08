Dockerfile Generator
====================

## To build: 

1. Make sure you have `make` working, and a working go 1.20+ environment.
2. Run `make all`


## To run:

1. Make the binary (see above)
2. Execute it `dist/dockergen-macos-arm64 -L {language} -I {input yaml file}`

### Supported Languages:
* Go (`-L golang`)
* Python (`-L python`)
* Java (`-L java`)  -- Untested. ;)

### Input Yaml Examples:
`input.golang.yaml`, `input.python.yaml`, `input.java.yaml`

These set out the bare-minimum for what you need to define to build a dockerfile. If you don't need environment variables, you can simply omit the block in yaml. 

Generated files will appear in `generated/`
