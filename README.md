# Go CupID

Performing mapping through the Cupid API is very straightforward. You'll first upload your properties inventory to be used as reference and then send a list of properties to be mapped to the API and get the results directly in the response.  #### Getting started is easy:  <!-- type: tab title: 1. Sign-up -->  1. Head to [cupid website](http://mapping.nuitee.com:8084), enter your details to sign up. 2. In your dashboard, click on \"API Keys\" in the menu and generate your first API key. 3. You'll need to add this key as a ```x-api-key``` request header to each request. 4. Now use the tabs above to navigate to the second step and \"2. Setup your workspace\".  <!-- type: tab title: 2. Setup your workspace -->  Upload a CSV file with your properties inventory using the [upload catalog endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDg1OTM1MTE-upload-inventory).  <!-- type: tab title: 3. Start mapping -->  After you have uploaded your inventory, you can send a list of properties using the [mapping endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDUxMTM0NTA-map-hotel-list) and get their mapping result instantly.  <!-- type: tab-end -->

For more information, please visit [https://cupid.travel/support](https://cupid.travel/support)

## Installation

Make sure your project is using Go Modules (it will have a go.mod file in its root if it already is):

```
go mod init
```

Then, reference cupid-go in a Go program with import:

```golang
import cupid "github.com/cupid-mapping/cupid-go"
```

Run any of the normal go commands (build/install/test). The Go toolchain will resolve and fetch the cupid-go module automatically.

Alternatively, you can also explicitly go get the package into a project:

```
go get github.com/cupid-mapping/cupid-go
```

## Usage

### Import
```golang
import cupid "github.com/cupid-mapping/cupid-go"
import cupidModels "github.com/cupid-mapping/cupid-go/models"
````

#### Create client
```golang
cfg := cupid.NewConfiguration()
client := cupid.NewAPIClient(cfg)
auth := context.WithValue(context.Background(), cupid.ContextAPIKey, cupid.APIKey{
	Key: "API-KEY",
})
```

#### Upload inventory
```golang
f, err := os.Open("inventory.csv")
if err != nil {
	// handle err
}
defer f.Close()
resp, _, err := client.DefaultApi.UploadInventory(auth, f, "My inventory name", cupid.ColumnsIndexOpts{
	Id:          0,
	Name:        1,
	Address:     2,
	City:        3,
	CountryCode: 4,
	Latitude:    5,
	Longitude:   6,
})
fmt.Println(resp.InventoryId)
```

#### List inventories
```golang
resp, _, err := client.DefaultApi.ListInventories(auth)
```

#### Map list of hotels
```golang
var hotels = []cupidModels.HotelItem{
	{
		HotelCode:   "123",
		Name:        "Best Western Inn Florence",
		Address:     "7821 Commerce Drive",
		CountryCode: "US",
		Latitude:    39.0014,
		Longitude:   -84.6445,
	},
	{
		HotelCode:   "456",
		Name:        "Baymont by Wyndham Arlington",
		Address:     "2401 Diplomacy Drive",
		CountryCode: "US",
		Latitude:    32.7539,
		Longitude:   -97.0652,
	},
}

opts := &cupid.DefaultApiMapHotelsOpts{
	Body: optional.NewInterface(hotels),
}
resp, _, err := client.DefaultApi.MapHotels(auth, opts)
```

## Models

 - [BadRequestResult](docs/BadRequestResult.md)
 - [ConflictResult](docs/ConflictResult.md)
 - [HotelItem](docs/HotelItem.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InvalidHotel](docs/InvalidHotel.md)
 - [InventoriesUploadBody](docs/InventoriesUploadBody.md)
 - [Inventory](docs/Inventory.md)
 - [MapHotelsResponse](docs/MapHotelsResponse.md)
 - [MappedHotelItem](docs/MappedHotelItem.md)
 - [NotFoundResult](docs/NotFoundResult.md)
 - [Result](docs/Result.md)
 - [ServerErrorResult](docs/ServerErrorResult.md)
 - [UnauthorizedErrorResult](docs/UnauthorizedErrorResult.md)
 - [UploadInventoryResponse](docs/UploadInventoryResponse.md)

## Author

hello@cupid.travel
