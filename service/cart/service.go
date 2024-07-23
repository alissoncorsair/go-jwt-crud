package cart

import (
	"fmt"

	"github.com/alissoncorsair/goapi/types"
)

func getCartItemsIDS(item []types.CartItem) ([]int, error) {
	ids := make([]int, len(item))

	for i, value := range item {
		if value.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", value.ProductID)
		}
		ids[i] = value.ProductID
	}

	return ids, nil
}

func (h *Handler) createOrder(products []*types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)

	for _, product := range products {
		productMap[product.ID] = *product
	}
	//check if all products are actually in stock
	if err := checkIfProductIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantity of products in our db
	for _, item := range items {
		prod := productMap[item.ProductID]
		prod.Quantity -= item.Quantity
		h.productStore.UpdateProduct(prod)
	}

	// create the order
	orderId, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})

	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range items {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID:   orderId,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}
	return orderId, totalPrice, nil
}

func checkIfProductIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]

		if !ok {
			return fmt.Errorf("product %d is not available on stock", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d has only %d left", product.ID, product.Quantity)
		}
	}
	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64
	for _, item := range cartItems {
		total += products[item.ProductID].Price * float64(item.Quantity)
	}
	return total
}
