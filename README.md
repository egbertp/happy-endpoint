# Happy endpoint

This GoLang mini-project creates a binary file that returns a JSON object with a simple `Hello Wolrd!` message.

I use this `happy-endpoint` project to verify the correct working of my infrastructure.

## Usage

Download:

```
mkdir $GOPATH/src/github.com/egbertp/happy-endpoint
cd $GOPATH/src/github.com/egbertp/happy-endpoint
git clone https://github.com/egbertp/happy-endpoint.git
```

### Build and Run

Install dependencies using [Glide Package Management for Go](https://glide.sh/)

```
$ glide install
```

Build the binary
```
$ make build
```

Release the app
```
$ make release
```

Run the app
```
$ ./happy-endpoint
```

### Use tha app

Point your browser to `http://localhost:7000/`
```
{
	message: "Hello world"
}
```
