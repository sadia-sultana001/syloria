package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {

			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for ind, p := range productList {
		if p.ID == product.ID {
			productList[ind] = product
		}
	}
}

func Detele(productID int) {
	var tempList []Product = make([]Product, 0)
	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "A juicy citrus fruit rich in vitamin C and refreshing flavor.",
		Price:       150,
		ImgUrl:      "https://imgs.search.brave.com/2s3-TaRP-gXmQ-4OGmcCA-bN63gOvCPZmdOCLYNro1o/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wMzkv/NjIyLzcxMi9zbWFs/bC9haS1nZW5lcmF0/ZWQtZnJlc2huZXNz/LW9mLWNpdHJ1cy1m/cnVpdC1pbi1uYXR1/cmUtdmlicmFudC1j/b2xvcnMtb2Ytb3Jh/bmdlLXllbGxvdy1h/bmQtZ3JlZW4tZ2Vu/ZXJhdGVkLWJ5LWFp/LWZyZWUtcGhvdG8u/anBn",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "A crisp, sweet fruit known for its fiber and daily health benefit",
		Price:       350,
		ImgUrl:      "https://imgs.search.brave.com/Q_-lNtruM82FHTdMEwWJb8IMCQHn3U_wP1ebaQ6rTdQ/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9ncmVl/bi1hcHBsZS1mcnVp/dHMtMjAzNTM1Nzku/anBn",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "A soft, energy-boosting fruit packed with potassium.",
		Price:       40,
		ImgUrl:      "https://imgs.search.brave.com/-sBhiHfvfEvYXUTqSoo5s30I6sOPMGT1Y7MZsQxlQF8/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly90My5m/dGNkbi5uZXQvanBn/LzA3Lzc0LzI0LzA0/LzM2MF9GXzc3NDI0/MDQ3NF9VV1N0UGlr/ZTFTS1VzVVNNRTdz/Y2pRNk4zM3BGTm11/UC5qcGc",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Grape",
		Description: "A soft, energy-boosting fruit packed with potassium.",
		Price:       540,
		ImgUrl:      "https://imgs.search.brave.com/njVrk3F7HT3TftQUGaa1_UAPtB8TaNGIFQYFvTnxFls/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly91cGxv/YWQud2lraW1lZGlh/Lm9yZy93aWtpcGVk/aWEvY29tbW9ucy8z/LzM3L1NlZWRsZXNz/X2dyYXBlc19vZl9L/YWxsaWRhaWt1cmlj/aGkuanBn",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "A sweet, aromatic tropical fruit often called the “king of fruits.",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/j59IlxitzVSekHpV2bBX_ksJ-tZvv-MaJIJU1C9RnVc/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvOTgw/ODEyNTkwL3Bob3Rv/L2ZyZXNoLXJhdy1t/YW5nb2VzLmpwZz9z/PTYxMng2MTImdz0w/Jms9MjAmYz1jTXlL/ZVVrM3R2MHIyOTVq/TVRaaVdMQ1pfV0FB/c2FqSnFSOWNuYWZx/N1BBPQ",
	}
	prd6 := Product{
		ID: 6,

		Title:       "Pomegranate",
		Description: " A ruby-red fruit filled with tangy, antioxidant-rich seeds.",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/dludsujGmAg-Aqzsuch41owrhWEOe7d4OS_BSEE7TTE/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvODMy/MDk5NTY0L3Bob3Rv/L3dob2xlLWFuZC1w/YXJ0LW9mLXBvbWVn/cmFuYXRlLXdpdGgt/bGVhdmVzLWlzb2xh/dGVkLW9uLXdoaXRl/LmpwZz9zPTYxMng2/MTImdz0wJms9MjAm/Yz1vZ0NRS0Jwa3VD/X05HT2w5dzg2VWlR/UGlrVERDY3p1c2o2/R29MdlVFQXVnPQ",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}
