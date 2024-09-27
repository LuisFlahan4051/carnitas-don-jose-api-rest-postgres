
export type Branch = {
	name: string;
	address: string;
	branch_safeboxes: any;
	branch_products_stock: any;
	branch_supplies_stock: any;
	branch_articles_stock: any;
	users: any;
	branch_user_roles: any;
	turns: any;
	sales: any;
	inventories: any;
	admin_notifications: any;
};
export type BranchSafebox = {
	name: string;
	content: any;
	branch_id: any;
	safebox_id: any;
};
export type BranchProductStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: any;
	branch_id: any;
	product_id: any;
};
export type BranchSupplyStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: any;
	branch_id: any;
	supply_id: any;
};
export type BranchArticleStock = {
	unit_quantity: any;
	grammage_quantity: any;
	in_use: any;
	branch_id: any;
	article_id: any;
};