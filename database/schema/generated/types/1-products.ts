
export type ID = {
	id: any;
	created_at: any;
	updated_at: any;
	deleted_at: any;
};
export type FoodType = {
	name: string;
	foods: any;
};
export type FoodMeat = {
	name: string;
	foods: any;
};
export type Food = {
	name: string;
	description: any;
	photo: any;
	food_type_id: any;
	food_meat_id: any;
	product_foods: any;
};
export type DrinkSize = {
	name: string;
	drinks: any;
};
export type DrinkFlavor = {
	name: string;
	drinks: any;
};
export type Drink = {
	name: string;
	description: string;
	photo: string;
	drink_size_id: any;
	drink_flavor_id: any;
	product_drinks: any;
};
export type Product = {
	name: string;
	description: any;
	price: number;
	photo: any;
	product_foods: any;
	product_drinks: any;
	branch_products_stock: any;
	sales_products: any;
	inventory_products_stock: any;
};
export type ProductFood = {
	unit_quantity: any;
	grammage_quantity: any;
	food_id: any;
	product_id: any;
};
export type ProductDrink = {
	unit_quantity: any;
	grammage_quantity: any;
	drink_id: any;
	product_id: any;
};