package product

const (
	selectProductQuery = "SELECT * FROM product LIMIT $1 OFFSET $2"

	selectProductByIDQuery = "SELECT * FROM product WHERE id = $1"

	createProductByIDQuery = "INSERT INTO product (name, description, price) VALUES ($1, $2, $3) RETURNING id"

	updateProductQuery = "UPDATE product SET %s WHERE id = %d"

	deleteProductByIDQuery = "DELETE FROM product WHERE id = $1"
)
