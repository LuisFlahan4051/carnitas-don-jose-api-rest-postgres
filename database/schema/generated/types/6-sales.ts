
export type Sale = {
	entry_date: any;
	exit_date: any;
	table_number: any;
	packed: any;
	user_id: any;
	branch_id: any;
	turn_id: any;
	sales_products: any;
	sales_incomes: any;
	sales_expenses: any;
	sales_arguments: any;
};
export type SaleProduct = {
	done: boolean;
	grammage_quantity: any;
	kilogrammage_price: any;
	unit_quantity: any;
	unit_price: any;
	discount: any;
	tax: any;
	sale_id: any;
	product_id: any;
};
export type SaleIncome = {
	sale_id: any;
	income_id: any;
};
export type SaleExpense = {
	sale_id: any;
	expense_id: any;
};
export type SaleArgument = {
	sale_id: any;
	argument_id: any;
};