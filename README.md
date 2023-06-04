# Go Bindings for TP-Link EAP Devices

Provides bindings for interacting with TP-Link access points by using the EAP's
webserver's "API", such as it is. Pretty jank, all in all.

## Disclaimer

So far, this has only been tested on my own EAP-653, mainly due to personal
interests as well as lack of time / hardware. As such, there is guarantee that
this will work on other models as well, although I would be surprised if the
other access points' APIs are drastically different.

## Basic Usage

```go
import (
	"fmt"
	"log"
	"os"

	"github.com/toalaah/go-tplink-eap/pkg/tplink"
)

func main() {
	baseAddr := os.Getenv("TPLINK_ADDR")
	username := os.Getenv("TPLINK_USERNAME")
	password := os.Getenv("TPLINK_PASSWORD")

	c := tplink.NewClient(baseAddr, username, password)
	c.Authenticate()

	if body, err := c.GetDeviceInfo(); err == nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Status: %s\n", body.DeviceName)
	}
}
```

See the [examples](./examples) folder for more.


# Planned Features

- [x] Get / set SSH server
- [ ] Get / set SSIDs
- [x] Get / set LED status
- [ ] Get / set admin credentials
- [ ] Get / set radios
- [ ] ...

## License

This project is released under the terms of the [GPLv3](./LICENSE) license.
