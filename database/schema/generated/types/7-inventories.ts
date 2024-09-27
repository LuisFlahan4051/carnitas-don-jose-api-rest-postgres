
export type InventoryType = {
	type: string;
	inventories: any;
};
export type Inventory = {
	acepted: boolean;
	inventory_type_id: any;
	branch_id: any;
	turn_id: any;
	inventory_products_stocks: any;
	inventory_supplies_stocks: any;
	inventory_articles_stocks: any;
};
export type InventoryProductStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: boolean;
	inventory_id: any;
	product_id: any;
};
export type InventorySupplyStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: boolean;
	inventory_id: any;
	supply_id: any;
};
export type InventoryArticleStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: boolean;
	inventory_id: any;
	article_id: any;
};
export type InventorySafebox = {
	inventory_id: any;
	safebox_id: any;
};