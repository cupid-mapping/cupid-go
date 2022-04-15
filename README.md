# Go Cupid

Performing mapping through the Cupid API is very straightforward. You'll first upload your properties inventory to be used as reference and then send a list of properties to be mapped to the API and get the results directly in the response.

1. Head to [cupid website](https://mapping.nuitee.com), enter your details to sign up.
2. In your dashboard, click on "API Keys" in the menu and generate your first API key.
3. You'll need to use this key to authenticate the client requests as follow.
```golang
auth := context.WithValue(context.Background(), cupid.ContextAPIKey, cupid.APIKey{
	Key: "API-KEY",
})
```

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

#### Import
```golang
import cupid "github.com/cupid-mapping/cupid-go"
import cupidModels "github.com/cupid-mapping/cupid-go/models"
```

#### Create client
```golang
cfg := cupid.NewConfiguration()
client := cupid.NewAPIClient(cfg)
auth := context.WithValue(context.Background(), cupid.ContextAPIKey, cupid.APIKey{
	Key: "API-KEY",
})
```

#### Upload inventory

This function can be used to upload an inventory with a provided CSV file.
You will have to provide the inventory's name and for each column in your file, you'll need match its index with the standard field names provided below.<br>

Assuming you have a CSV file that looks like the following table:

| HotelID     | Hotel_Name   | Hotel_Address | Hotel_Lat  | Hotel_Lon  | Country | City     | ... |
| ----------- | ------------ | ------------- | ---------- | ---------- | ------- | -------- | --- |
| 1408        | Smile Hotel  | 123 main st   | 5.419514   | 3.392694   | AU      | Adelaid  | ... |
| 123         | Cupid resort | 456 north av  | 55.69254   | 37.3167887 | MA      | Zagoura  | ... |

In this case the cupid.ColumnsIndexOpts would be:
```golang
columns := cupid.ColumnsIndexOpts{
	Id:          0,
	Name:        1,
	Address:     2,
	City:        6,
	CountryCode: 5,
	Latitude:    3,
	Longitude:   4,
}
```

> ##### Watch Out!
>
> When uploading your inventory it's important to make sure the process is complete before trying to start mapping.
> Depending on the size of your file, this can take several minutes.
> You can use the ```InventoryDetails``` function to check the processing status id.
>
> **If the inventory is still processing, the mapping function will return an error**


Possible values for the inventory's status ```MappingProcessStatusId```:

| value     | Name           | Description                                        | Action                                                     |
| --------- | -------------- | -------------------------------------------------- | ---------------------------------------------------------- |
| -1        | **Invalid**    | The inventory doesn't contain any valid hotels       | Correct your catalog and try again.                        |
| 0         | **Pending**    | The inventory is being uploaded                      | Wait, no action required.                                  |
| 1         | **Processing** | The inventory is being processed                     | Wait, no action required.                                  |
| 2         | **Done**       | The process is complete                            | **You can start mapping.**                                 |
| 3         | **Failed**     | There was an error while processing                | Retry and if the error persists, feel free to contact us.  |

When you have an inventory with ```Active=true``` and ```MappingProcessStatusId=2``` you can use the following endpoint to perfom mapping:


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
