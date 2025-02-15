package cart

import (
	"fmt"

	"github.com/willy-r/ecom-example/types"
)

func GetCartItemsIds(items []types.CartItem) ([]int, error) {
	itemsIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product with id %d", item.ProductID)
		}

		itemsIds[i] = item.ProductID
	}

	return itemsIds, nil
}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userId int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, p := range ps {
		productMap[p.ID] = p
	}

	if err := CheckIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	totalPrice := CalculateTotalPrice(items, productMap)

	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	orderId, err := h.store.CreateOrder(types.Order{
		UserID:  userId,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}

	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderId,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderId, totalPrice, nil
}

func CheckIfCartIsInStock(items []types.CartItem, productMap map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		p, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if p.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", p.Name)
		}
	}

	return nil
}

func CalculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	var total float64

	for _, item := range items {
		p := productMap[item.ProductID]
		total += p.Price * float64(item.Quantity)
	}

	return total
}
