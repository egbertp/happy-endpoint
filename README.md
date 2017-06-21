# Azure Active Directory Authentication

Using OAuth2

## Usage

Download:

```
mkdir $GOPATH/src/git.rabobank.nl/spn-rotation-tool
cd $GOPATH/src/git.rabobank.nl/spn-rotation-tool
git clone https://potebbb@git.rabobank.nl/scm/~potebbb/spn-rotation-tool.git
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

Run the app
```
$ ./spn-rotation-app
```

Release the app
``` 
$ make release
```


### Use tha app

Point your browser to `http://localhost:7000/AzureLogin`
