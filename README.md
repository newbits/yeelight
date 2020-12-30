## yeelight
This is a library to communicate with the Yeelights in your network. It allows you to easily change the color, brightness, power, etc.

### Installation
To install, run:
```sh
go get github.com/newbits/yeelight
```

### Usage
```go
func main() {
  // Discover and retrieve a single Yeelight in the network
  y, err := yeelight.Discover()
  if err != nil {
    log.Fatal(err)
  }

  // Toggle power
  y.Power()

  // Set color
  y.Color("#112233")

  // Set brightness (percentage)
  y.Brightness(85)

  // Retrieve single property
  prop, _ := y.Prop("power")
  fmt.Printf("Power is %s", prop[0].(string))
}
```

### API Specification
The Yeelight API Specification [can be found here](https://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf)

### Legal
Yeelight is a registered trademark of [Yeelight](https://www.yeelight.com/).

