# Quiz

create an api for merchant seller.
send to: teerapat.sak@kbtg.tech

## Requirement
- language: golang
- database: any

## Point (100 point)
- separate packages via structure (20 point)
- separate config between SIT and PROD (15 point)
- show product `amount` and `stock` as humanize (15 point)
- can be run on docker or docker-compose (20 point)
- connection pool (20 point)
- interface usages (20 point)

## Merchant Fields
- Name
- Bank Account 
- Username
- Password

## Product Fields
- Name
- Amount
- Stocks

## APIs
| Name                 | Method | Endpoint                          |
|----------------------|--------|-----------------------------------|
| Register Merchant    | POST   | /merchant/register                |
| Merchant Information | GET    | /merchant/information/:id         |
| List All Products    | GET    | /merchant/:id/products            |
| Add Product          | POST   | /merchant/:id/product             |

### Register Merchant
- auto generate username and password
- each api must check username/password except register and buy product
- cannot register using the same bank account

### Merchant Information
- response merchant information

### Update Merchant
- can only update name

### List All Products
- list all merchant products with name and amount

### Add Product
- add product for each merchant 
- amount can be present in 2 precision, ex. 100.01, 250.35
- maximum products is 5

### Update Product
- can only update amount
- in case of user already brought product, in sell report must calculate by old amount

### Remove Product
- remove product by product id
- cannot remove if user already brought product

### Sell Reports
- sell report range only by date
- provide list of selling products and amount accumulate with precision point 2 digit
- ensure there is index in all related fields collections,  must prove that there is no table scan

```json
{
	"date": "2018-11-01",
	"products": [
		{"name": "ABC", "selling_volume": 10},
		{"name": "DEF", "selling_volume": 5}
	],
	"accumulate": 100.25 
}
```

### Buy Product
- buy product from merchant with volume

